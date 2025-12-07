package v3

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/sentinelhub/v13/x/subscription/keeper"
	"github.com/sentinel-official/sentinelhub/v13/x/subscription/types/v3"
)

var (
	_ v3.MsgServiceServer = (*msgServer)(nil)
)

type msgServer struct {
	keeper.Keeper
}

func NewMsgServiceServer(k keeper.Keeper) v3.MsgServiceServer {
	return &msgServer{k}
}

func (m *msgServer) MsgCancelSubscription(c context.Context, req *v3.MsgCancelSubscriptionRequest) (*v3.MsgCancelSubscriptionResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	return m.HandleMsgCancelSubscription(ctx, req)
}

func (m *msgServer) MsgRenewSubscription(c context.Context, req *v3.MsgRenewSubscriptionRequest) (*v3.MsgRenewSubscriptionResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	return m.HandleMsgRenewSubscription(ctx, req)
}

func (m *msgServer) MsgShareSubscription(c context.Context, req *v3.MsgShareSubscriptionRequest) (*v3.MsgShareSubscriptionResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	return m.HandleMsgShareSubscription(ctx, req)
}

func (m *msgServer) MsgStartSubscription(c context.Context, req *v3.MsgStartSubscriptionRequest) (*v3.MsgStartSubscriptionResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	return m.HandleMsgStartSubscription(ctx, req)
}

func (m *msgServer) MsgUpdateSubscription(c context.Context, req *v3.MsgUpdateSubscriptionRequest) (*v3.MsgUpdateSubscriptionResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	return m.HandleMsgUpdateSubscription(ctx, req)
}

func (m *msgServer) MsgStartSession(c context.Context, req *v3.MsgStartSessionRequest) (*v3.MsgStartSessionResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	return m.HandleMsgStartSession(ctx, req)
}

func (m *msgServer) MsgUpdateParams(c context.Context, req *v3.MsgUpdateParamsRequest) (*v3.MsgUpdateParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	return m.HandleMsgUpdateParams(ctx, req)
}
