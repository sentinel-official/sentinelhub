package v1

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/sentinelhub/v12/x/swap/keeper"
	"github.com/sentinel-official/sentinelhub/v12/x/swap/types/v1"
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

func (m *msgServer) MsgSwap(c context.Context, req *v1.MsgSwapRequest) (*v1.MsgSwapResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	return m.HandleMsgSwap(ctx, req)
}
