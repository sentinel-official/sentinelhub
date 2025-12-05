package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkbech32 "github.com/cosmos/cosmos-sdk/types/bech32"
	"gopkg.in/yaml.v3"
)

const (
	// Bech32MainPrefix defines the main prefix for account addresses in the blockchain.
	Bech32MainPrefix = "sent"

	// PrefixValidator is the prefix used for validator keys.
	PrefixValidator = "val"

	// PrefixConsensus is the prefix used for consensus keys.
	PrefixConsensus = "cons"

	// PrefixPublic is the prefix used for public keys.
	PrefixPublic = "pub"

	// PrefixOperator is the prefix used for operator keys.
	PrefixOperator = "oper"

	// PrefixProvider is the prefix used for provider keys.
	PrefixProvider = "prov"

	// PrefixNode is the prefix used for node keys.
	PrefixNode = "node"

	// Bech32PrefixAccAddr defines the Bech32 prefix for an account address in the blockchain.
	Bech32PrefixAccAddr = Bech32MainPrefix

	// Bech32PrefixAccPub defines the Bech32 prefix for an account public key in the blockchain.
	Bech32PrefixAccPub = Bech32MainPrefix + PrefixPublic

	// Bech32PrefixValAddr defines the Bech32 prefix for a validator operator address in the blockchain.
	Bech32PrefixValAddr = Bech32MainPrefix + PrefixValidator + PrefixOperator

	// Bech32PrefixValPub defines the Bech32 prefix for a validator operator public key in the blockchain.
	Bech32PrefixValPub = Bech32MainPrefix + PrefixValidator + PrefixOperator + PrefixPublic

	// Bech32PrefixConsAddr defines the Bech32 prefix for a validator consensus address in the blockchain.
	Bech32PrefixConsAddr = Bech32MainPrefix + PrefixValidator + PrefixConsensus

	// Bech32PrefixConsPub defines the Bech32 prefix for a validator consensus public key in the blockchain.
	Bech32PrefixConsPub = Bech32MainPrefix + PrefixValidator + PrefixConsensus + PrefixPublic

	// Bech32PrefixProvAddr defines the Bech32 prefix for a provider address in the blockchain.
	Bech32PrefixProvAddr = Bech32MainPrefix + PrefixProvider

	// Bech32PrefixProvPub defines the Bech32 prefix for a provider public key in the blockchain.
	Bech32PrefixProvPub = Bech32MainPrefix + PrefixProvider + PrefixPublic

	// Bech32PrefixNodeAddr defines the Bech32 prefix for a node address in the blockchain.
	Bech32PrefixNodeAddr = Bech32MainPrefix + PrefixNode

	// Bech32PrefixNodePub defines the Bech32 prefix for a node public key in the blockchain.
	Bech32PrefixNodePub = Bech32MainPrefix + PrefixNode + PrefixPublic
)

var (
	_ sdk.Address    = ProvAddress{}
	_ yaml.Marshaler = ProvAddress{}

	_ sdk.Address    = NodeAddress{}
	_ yaml.Marshaler = NodeAddress{}
)

type ProvAddress []byte

func (p ProvAddress) Equals(addr sdk.Address) bool {
	if p.Empty() && addr == nil {
		return true
	}

	return bytes.Equal(p.Bytes(), addr.Bytes())
}

func (p ProvAddress) Empty() bool {
	return bytes.Equal(p.Bytes(), ProvAddress{}.Bytes())
}

func (p ProvAddress) Bytes() []byte {
	return p
}

func (p ProvAddress) String() string {
	if p.Empty() {
		return ""
	}

	str, err := sdkbech32.ConvertAndEncode(GetConfig().GetBech32ProviderAddrPrefix(), p.Bytes())
	if err != nil {
		panic(err)
	}

	return str
}

