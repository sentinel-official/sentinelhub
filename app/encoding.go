package app

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	standard "github.com/cosmos/cosmos-sdk/std"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
)

type EncodingConfig struct {
	Amino             *codec.LegacyAmino
	Codec             codec.Codec
	InterfaceRegistry codectypes.InterfaceRegistry
	TxConfig          client.TxConfig
}

func NewEncodingConfig() EncodingConfig {
	var (
		amino             = codec.NewLegacyAmino()
		interfaceRegistry = codectypes.NewInterfaceRegistry()
		cdc               = codec.NewProtoCodec(interfaceRegistry)
		txConfig          = authtx.NewTxConfig(cdc, authtx.DefaultSignModes)
	)

	return EncodingConfig{
		Amino:             amino,
		Codec:             cdc,
		InterfaceRegistry: interfaceRegistry,
		TxConfig:          txConfig,
	}
}

func DefaultEncodingConfig() EncodingConfig {
	v := NewEncodingConfig()
	standard.RegisterLegacyAminoCodec(v.Amino)
	standard.RegisterInterfaces(v.InterfaceRegistry)
	ModuleBasics.RegisterLegacyAminoCodec(v.Amino)
	ModuleBasics.RegisterInterfaces(v.InterfaceRegistry)

	return v
}
