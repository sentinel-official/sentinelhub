package v2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = (*MsgCreateRequest)(nil)
	_ sdk.Msg = (*MsgUpdateStatusRequest)(nil)
	_ sdk.Msg = (*MsgLinkNodeRequest)(nil)
	_ sdk.Msg = (*MsgUnlinkNodeRequest)(nil)
	_ sdk.Msg = (*MsgSubscribeRequest)(nil)
)

func (m *MsgCreateRequest) ValidateBasic() error         { return nil }
func (m *MsgCreateRequest) GetSigners() []sdk.AccAddress { return nil }

func (m *MsgUpdateStatusRequest) ValidateBasic() error         { return nil }
func (m *MsgUpdateStatusRequest) GetSigners() []sdk.AccAddress { return nil }

func (m *MsgLinkNodeRequest) ValidateBasic() error         { return nil }
func (m *MsgLinkNodeRequest) GetSigners() []sdk.AccAddress { return nil }

func (m *MsgUnlinkNodeRequest) ValidateBasic() error         { return nil }
func (m *MsgUnlinkNodeRequest) GetSigners() []sdk.AccAddress { return nil }

func (m *MsgSubscribeRequest) ValidateBasic() error         { return nil }
func (m *MsgSubscribeRequest) GetSigners() []sdk.AccAddress { return nil }
