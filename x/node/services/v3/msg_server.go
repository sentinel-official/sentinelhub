package v3

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/sentinelhub/v12/x/node/keeper"
	"github.com/sentinel-official/sentinelhub/v12/x/node/types/v3"
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

func (m *msgServer) MsgRegisterNode(c context.Context, req *v3.MsgRegisterNodeRequest) (*v3.MsgRegisterNodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return m.HandleMsgRegisterNode(ctx, req)
}

func (m *msgServer) MsgUpdateNodeDetails(c context.Context, req *v3.MsgUpdateNodeDetailsRequest) (*v3.MsgUpdateNodeDetailsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return m.HandleMsgUpdateNodeDetails(ctx, req)
}

func (m *msgServer) MsgUpdateNodeStatus(c context.Context, req *v3.MsgUpdateNodeStatusRequest) (*v3.MsgUpdateNodeStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return m.HandleMsgUpdateNodeStatus(ctx, req)
}

func (m *msgServer) MsgStartSession(c context.Context, req *v3.MsgStartSessionRequest) (*v3.MsgStartSessionResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return m.HandleMsgStartSession(ctx, req)
}

func (m *msgServer) MsgUpdateParams(c context.Context, req *v3.MsgUpdateParamsRequest) (*v3.MsgUpdateParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return m.HandleMsgUpdateParams(ctx, req)
}
