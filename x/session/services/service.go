package services

import (
	sdkmodule "github.com/cosmos/cosmos-sdk/types/module"

	"github.com/sentinel-official/sentinelhub/v13/x/session/keeper"
	"github.com/sentinel-official/sentinelhub/v13/x/session/services/v3"
	v3types "github.com/sentinel-official/sentinelhub/v13/x/session/types/v3"
)

func RegisterServices(configurator sdkmodule.Configurator, k keeper.Keeper) {
	v3types.RegisterMsgServiceServer(configurator.MsgServer(), v3.NewMsgServiceServer(k))

	v3types.RegisterQueryServiceServer(configurator.QueryServer(), v3.NewQueryServiceServer(k))
}
