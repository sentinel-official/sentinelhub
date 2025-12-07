package services

import (
	sdkmodule "github.com/cosmos/cosmos-sdk/types/module"

	"github.com/sentinel-official/sentinelhub/v13/x/swap/keeper"
	"github.com/sentinel-official/sentinelhub/v13/x/swap/services/v1"
	v1types "github.com/sentinel-official/sentinelhub/v13/x/swap/types/v1"
)

func RegisterServices(configurator sdkmodule.Configurator, k keeper.Keeper) {
	v1types.RegisterMsgServiceServer(configurator.MsgServer(), v1.NewMsgServiceServer(k))
	v1types.RegisterQueryServiceServer(configurator.QueryServer(), v1.NewQueryServiceServer(k))
}
