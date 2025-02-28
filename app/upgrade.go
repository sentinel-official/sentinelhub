package app

import (
	"fmt"
	"math"
	"time"

	sdkmath "cosmossdk.io/math"
	v2wasmmigrations "github.com/CosmWasm/wasmd/x/wasm/migrations/v2"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkmodule "github.com/cosmos/cosmos-sdk/types/module"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authvestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	v1govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/cosmos/cosmos-sdk/x/group"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/nft"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/cosmos/ibc-go/v7/modules/core/exported"
	ibctmmigrations "github.com/cosmos/ibc-go/v7/modules/light-clients/07-tendermint/migrations"

	oracletypes "github.com/sentinel-official/hub/v12/x/oracle/types"
)

const (
	UpgradeName = "v12_0_0"
)

var (
	StoreUpgrades = &storetypes.StoreUpgrades{
		Added: []string{
			consensustypes.ModuleName,
			crisistypes.ModuleName,
			group.ModuleName,
			nft.ModuleName,
			oracletypes.ModuleName,
		},
	}
)

func UpgradeHandler(
	cdc codec.Codec, mm *sdkmodule.Manager, configurator sdkmodule.Configurator, keepers Keepers,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, fromVM sdkmodule.VersionMap) (sdkmodule.VersionMap, error) {
		keyTables := map[string]paramstypes.KeyTable{
			// Cosmos SDK subspaces
			authtypes.ModuleName:         authtypes.ParamKeyTable(),
			banktypes.ModuleName:         banktypes.ParamKeyTable(),
			crisistypes.ModuleName:       crisistypes.ParamKeyTable(),
			distributiontypes.ModuleName: distributiontypes.ParamKeyTable(),
			govtypes.ModuleName:          v1govtypes.ParamKeyTable(),
			minttypes.ModuleName:         minttypes.ParamKeyTable(),
			slashingtypes.ModuleName:     slashingtypes.ParamKeyTable(),
			stakingtypes.ModuleName:      stakingtypes.ParamKeyTable(),

			// Other subspaces
			wasmtypes.ModuleName: v2wasmmigrations.ParamKeyTable(),
		}

		for name, table := range keyTables {
			subspace, ok := keepers.ParamsKeeper.GetSubspace(name)
			if !ok {
				return nil, fmt.Errorf("params subspace does not exist for module: %s", name)
			}
			if subspace.HasKeyTable() {
				continue
			}

			subspace.WithKeyTable(table)
		}

		legacyParamStore := keepers.ParamsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramstypes.ConsensusParamsKeyTable())
		baseapp.MigrateParams(ctx, legacyParamStore, &keepers.ConsensusKeeper)

		newVM, err := mm.RunMigrations(ctx, configurator, fromVM)
		if err != nil {
			return newVM, err
		}

		_, err = ibctmmigrations.PruneExpiredConsensusStates(ctx, cdc, keepers.IBCKeeper.ClientKeeper)
		if err != nil {
			return nil, err
		}

		govParams := keepers.GovKeeper.GetParams(ctx)
		govParams.MinInitialDepositRatio = sdkmath.LegacyNewDecWithPrec(2, 1).String()
		if err := keepers.GovKeeper.SetParams(ctx, govParams); err != nil {
			return nil, err
		}

		stakingParams := keepers.StakingKeeper.GetParams(ctx)
		stakingParams.MinCommissionRate = sdkmath.LegacyNewDecWithPrec(5, 2)
		if err := keepers.StakingKeeper.SetParams(ctx, stakingParams); err != nil {
			return nil, err
		}

		validators := keepers.StakingKeeper.GetAllValidators(ctx)
		for _, validator := range validators {
			if validator.Commission.Rate.GTE(stakingParams.MinCommissionRate) {
				continue
			}

			validator.Commission.Rate = stakingParams.MinCommissionRate
			validator.Commission.UpdateTime = ctx.BlockTime()
			if validator.Commission.MaxRate.LT(validator.Commission.Rate) {
				validator.Commission.MaxRate = validator.Commission.Rate
			}

			if err := keepers.StakingKeeper.Hooks().BeforeValidatorModified(ctx, validator.GetOperator()); err != nil {
				return nil, err
			}

			keepers.StakingKeeper.SetValidator(ctx, validator)
			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					stakingtypes.EventTypeEditValidator,
					sdk.NewAttribute(stakingtypes.AttributeKeyCommissionRate, validator.Commission.String()),
					sdk.NewAttribute(stakingtypes.AttributeKeyMinSelfDelegation, validator.MinSelfDelegation.String()),
				),
			)
		}

		ibcClientParams := keepers.IBCKeeper.ClientKeeper.GetParams(ctx)
		ibcClientParams.AllowedClients = append(ibcClientParams.AllowedClients, exported.Localhost)
		keepers.IBCKeeper.ClientKeeper.SetParams(ctx, ibcClientParams)

		if err := migrateFoundationAccount(ctx, keepers.AccountKeeper, keepers.BankKeeper, keepers.StakingKeeper); err != nil {
			return nil, err
		}

		return newVM, nil
	}
}

