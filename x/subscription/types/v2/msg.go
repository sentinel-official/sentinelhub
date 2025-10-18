package v2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = (*MsgCancelRequest)(nil)
	_ sdk.Msg = (*MsgAllocateRequest)(nil)
)

func (m *MsgCancelRequest) ValidateBasic() error         { return nil }
func (m *MsgCancelRequest) GetSigners() []sdk.AccAddress { return nil }

func (m *MsgAllocateRequest) ValidateBasic() error         { return nil }
func (m *MsgAllocateRequest) GetSigners() []sdk.AccAddress { return nil }
