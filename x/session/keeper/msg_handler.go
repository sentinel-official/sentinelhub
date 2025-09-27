package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	v1base "github.com/sentinel-official/sentinelhub/v12/types/v1"
	"github.com/sentinel-official/sentinelhub/v12/x/session/types"
	"github.com/sentinel-official/sentinelhub/v12/x/session/types/v3"
)

// HandleMsgCancelSession handles a request to cancel an active session.
// It verifies ownership, updates the session status, and emits a status change event.
func (k *Keeper) HandleMsgCancelSession(ctx sdk.Context, msg *v3.MsgCancelSessionRequest) (*v3.MsgCancelSessionResponse, error) {
	// Parse and validate the requester's address
	fromAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	// Fetch the session and verify existence and active status
	session, found := k.GetSession(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorSessionNotFound(msg.ID)
	}

	if !session.GetStatus().Equal(v1base.StatusActive) {
		return nil, types.NewErrorInvalidSessionStatus(session.GetID(), session.GetStatus())
	}

	// Parse and verify ownership against the session's account address
	accAddr, err := sdk.AccAddressFromBech32(session.GetAccAddress())
	if err != nil {
		return nil, err
	}

	if !fromAddr.Equals(accAddr) {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	// Clear old inactive index before status update
	k.DeleteSessionForInactiveAt(ctx, session.GetInactiveAt(), session.GetID())

	// Update session status to inactive pending and set timestamps
	inactiveAt := k.GetInactiveAt(ctx)

	session.SetStatus(v1base.StatusInactivePending)
	session.SetInactiveAt(inactiveAt)
	session.SetStatusAt(ctx.BlockTime())

	// Persist the session with new status and reindex
	k.SetSession(ctx, session)
	k.SetSessionForInactiveAt(ctx, session.GetInactiveAt(), session.GetID())

	// Emit event indicating session status change
	ctx.EventManager().EmitTypedEvent(
		&v3.EventUpdateStatus{
			SessionID:   session.GetID(),
			AccAddress:  session.GetAccAddress(),
			NodeAddress: session.GetNodeAddress(),
			Status:      session.GetStatus().String(),
		},
	)

	return &v3.MsgCancelSessionResponse{}, nil
}

// HandleMsgUpdateSession handles a request to update session metrics (bytes, duration).
// It validates ownership, checks monotonic increases, optionally verifies proof, and applies updates.
func (k *Keeper) HandleMsgUpdateSession(ctx sdk.Context, msg *v3.MsgUpdateSessionRequest) (*v3.MsgUpdateSessionResponse, error) {
	// Fetch the session and verify existence
	session, found := k.GetSession(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorSessionNotFound(msg.ID)
	}

	// Confirm node address authorization
	if msg.From != session.GetNodeAddress() {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	// Validate updated usage metrics are not regressing
	if msg.DownloadBytes.LT(session.GetDownloadBytes()) {
		return nil, types.NewErrorInvalidDownloadBytes(msg.DownloadBytes)
	}

	if msg.UploadBytes.LT(session.GetUploadBytes()) {
		return nil, types.NewErrorInvalidUploadBytes(msg.UploadBytes)
	}

	if msg.Duration < session.GetDuration() {
		return nil, types.NewErrorInvalidDuration(msg.Duration)
	}

	// Optionally verify proof of update if enabled by module parameters
	if ok := k.ProofVerificationEnabled(ctx); ok {
		accAddr, err := sdk.AccAddressFromBech32(session.GetAccAddress())
		if err != nil {
			return nil, err
		}

		if err := k.VerifySignature(ctx, accAddr, msg.Proof(), msg.Signature); err != nil {
			return nil, types.NewErrorInvalidSignature(msg.Signature)
		}
	}

	// Execute pre-update hook to allow custom logic
	if err := k.SessionUpdatePreHook(ctx, session.GetID(), msg.Bytes()); err != nil {
		return nil, err
	}

	// Remove old inactive index if the session is still active
	if session.GetStatus().Equal(v1base.StatusActive) {
		k.DeleteSessionForInactiveAt(ctx, session.GetInactiveAt(), session.GetID())
	}

	// Apply new metrics to the session
	session.SetDownloadBytes(msg.DownloadBytes)
	session.SetUploadBytes(msg.UploadBytes)
	session.SetDuration(msg.Duration)

	// If active, refresh inactivity timeout
	if session.GetStatus().Equal(v1base.StatusActive) {
		inactiveAt := k.GetInactiveAt(ctx)
		session.SetInactiveAt(inactiveAt)
	}

	// Persist updated session and reindex if still active
	k.SetSession(ctx, session)

	if session.GetStatus().Equal(v1base.StatusActive) {
		k.SetSessionForInactiveAt(ctx, session.GetInactiveAt(), session.GetID())
	}

	// Emit event reflecting the updated session usage
	ctx.EventManager().EmitTypedEvent(
		&v3.EventUpdateDetails{
			SessionID:     session.GetID(),
			AccAddress:    session.GetAccAddress(),
			NodeAddress:   session.GetNodeAddress(),
			DownloadBytes: session.GetDownloadBytes().String(),
			UploadBytes:   session.GetUploadBytes().String(),
			Duration:      session.GetDuration().String(),
		},
	)

	return &v3.MsgUpdateSessionResponse{}, nil
}

// HandleMsgUpdateParams allows the module authority to update session module parameters.
// It enforces authority checks and saves the updated configuration.
func (k *Keeper) HandleMsgUpdateParams(ctx sdk.Context, msg *v3.MsgUpdateParamsRequest) (*v3.MsgUpdateParamsResponse, error) {
	// Restrict access to the designated authority account
	if msg.From != k.authority {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	// Persist the new parameter configuration
	k.SetParams(ctx, msg.Params)

	return &v3.MsgUpdateParamsResponse{}, nil
}
