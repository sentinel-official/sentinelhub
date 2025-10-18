package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = (*MsgSubscribeToNodeRequest)(nil)
	_ sdk.Msg = (*MsgSubscribeToPlanRequest)(nil)
	_ sdk.Msg = (*MsgCancelRequest)(nil)
	_ sdk.Msg = (*MsgAddQuotaRequest)(nil)
	_ sdk.Msg = (*MsgUpdateQuotaRequest)(nil)
)

func (m *MsgSubscribeToNodeRequest) ValidateBasic() error         { return nil }
func (m *MsgSubscribeToNodeRequest) GetSigners() []sdk.AccAddress { return nil }

func (m *MsgSubscribeToPlanRequest) ValidateBasic() error         { return nil }
func (m *MsgSubscribeToPlanRequest) GetSigners() []sdk.AccAddress { return nil }

func (m *MsgCancelRequest) ValidateBasic() error         { return nil }
func (m *MsgCancelRequest) GetSigners() []sdk.AccAddress { return nil }

func (m *MsgAddQuotaRequest) ValidateBasic() error         { return nil }
func (m *MsgAddQuotaRequest) GetSigners() []sdk.AccAddress { return nil }

func (m *MsgUpdateQuotaRequest) ValidateBasic() error         { return nil }
func (m *MsgUpdateQuotaRequest) GetSigners() []sdk.AccAddress { return nil }