func completeAllRedelegations(
	ctx sdk.Context, k *stakingkeeper.Keeper, accAddr sdk.AccAddress, completionTime time.Time,
) error {
	for _, item := range k.GetRedelegations(ctx, accAddr, math.MaxInt16) {
		for i := range item.Entries {
			item.Entries[i].CompletionTime = completionTime
		}

		k.SetRedelegation(ctx, item)

		fromAddr, err := sdk.ValAddressFromBech32(item.ValidatorSrcAddress)
		if err != nil {
			return err
		}

		toAddr, err := sdk.ValAddressFromBech32(item.ValidatorDstAddress)
		if err != nil {
			return err
		}

		balances, err := k.CompleteRedelegation(ctx, accAddr, fromAddr, toAddr)
		if err != nil {
			return err
		}

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				stakingtypes.EventTypeCompleteRedelegation,
				sdk.NewAttribute(sdk.AttributeKeyAmount, balances.String()),
				sdk.NewAttribute(stakingtypes.AttributeKeyDelegator, item.DelegatorAddress),
				sdk.NewAttribute(stakingtypes.AttributeKeySrcValidator, item.ValidatorSrcAddress),
				sdk.NewAttribute(stakingtypes.AttributeKeyDstValidator, item.ValidatorDstAddress),
			),
		)
	}

	return nil
}

func undelegateAllDelegations(ctx sdk.Context, k *stakingkeeper.Keeper, accAddr sdk.AccAddress) error {
	for _, item := range k.GetAllDelegatorDelegations(ctx, accAddr) {
		valAddr, err := sdk.ValAddressFromBech32(item.ValidatorAddress)
		if err != nil {
			return err
		}

		validator, found := k.GetValidator(ctx, valAddr)
		if !found {
			return fmt.Errorf("validator %s does not exist", valAddr)
		}

		completionTime, err := k.Undelegate(ctx, accAddr, valAddr, item.GetShares())
		if err != nil {
			return err
		}

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				stakingtypes.EventTypeUnbond,
				sdk.NewAttribute(stakingtypes.AttributeKeyValidator, item.ValidatorAddress),
				sdk.NewAttribute(sdk.AttributeKeyAmount, validator.TokensFromSharesTruncated(item.GetShares()).String()),
				sdk.NewAttribute(stakingtypes.AttributeKeyDelegator, item.DelegatorAddress),
				sdk.NewAttribute(stakingtypes.AttributeKeyCompletionTime, completionTime.Format(time.RFC3339)),
			),
		)
	}

	return nil
}

func completeAllUnbondingDelegations(
	ctx sdk.Context, k *stakingkeeper.Keeper, accAddr sdk.AccAddress, completionTime time.Time,
) error {
	for _, item := range k.GetAllUnbondingDelegations(ctx, accAddr) {
		for i := range item.Entries {
			item.Entries[i].CompletionTime = completionTime
		}

		k.SetUnbondingDelegation(ctx, item)

		valAddr, err := sdk.ValAddressFromBech32(item.ValidatorAddress)
		if err != nil {
			return err
		}

		balances, err := k.CompleteUnbonding(ctx, accAddr, valAddr)
		if err != nil {
			return err
		}

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				stakingtypes.EventTypeCompleteUnbonding,
				sdk.NewAttribute(sdk.AttributeKeyAmount, balances.String()),
				sdk.NewAttribute(stakingtypes.AttributeKeyValidator, item.ValidatorAddress),
				sdk.NewAttribute(stakingtypes.AttributeKeyDelegator, item.DelegatorAddress),
			),
		)
	}

	return nil
}
func migrateFoundationAccount(
	ctx sdk.Context, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, sk *stakingkeeper.Keeper,
) error {
	// Parse Bech32 account address
	addr, err := sdk.AccAddressFromBech32("sent1vv8kmwrs24j5emzw8dp7k8satgea62l7knegd7")
	if err != nil {
		return fmt.Errorf("failed to parse address: %w", err)
	}

	// Complete all redelegations
	if err := completeAllRedelegations(ctx, sk, addr, ctx.BlockTime()); err != nil {
		return fmt.Errorf("failed to complete redelegations: %w", err)
	}

	// Undelegate all delegations
	if err := undelegateAllDelegations(ctx, sk, addr); err != nil {
		return fmt.Errorf("failed to undelegate delegations: %w", err)
	}

	// Complete all unbonding delegations
	if err := completeAllUnbondingDelegations(ctx, sk, addr, ctx.BlockTime()); err != nil {
		return fmt.Errorf("failed to complete unbonding delegations: %w", err)
	}

	// Retrieve account
	account := ak.GetAccount(ctx, addr)

	// Ensure the account is a ContinuousVestingAccount
	vestingAccount, ok := account.(*authvestingtypes.ContinuousVestingAccount)
	if !ok {
		return fmt.Errorf("invalid account type; expected ContinuousVestingAccount, got %T", account)
	}

	// Create a new ContinuousVestingAccount with updated end time
	vestingAccount = authvestingtypes.NewContinuousVestingAccount(
		vestingAccount.BaseAccount,
		vestingAccount.OriginalVesting,
		0,
		ctx.BlockTime().Unix(),
	)

	// Get balances and calculate total bonded and unbonding amounts
	balances := bk.GetAllBalances(ctx, addr)
	bonded := sk.GetDelegatorBonded(ctx, addr)
	unbonding := sk.GetDelegatorUnbonding(ctx, addr)

	// Add bonded and unbonding amounts to the balance
	amount := sdk.NewCoin("udvpn", bonded.Add(unbonding))
	balance := balances.Add(amount)

	// Track delegation and update account
	vestingAccount.TrackDelegation(ctx.BlockTime(), balance, sdk.NewCoins(amount))
	ak.SetAccount(ctx, vestingAccount)

	// Transfer spendable coins to new address
	toAddr, err := sdk.AccAddressFromBech32("")
	if err != nil {
		return err
	}

	// Retrieve the spendable balance
	spendableCoins := bk.SpendableCoins(ctx, addr)

	// Transfer spendable balance to new address
	if err := bk.SendCoins(ctx, addr, toAddr, spendableCoins); err != nil {
		return err
	}

	return nil
}
