package mint

import (
	"encoding/json"

	abcitypes "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkmodule "github.com/cosmos/cosmos-sdk/types/module"
	sdksimulation "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/sentinelhub/v12/x/mint/keeper"
	"github.com/sentinel-official/sentinelhub/v12/x/mint/types"
	"github.com/sentinel-official/sentinelhub/v12/x/mint/types/v1"
)

var (
	_ sdkmodule.AppModuleBasic   = AppModuleBasic{}
	_ sdkmodule.HasGenesisBasics = AppModuleBasic{}

	_ sdkmodule.AppModuleGenesis    = AppModule{}
	_ sdkmodule.AppModuleSimulation = AppModule{}
	_ sdkmodule.BeginBlockAppModule = AppModule{}
	_ sdkmodule.EndBlockAppModule   = AppModule{}
	_ sdkmodule.HasConsensusVersion = AppModule{}
	_ sdkmodule.HasInvariants       = AppModule{}
	_ sdkmodule.HasServices         = AppModule{}
)

type AppModuleBasic struct{}

func (amb AppModuleBasic) Name() string { return types.ModuleName }

func (amb AppModuleBasic) RegisterLegacyAminoCodec(_ *codec.LegacyAmino) {}

func (amb AppModuleBasic) RegisterInterfaces(_ codectypes.InterfaceRegistry) {}

func (amb AppModuleBasic) RegisterGRPCGatewayRoutes(_ client.Context, _ *runtime.ServeMux) {}

func (amb AppModuleBasic) GetTxCmd() *cobra.Command { return nil }

func (amb AppModuleBasic) GetQueryCmd() *cobra.Command { return nil }

func (amb AppModuleBasic) DefaultGenesis(jsonCodec codec.JSONCodec) json.RawMessage {
	state := v1.DefaultGenesisState()

	return jsonCodec.MustMarshalJSON(state)
}

func (amb AppModuleBasic) ValidateGenesis(jsonCodec codec.JSONCodec, _ client.TxEncodingConfig, message json.RawMessage) error {
	var state v1.GenesisState
	if err := jsonCodec.UnmarshalJSON(message, &state); err != nil {
		return err
	}

	return state.Validate()
}

type AppModule struct {
	AppModuleBasic

	cdc    codec.Codec
	keeper keeper.Keeper
}

func NewAppModule(cdc codec.Codec, k keeper.Keeper) AppModule {
	return AppModule{
		cdc:    cdc,
		keeper: k,
	}
}

func (am AppModule) InitGenesis(ctx sdk.Context, jsonCodec codec.JSONCodec, message json.RawMessage) []abcitypes.ValidatorUpdate {
	var state v1.GenesisState
	jsonCodec.MustUnmarshalJSON(message, &state)
	am.keeper.InitGenesis(ctx, &state)

	return nil
}

func (am AppModule) ExportGenesis(ctx sdk.Context, jsonCodec codec.JSONCodec) json.RawMessage {
	state := am.keeper.ExportGenesis(ctx)

	return jsonCodec.MustMarshalJSON(state)
}

func (am AppModule) BeginBlock(ctx sdk.Context, _ abcitypes.RequestBeginBlock) {
	am.keeper.BeginBlock(ctx)
}

func (am AppModule) EndBlock(_ sdk.Context, _ abcitypes.RequestEndBlock) []abcitypes.ValidatorUpdate {
	return nil
}

func (am AppModule) GenerateGenesisState(_ *sdkmodule.SimulationState) {}

func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

func (am AppModule) WeightedOperations(_ sdkmodule.SimulationState) []sdksimulation.WeightedOperation {
	return nil
}

func (am AppModule) ConsensusVersion() uint64 { return 1 }

func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

func (am AppModule) RegisterServices(_ sdkmodule.Configurator) {}
