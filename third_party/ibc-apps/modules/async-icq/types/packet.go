package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (iqpd InterchainQueryPacketData) ValidateBasic() error {
	return nil
}

func (iqpd InterchainQueryPacketData) GetBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&iqpd))
}