func (p ProvAddress) Format(f fmt.State, c rune) {
	switch c {
	case 's':
		_, _ = f.Write([]byte(p.String()))
	case 'p':
		_, _ = fmt.Fprintf(f, "%p", p)
	default:
		_, _ = fmt.Fprintf(f, "%X", p.Bytes())
	}
}

func (p ProvAddress) Marshal() ([]byte, error) {
	return p, nil
}

func (p ProvAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p ProvAddress) MarshalYAML() (any, error) {
	return p.String(), nil
}

func (p *ProvAddress) Unmarshal(data []byte) error {
	*p = data

	return nil
}

func (p *ProvAddress) UnmarshalJSON(data []byte) (err error) {
	var str string
	if err = json.Unmarshal(data, &str); err != nil {
		return err
	}

	addr, err := ProvAddressFromBech32(str)
	if err != nil {
		return err
	}

	*p = addr

	return nil
}

func (p *ProvAddress) UnmarshalYAML(data []byte) (err error) {
	var str string
	if err = yaml.Unmarshal(data, &str); err != nil {
		return err
	}

	addr, err := ProvAddressFromBech32(str)
	if err != nil {
		return err
	}

	*p = addr

	return nil
}

func ProvAddressFromBech32(str string) (ProvAddress, error) {
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return ProvAddress{}, errors.New("empty address string is not allowed")
	}

	buf, err := sdk.GetFromBech32(str, GetConfig().GetBech32ProviderAddrPrefix())
	if err != nil {
		return nil, err
	}

	if err = sdk.VerifyAddressFormat(buf); err != nil {
		return nil, err
	}

	return buf, nil
}

type NodeAddress []byte

func (n NodeAddress) Equals(addr sdk.Address) bool {
	if n.Empty() && addr == nil {
		return true
	}

	return bytes.Equal(n.Bytes(), addr.Bytes())
}

func (n NodeAddress) Empty() bool {
	return bytes.Equal(n.Bytes(), NodeAddress{}.Bytes())
}

func (n NodeAddress) Bytes() []byte {
	return n
}

func (n NodeAddress) String() string {
	if n.Empty() {
		return ""
	}

	str, err := sdkbech32.ConvertAndEncode(GetConfig().GetBech32NodeAddrPrefix(), n.Bytes())
	if err != nil {
		panic(err)
	}

	return str
}

func (n NodeAddress) Format(f fmt.State, c rune) {
	switch c {
	case 's':
		_, _ = f.Write([]byte(n.String()))
	case 'p':
		_, _ = fmt.Fprintf(f, "%p", n)
	default:
		_, _ = fmt.Fprintf(f, "%X", n.Bytes())
	}
}

func (n NodeAddress) Marshal() ([]byte, error) {
	return n, nil
}

func (n NodeAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.String())
}

func (n NodeAddress) MarshalYAML() (any, error) {
	return n.String(), nil
}

func (n *NodeAddress) Unmarshal(data []byte) error {
	*n = data

	return nil
}

func (n *NodeAddress) UnmarshalJSON(data []byte) (err error) {
	var str string
	if err = json.Unmarshal(data, &str); err != nil {
		return err
	}

	addr, err := NodeAddressFromBech32(str)
	if err != nil {
		return err
	}

	*n = addr

	return nil
}

func (n *NodeAddress) UnmarshalYAML(data []byte) (err error) {
	var str string
	if err = yaml.Unmarshal(data, &str); err != nil {
		return err
	}

	addr, err := NodeAddressFromBech32(str)
	if err != nil {
		return err
	}

	*n = addr

	return nil
}

func NodeAddressFromBech32(str string) (NodeAddress, error) {
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return NodeAddress{}, errors.New("empty address string is not allowed")
	}

	buf, err := sdk.GetFromBech32(str, GetConfig().GetBech32NodeAddrPrefix())
	if err != nil {
		return nil, err
	}

	if err = sdk.VerifyAddressFormat(buf); err != nil {
		return nil, err
	}

	return buf, nil
}
