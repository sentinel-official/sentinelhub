package v1

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/sentinelhub/v13/x/oracle/keeper"
	"github.com/sentinel-official/sentinelhub/v13/x/oracle/types/v1"
)

var (
	_ v1.MsgServiceServer = (*msgServer)(nil)
)

type msgServer struct {
	keeper.Keeper
}

func NewMsgServiceServer(k keeper.Keeper) v1.MsgServiceServer {
	return &msgServer{k}
}

func (m *msgServer) MsgCreateAsset(c context.Context, req *v1.MsgCreateAssetRequest) (*v1.MsgCreateAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	return m.HandleMsgCreateAsset(ctx, req)
}

func (m *msgServer) MsgDeleteAsset(c context.Context, req *v1.MsgDeleteAssetRequest) (*v1.MsgDeleteAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	return m.HandleMsgDeleteAsset(ctx, req)
}

func (m *msgServer) MsgUpdateAsset(c context.Context, req *v1.MsgUpdateAssetRequest) (*v1.MsgUpdateAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	return m.HandleMsgUpdateAsset(ctx, req)
}

func (m *msgServer) MsgUpdateParams(c context.Context, req *v1.MsgUpdateParamsRequest) (*v1.MsgUpdateParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	return m.HandleMsgUpdateParams(ctx, req)
}
