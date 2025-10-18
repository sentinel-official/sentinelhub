package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = (*MsgAddRequest)(nil)
	_ sdk.Msg = (*MsgSetStatusRequest)(nil)
	_ sdk.Msg = (*MsgAddNodeRequest)(nil)
	_ sdk.Msg = (*MsgRemoveNodeRequest)(nil)
)

func (m *MsgAddRequest) ValidateBasic() error         { return nil }
func (m *MsgAddRequest) GetSigners() []sdk.AccAddress { return nil }

func (m *MsgSetStatusRequest) ValidateBasic() error         { return nil }
func (m *MsgSetStatusRequest) GetSigners() []sdk.AccAddress { return nil }

func (m *MsgAddNodeRequest) ValidateBasic() error         { return nil }
func (m *MsgAddNodeRequest) GetSigners() []sdk.AccAddress { return nil }

func (m *MsgRemoveNodeRequest) ValidateBasic() error         { return nil }
func (m *MsgRemoveNodeRequest) GetSigners() []sdk.AccAddress { return nil }
