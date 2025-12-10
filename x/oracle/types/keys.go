package types

import (
	ibcicqtypes "github.com/sentinel-official/sentinelhub/v13/third_party/ibc-apps/modules/async-icq/types"
)

const (
	ModuleName = "oracle"
	PortID     = "oracle-1"
	StoreKey   = ModuleName
	Version    = ibcicqtypes.Version
)

var (
	ParamsKey = []byte{0x00}
	PortIDKey = []byte{0x01}

	AssetKeyPrefix = []byte{0x10}
)

func AssetKey(denom string) []byte {
	return append(AssetKeyPrefix, []byte(denom)...)
}
