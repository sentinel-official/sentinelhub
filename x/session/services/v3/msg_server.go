package v3

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/sentinelhub/v13/x/session/keeper"
	"github.com/sentinel-official/sentinelhub/v13/x/session/types/v3"
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

func (m *msgServer) MsgCancelSession(c context.Context, req *v3.MsgCancelSessionRequest) (*v3.MsgCancelSessionResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	return m.HandleMsgCancelSession(ctx, req)
}

func (m *msgServer) MsgUpdateSession(c context.Context, req *v3.MsgUpdateSessionRequest) (*v3.MsgUpdateSessionResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	return m.HandleMsgUpdateSession(ctx, req)
}

func (m *msgServer) MsgUpdateParams(c context.Context, req *v3.MsgUpdateParamsRequest) (*v3.MsgUpdateParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	return m.HandleMsgUpdateParams(ctx, req)
}
