package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	baseutils "github.com/sentinel-official/hub/v12/utils"
	"github.com/sentinel-official/hub/v12/x/lease/types/v1"
)

// handleInactiveLeases processes all leases that have become inactive at the current block time.
func (k *Keeper) handleInactiveLeases(ctx sdk.Context) {
	// Iterate through leases that have become inactive at the current block time
	k.IterateLeasesForInactiveAt(ctx, ctx.BlockTime(), func(_ int, item v1.Lease) bool {
		k.Logger(ctx).Info("Handling inactive lease", "id", item.ID)

		// Create a message to end the lease
		msg := &v1.MsgEndLeaseRequest{
			From: item.ProvAddress,
			ID:   item.ID,
		}

		// Get the appropriate handler for processing the message
		handler := k.router.Handler(msg)
		if handler == nil {
			panic(fmt.Errorf("nil handler for message route: %s", sdk.MsgTypeURL(msg)))
		}

		// Execute the handler and process the lease termination request
		resp, err := handler(ctx, msg)
		if err != nil {
			panic(err)
		}

		// Emit any events generated during the lease termination process
		ctx.EventManager().EmitEvents(resp.GetEvents())
		return false
	})
}

// handleLeasePayouts processes payouts for leases that are due at the current block time.
func (k *Keeper) handleLeasePayouts(ctx sdk.Context) {
	// Get the staking share to calculate the reward portion
	share := k.StakingShare(ctx)

	// Iterate through leases that are due for payout at the current block time
	k.IterateLeasesForPayoutAt(ctx, ctx.BlockTime(), func(_ int, item v1.Lease) bool {
		k.Logger(ctx).Info("Handling lease payout", "id", item.ID)

		// Get node and provider addresses from Bech32 strings
		nodeAddr, err := base.NodeAddressFromBech32(item.NodeAddress)
		if err != nil {
			panic(err)
		}

		provAddr, err := base.ProvAddressFromBech32(item.ProvAddress)
		if err != nil {
			panic(err)
		}

		totalPayment := item.Price.QuotePrice()

		// Calculate the staking reward and send it to the module
		reward := baseutils.GetProportionOfCoin(totalPayment, share)
		if err := k.SendCoinFromDepositToModule(ctx, provAddr.Bytes(), k.feeCollectorName, reward); err != nil {
			panic(err)
		}

		// Calculate the remaining payment and send it to the node address
		payment := totalPayment.Sub(reward)
		if err := k.SendCoinFromDepositToAccount(ctx, provAddr.Bytes(), nodeAddr.Bytes(), payment); err != nil {
			panic(err)
		}

		// Emit an event for the payment processing
		ctx.EventManager().EmitTypedEvent(
			&v1.EventPay{
				ID:            item.ID,
				NodeAddress:   item.NodeAddress,
				ProvAddress:   item.ProvAddress,
				Payment:       payment.String(),
				StakingReward: reward.String(),
			},
		)

		// Remove the lease from the payout queue as it has been processed
		k.DeleteLeaseForPayoutAt(ctx, item.PayoutAt(), item.ID)

		// Update lease hours
		item.Hours = item.Hours + 1

		// Update the lease in the store with new details
		k.SetLease(ctx, item)
		k.SetLeaseForPayoutAt(ctx, item.PayoutAt(), item.ID)

		// Emit an event for the updated lease details
		ctx.EventManager().EmitTypedEvent(
			&v1.EventUpdate{
				ID:          item.ID,
				NodeAddress: item.NodeAddress,
				ProvAddress: item.ProvAddress,
				Hours:       item.Hours,
				PayoutAt:    item.PayoutAt().String(),
			},
		)

		return false
	})
}

// handleLeaseRenewals processes lease renewals that are due at the current block time.
func (k *Keeper) handleLeaseRenewals(ctx sdk.Context) {
	// Iterate through leases that are due for renewal at the current block time
	k.IterateLeasesForRenewalAt(ctx, ctx.BlockTime(), func(_ int, item v1.Lease) bool {
		k.Logger(ctx).Info("Handling lease renewal", "id", item.ID)

		// Create a message to renew the lease
		msg := &v1.MsgRenewLeaseRequest{
			From:  item.ProvAddress,
			ID:    item.ID,
			Hours: item.MaxHours,
			Denom: item.Price.Denom,
		}

		// Get the appropriate handler for processing the renewal message
		handler := k.router.Handler(msg)
		if handler == nil {
			panic(fmt.Errorf("nil handler for message route: %s", sdk.MsgTypeURL(msg)))
		}

		// Create a cache context to prevent state changes until lease renewal is successful
		cc, write := ctx.CacheContext()

		// Execute the handler to process the lease renewal
		resp, err := handler(cc, msg)
		if err != nil {
			k.Logger(cc).Error("Failed to handle lease renewal", "id", item.ID, "msg", err)
			return false
		}

		// If the renewal is successful, commit the changes to the main context
		defer write()

		// Emit any events generated during the lease renewal process
		cc.EventManager().EmitEvents(resp.GetEvents())
		return false
	})
}

// BeginBlock is called at the beginning of each block to handle lease-related operations.
func (k *Keeper) BeginBlock(ctx sdk.Context) {
	// Handle lease renewals at the beginning of each block
	k.handleLeaseRenewals(ctx)

	// Handle lease payouts at the beginning of each block
	k.handleLeasePayouts(ctx)

	// Handle leases that have become inactive at the beginning of each block
	k.handleInactiveLeases(ctx)
}
