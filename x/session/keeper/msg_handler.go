package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/session/types"
	"github.com/sentinel-official/hub/v12/x/session/types/v3"
)

func (k *Keeper) HandleMsgCancelSession(ctx sdk.Context, msg *v3.MsgCancelSessionRequest) (*v3.MsgCancelSessionResponse, error) {
	fromAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	session, found := k.GetSession(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorSessionNotFound(msg.ID)
	}
	if !session.GetStatus().Equal(v1base.StatusActive) {
		return nil, types.NewErrorInvalidSessionStatus(session.GetID(), session.GetStatus())
	}

	accAddr, err := sdk.AccAddressFromBech32(session.GetAccAddress())
	if err != nil {
		return nil, err
	}

	if !fromAddr.Equals(accAddr) {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	k.DeleteSessionForInactiveAt(ctx, session.GetInactiveAt(), session.GetID())

	inactiveAt := k.GetInactiveAt(ctx)
	session.SetStatus(v1base.StatusInactivePending)
	session.SetInactiveAt(inactiveAt)
	session.SetStatusAt(ctx.BlockTime())

	k.SetSession(ctx, session)
	k.SetSessionForInactiveAt(ctx, session.GetInactiveAt(), session.GetID())

	ctx.EventManager().EmitTypedEvent(
		&v3.EventUpdateStatus{
			ID:          session.GetID(),
			AccAddress:  session.GetAccAddress(),
			NodeAddress: session.GetNodeAddress(),
			Status:      session.GetStatus(),
			StatusAt:    session.GetStatusAt().String(),
		},
	)

	return &v3.MsgCancelSessionResponse{}, nil
}

func (k *Keeper) HandleMsgUpdateSession(ctx sdk.Context, msg *v3.MsgUpdateSessionRequest) (*v3.MsgUpdateSessionResponse, error) {
	session, found := k.GetSession(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorSessionNotFound(msg.ID)
	}
	if msg.From != session.GetNodeAddress() {
		return nil, types.NewErrorUnauthorized(msg.From)
	}
	if session.GetStatus().Equal(v1base.StatusInactive) {
		return nil, types.NewErrorInvalidSessionStatus(session.GetID(), session.GetStatus())
	}

	if msg.DownloadBytes.LT(session.GetDownloadBytes()) {
		return nil, types.NewErrorInvalidDownloadBytes(msg.DownloadBytes)
	}
	if msg.UploadBytes.LT(session.GetUploadBytes()) {
		return nil, types.NewErrorInvalidUploadBytes(msg.UploadBytes)
	}
	if msg.Duration < session.GetDuration() {
		return nil, types.NewErrorInvalidDuration(msg.Duration)
	}

	if ok := k.ProofVerificationEnabled(ctx); ok {
		accAddr, err := sdk.AccAddressFromBech32(session.GetAccAddress())
		if err != nil {
			return nil, err
		}

		if err := k.VerifySignature(ctx, accAddr, msg.Proof(), msg.Signature); err != nil {
			return nil, types.NewErrorInvalidSignature(msg.Signature)
		}
	}

	if err := k.SessionUpdatePreHook(ctx, session.GetID(), msg.Bytes()); err != nil {
		return nil, err
	}

	if session.GetStatus().Equal(v1base.StatusActive) {
		k.DeleteSessionForInactiveAt(ctx, session.GetInactiveAt(), session.GetID())
	}

	session.SetDownloadBytes(msg.DownloadBytes)
	session.SetUploadBytes(msg.UploadBytes)
	session.SetDuration(msg.Duration)

	if session.GetStatus().Equal(v1base.StatusActive) {
		inactiveAt := k.GetInactiveAt(ctx)
		session.SetInactiveAt(inactiveAt)
	}

	k.SetSession(ctx, session)
	if session.GetStatus().Equal(v1base.StatusActive) {
		k.SetSessionForInactiveAt(ctx, session.GetInactiveAt(), session.GetID())
	}

	ctx.EventManager().EmitTypedEvent(
		&v3.EventUpdateDetails{
			ID:            session.GetID(),
			AccAddress:    session.GetAccAddress(),
			NodeAddress:   session.GetNodeAddress(),
			DownloadBytes: session.GetDownloadBytes().String(),
			UploadBytes:   session.GetUploadBytes().String(),
			Duration:      session.GetDuration(),
		},
	)

	return &v3.MsgUpdateSessionResponse{}, nil
}

func (k *Keeper) HandleMsgUpdateParams(ctx sdk.Context, msg *v3.MsgUpdateParamsRequest) (*v3.MsgUpdateParamsResponse, error) {
	if msg.From != k.authority {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	k.SetParams(ctx, msg.Params)
	return &v3.MsgUpdateParamsResponse{}, nil
}
