package v2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = (*MsgStartRequest)(nil)
	_ sdk.Msg = (*MsgUpdateDetailsRequest)(nil)
	_ sdk.Msg = (*MsgEndRequest)(nil)
)

func (m *MsgStartRequest) ValidateBasic() error         { return nil }
func (m *MsgStartRequest) GetSigners() []sdk.AccAddress { return nil }

func (m *MsgUpdateDetailsRequest) ValidateBasic() error         { return nil }
func (m *MsgUpdateDetailsRequest) GetSigners() []sdk.AccAddress { return nil }

func (m *MsgEndRequest) ValidateBasic() error         { return nil }
func (m *MsgEndRequest) GetSigners() []sdk.AccAddress { return nil }
