package app

import (
	"cosmossdk.io/x/tx/signing"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/address"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/cosmos/gogoproto/proto"
)

type EncodingConfig struct {
	Amino             *codec.LegacyAmino
	Codec             codec.Codec
	InterfaceRegistry codectypes.InterfaceRegistry
	TxConfig          client.TxConfig
}

func NewEncodingConfig() EncodingConfig {
	amino := codec.NewLegacyAmino()

	interfaceRegistry, err := codectypes.NewInterfaceRegistryWithOptions(
		codectypes.InterfaceRegistryOptions{
			ProtoFiles: proto.HybridResolver,
			SigningOptions: signing.Options{
				AddressCodec: address.Bech32Codec{
					Bech32Prefix: sdk.GetConfig().GetBech32AccountAddrPrefix(),
				},
				ValidatorAddressCodec: address.Bech32Codec{
					Bech32Prefix: sdk.GetConfig().GetBech32ValidatorAddrPrefix(),
				},
			},
		},
	)
	if err != nil {
		panic(err)
	}

	cdc := codec.NewProtoCodec(interfaceRegistry)
	txConfig := authtx.NewTxConfig(cdc, authtx.DefaultSignModes)

	return EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Codec:             cdc,
		TxConfig:          txConfig,
		Amino:             amino,
	}
}

func DefaultEncodingConfig() EncodingConfig {
	v := NewEncodingConfig()
	std.RegisterLegacyAminoCodec(v.Amino)
	std.RegisterInterfaces(v.InterfaceRegistry)

	return v
}
