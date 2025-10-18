package v2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = (*MsgRegisterRequest)(nil)
	_ sdk.Msg = (*MsgUpdateDetailsRequest)(nil)
	_ sdk.Msg = (*MsgUpdateStatusRequest)(nil)
	_ sdk.Msg = (*MsgSubscribeRequest)(nil)
)

func (m *MsgRegisterRequest) ValidateBasic() error         { return nil }
func (m *MsgRegisterRequest) GetSigners() []sdk.AccAddress { return nil }

func (m *MsgUpdateDetailsRequest) ValidateBasic() error         { return nil }
func (m *MsgUpdateDetailsRequest) GetSigners() []sdk.AccAddress { return nil }

func (m *MsgUpdateStatusRequest) ValidateBasic() error         { return nil }
func (m *MsgUpdateStatusRequest) GetSigners() []sdk.AccAddress { return nil }

func (m *MsgSubscribeRequest) ValidateBasic() error         { return nil }
func (m *MsgSubscribeRequest) GetSigners() []sdk.AccAddress { return nil }
