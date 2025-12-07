package v3

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/sentinelhub/v13/x/provider/keeper"
	"github.com/sentinel-official/sentinelhub/v13/x/provider/types/v3"
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

func (m *msgServer) MsgRegisterProvider(c context.Context, req *v3.MsgRegisterProviderRequest) (*v3.MsgRegisterProviderResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	return m.HandleMsgRegisterProvider(ctx, req)
}

func (m *msgServer) MsgUpdateProviderDetails(c context.Context, req *v3.MsgUpdateProviderDetailsRequest) (*v3.MsgUpdateProviderDetailsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	return m.HandleMsgUpdateProviderDetails(ctx, req)
}

func (m *msgServer) MsgUpdateProviderStatus(c context.Context, req *v3.MsgUpdateProviderStatusRequest) (*v3.MsgUpdateProviderStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	return m.HandleMsgUpdateProviderStatus(ctx, req)
}

func (m *msgServer) MsgUpdateParams(c context.Context, req *v3.MsgUpdateParamsRequest) (*v3.MsgUpdateParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	return m.HandleMsgUpdateParams(ctx, req)
}
