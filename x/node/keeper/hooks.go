package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	baseutils "github.com/sentinel-official/hub/v12/utils"
	"github.com/sentinel-official/hub/v12/x/node/types"
	"github.com/sentinel-official/hub/v12/x/node/types/v3"
)

// SessionInactivePreHook handles operations when a session transitions to the inactive state.
func (k *Keeper) SessionInactivePreHook(ctx sdk.Context, id uint64) error {
	k.Logger(ctx).Info("Running session inactive pre-hook", "id", id)

	// Retrieve the session by ID and return an error if not found.
	item, found := k.GetSession(ctx, id)
	if !found {
		return types.NewErrorSessionNotFound(id)
	}

	// Cast the session to v3.Session; skip further processing if the type assertion fails.
	session, ok := item.(*v3.Session)
	if !ok {
		return nil
	}

	// Ensure the session status is "InactivePending"; return an error if not.
	if !session.Status.Equal(v1base.StatusInactivePending) {
		return types.NewErrorInvalidSessionStatus(session.ID, session.Status)
	}

	// Convert the session's account address from Bech32 format.
	accAddr, err := sdk.AccAddressFromBech32(session.AccAddress)
	if err != nil {
		return err
	}

	// Retrieve the staking share and compute the total payment amount for the session.
	share := k.SessionStakingShare(ctx)
	totalPayment := session.PaymentAmount()

	// Calculate the staking reward and transfer it to the fee collector module.
	reward := baseutils.GetProportionOfCoin(totalPayment, share)
	if err := k.SendCoinFromDepositToModule(ctx, accAddr, k.feeCollectorName, reward); err != nil {
		return err
	}

	// Convert the session's node address from Bech32 format.
	nodeAddr, err := base.NodeAddressFromBech32(session.NodeAddress)
	if err != nil {
		return err
	}

	// Transfer the remaining payment to the node's account.
	payment := totalPayment.Sub(reward)
	if err := k.SendCoinFromDepositToAccount(ctx, accAddr, nodeAddr.Bytes(), payment); err != nil {
		return err
	}

	// Emit an event indicating the payment and staking reward details.
	ctx.EventManager().EmitTypedEvent(
		&v3.EventPay{
			ID:            session.ID,
			AccAddress:    session.AccAddress,
			NodeAddress:   session.NodeAddress,
			Payment:       payment.String(),
			StakingReward: reward.String(),
		},
	)

	// Subtract the refund amount from the user's deposit and handle errors.
	refund := session.RefundAmount()
	if err := k.SubtractDeposit(ctx, accAddr, refund); err != nil {
		return err
	}

	// Emit an event indicating the refund processing details.
	ctx.EventManager().EmitTypedEvent(
		&v3.EventRefund{
			ID:         session.ID,
			AccAddress: session.AccAddress,
			Amount:     refund.String(),
		},
	)

	return nil
}
