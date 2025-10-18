package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = (*MsgRegisterRequest)(nil)
	_ sdk.Msg = (*MsgUpdateRequest)(nil)
)

func (m *MsgRegisterRequest) ValidateBasic() error         { return nil }
func (m *MsgRegisterRequest) GetSigners() []sdk.AccAddress { return nil }

func (m *MsgUpdateRequest) ValidateBasic() error         { return nil }
func (m *MsgUpdateRequest) GetSigners() []sdk.AccAddress { return nil }
