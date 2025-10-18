package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = (*MsgStartRequest)(nil)
	_ sdk.Msg = (*MsgUpdateRequest)(nil)
	_ sdk.Msg = (*MsgEndRequest)(nil)
)

func (m *MsgStartRequest) ValidateBasic() error         { return nil }
func (m *MsgStartRequest) GetSigners() []sdk.AccAddress { return nil }

func (m *MsgUpdateRequest) ValidateBasic() error         { return nil }
func (m *MsgUpdateRequest) GetSigners() []sdk.AccAddress { return nil }

func (m *MsgEndRequest) ValidateBasic() error         { return nil }
func (m *MsgEndRequest) GetSigners() []sdk.AccAddress { return nil }
