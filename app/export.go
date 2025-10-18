package app

import (
	"encoding/json"
	"log"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func (a *App) ExportAppStateAndValidators(
	forZeroHeight bool, jailAllowedAddrs []string, modulesToExport []string,
) (servertypes.ExportedApp, error) {
	ctx := a.NewContext(true, tmproto.Header{Height: a.LastBlockHeight()})

	height := a.LastBlockHeight() + 1
	if forZeroHeight {
		height = 0

		a.prepForZeroHeightGenesis(ctx, jailAllowedAddrs)
	}

	consensusParams := a.GetConsensusParams(ctx)
	genState := a.mm.ExportGenesis(ctx, a.Codec)

	validators, err := staking.WriteValidators(ctx, a.StakingKeeper)
	if err != nil {
		return servertypes.ExportedApp{}, err
	}

	appState, err := json.MarshalIndent(genState, "", "  ")
	if err != nil {
		return servertypes.ExportedApp{}, err
	}

	return servertypes.ExportedApp{
		AppState:        appState,
		Validators:      validators,
		Height:          height,
		ConsensusParams: consensusParams,
	}, nil
}

func (a *App) prepForZeroHeightGenesis(ctx sdk.Context, jailAllowedAddrs []string) {
	applyAllowedAddrs := len(jailAllowedAddrs) > 0

	allowedAddrsMap := make(map[string]bool)

	for _, addr := range jailAllowedAddrs {
		_, err := sdk.ValAddressFromBech32(addr)
		if err != nil {
			log.Fatal(err)
		}

		allowedAddrsMap[addr] = true
	}

	a.CrisisKeeper.AssertInvariants(ctx)

	a.StakingKeeper.IterateValidators(ctx, func(_ int64, val stakingtypes.ValidatorI) (stop bool) {
		_, _ = a.DistributionKeeper.WithdrawValidatorCommission(ctx, val.GetOperator())

		return false
	})

	dels := a.StakingKeeper.GetAllDelegations(ctx)
	for _, delegation := range dels {
		valAddr, err := sdk.ValAddressFromBech32(delegation.ValidatorAddress)
		if err != nil {
			panic(err)
		}

		delAddr, err := sdk.AccAddressFromBech32(delegation.DelegatorAddress)
		if err != nil {
			panic(err)
		}

		_, _ = a.DistributionKeeper.WithdrawDelegationRewards(ctx, delAddr, valAddr)
	}

	a.DistributionKeeper.DeleteAllValidatorSlashEvents(ctx)

	a.DistributionKeeper.DeleteAllValidatorHistoricalRewards(ctx)

	height := ctx.BlockHeight()
	ctx = ctx.WithBlockHeight(0)

	a.StakingKeeper.IterateValidators(ctx, func(_ int64, val stakingtypes.ValidatorI) (stop bool) {
		scraps := a.DistributionKeeper.GetValidatorOutstandingRewardsCoins(ctx, val.GetOperator())
		feePool := a.DistributionKeeper.GetFeePool(ctx)
		feePool.CommunityPool = feePool.CommunityPool.Add(scraps...)
		a.DistributionKeeper.SetFeePool(ctx, feePool)

		a.DistributionKeeper.Hooks().AfterValidatorCreated(ctx, val.GetOperator())

		return false
	})

	for _, del := range dels {
		valAddr, err := sdk.ValAddressFromBech32(del.ValidatorAddress)
		if err != nil {
			panic(err)
		}

		delAddr, err := sdk.AccAddressFromBech32(del.DelegatorAddress)
		if err != nil {
			panic(err)
		}

		a.DistributionKeeper.Hooks().BeforeDelegationCreated(ctx, delAddr, valAddr)
		a.DistributionKeeper.Hooks().AfterDelegationModified(ctx, delAddr, valAddr)
	}

	ctx = ctx.WithBlockHeight(height)

	a.StakingKeeper.IterateRedelegations(ctx, func(_ int64, red stakingtypes.Redelegation) (stop bool) {
		for i := range red.Entries {
			red.Entries[i].CreationHeight = 0
		}

		a.StakingKeeper.SetRedelegation(ctx, red)

		return false
	})

	a.StakingKeeper.IterateUnbondingDelegations(ctx, func(_ int64, ubd stakingtypes.UnbondingDelegation) (stop bool) {
		for i := range ubd.Entries {
			ubd.Entries[i].CreationHeight = 0
		}

		a.StakingKeeper.SetUnbondingDelegation(ctx, ubd)

		return false
	})

	store := ctx.KVStore(a.KV(stakingtypes.StoreKey))
	iterator := sdk.KVStoreReversePrefixIterator(store, stakingtypes.ValidatorsKey)
	counter := int16(0)

	for ; iterator.Valid(); iterator.Next() {
		addr := sdk.ValAddress(iterator.Key()[1:])

		validator, found := a.StakingKeeper.GetValidator(ctx, addr)
		if !found {
			panic("expected validator, not found")
		}

		validator.UnbondingHeight = 0
		if applyAllowedAddrs && !allowedAddrsMap[addr.String()] {
			validator.Jailed = true
		}

		a.StakingKeeper.SetValidator(ctx, validator)

		counter++
	}

	iterator.Close()

	_, err := a.StakingKeeper.ApplyAndReturnValidatorSetUpdates(ctx)
	if err != nil {
		log.Fatal(err)
	}

	a.SlashingKeeper.IterateValidatorSigningInfos(
		ctx,
		func(addr sdk.ConsAddress, info slashingtypes.ValidatorSigningInfo) (stop bool) {
			info.StartHeight = 0
			a.SlashingKeeper.SetValidatorSigningInfo(ctx, addr, info)

			return false
		},
	)
}
