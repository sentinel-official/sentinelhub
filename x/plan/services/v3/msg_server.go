package v3

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/sentinelhub/v12/x/plan/keeper"
	"github.com/sentinel-official/sentinelhub/v12/x/plan/types/v3"
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

func (m *msgServer) MsgCreatePlan(c context.Context, req *v3.MsgCreatePlanRequest) (*v3.MsgCreatePlanResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return m.HandleMsgCreatePlan(ctx, req)
}

func (m *msgServer) MsgLinkNode(c context.Context, req *v3.MsgLinkNodeRequest) (*v3.MsgLinkNodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return m.HandleMsgLinkNode(ctx, req)
}

func (m *msgServer) MsgUnlinkNode(c context.Context, req *v3.MsgUnlinkNodeRequest) (*v3.MsgUnlinkNodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return m.HandleMsgUnlinkNode(ctx, req)
}

func (m *msgServer) MsgUpdatePlanStatus(c context.Context, req *v3.MsgUpdatePlanStatusRequest) (*v3.MsgUpdatePlanStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return m.HandleMsgUpdatePlanStatus(ctx, req)
}

func (m *msgServer) MsgStartSession(c context.Context, req *v3.MsgStartSessionRequest) (*v3.MsgStartSessionResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return m.HandleMsgStartSession(ctx, req)
}
