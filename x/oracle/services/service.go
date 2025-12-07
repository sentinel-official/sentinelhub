package services

import (
	sdkmodule "github.com/cosmos/cosmos-sdk/types/module"

	"github.com/sentinel-official/sentinelhub/v13/x/oracle/keeper"
	"github.com/sentinel-official/sentinelhub/v13/x/oracle/services/v1"
	v1types "github.com/sentinel-official/sentinelhub/v13/x/oracle/types/v1"
)

func RegisterServices(configurator sdkmodule.Configurator, k keeper.Keeper) {
	v1types.RegisterMsgServiceServer(configurator.MsgServer(), v1.NewMsgServiceServer(k))

	v1types.RegisterQueryServiceServer(configurator.QueryServer(), v1.NewQueryServiceServer(k))
}
