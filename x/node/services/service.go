package services

import (
	sdkmodule "github.com/cosmos/cosmos-sdk/types/module"

	"github.com/sentinel-official/hub/v12/x/node/keeper"
	"github.com/sentinel-official/hub/v12/x/node/services/v3"
	v3types "github.com/sentinel-official/hub/v12/x/node/types/v3"
)

func RegisterServices(configurator sdkmodule.Configurator, k keeper.Keeper) {
	v3types.RegisterMsgServiceServer(configurator.MsgServer(), v3.NewMsgServiceServer(k))

	v3types.RegisterQueryServiceServer(configurator.QueryServer(), v3.NewQueryServiceServer(k))
}
