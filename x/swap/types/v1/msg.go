package v1

import (
	sdkerrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/sentinelhub/v12/x/swap/types"
)

var (
	_ sdk.Msg = (*MsgSwapRequest)(nil)
)

func NewMsgSwapRequest(from sdk.AccAddress, txHash types.EthereumHash, receiver sdk.AccAddress, amount sdkmath.Int) *MsgSwapRequest {
	return &MsgSwapRequest{
		From:     from.String(),
		TxHash:   txHash.Bytes(),
		Receiver: receiver.String(),
		Amount:   amount,
	}
}

func (m *MsgSwapRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrorInvalidFrom, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrapf(types.ErrorInvalidFrom, "%s", err)
	}
	if m.Receiver == "" {
		return sdkerrors.Wrap(types.ErrorInvalidReceiver, "receiver cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.Receiver); err != nil {
		return sdkerrors.Wrapf(types.ErrorInvalidReceiver, "%s", err)
	}
	if m.TxHash == nil {
		return sdkerrors.Wrap(types.ErrorInvalidTxHash, "tx_hash cannot be nil")
	}
	if len(m.TxHash) == 0 {
		return sdkerrors.Wrap(types.ErrorInvalidTxHash, "tx_hash cannot be empty")
	}
	if len(m.TxHash) < types.EthereumHashLength {
		return sdkerrors.Wrapf(types.ErrorInvalidTxHash, "tx_hash length cannot be less than %d", types.EthereumHashLength)
	}
	if len(m.TxHash) > types.EthereumHashLength {
		return sdkerrors.Wrapf(types.ErrorInvalidTxHash, "tx_hash length cannot be greater than %d", types.EthereumHashLength)
	}
	if m.Amount.IsNegative() {
		return sdkerrors.Wrap(types.ErrorInvalidAmount, "amount cannot be negative")
	}
	if m.Amount.IsZero() {
		return sdkerrors.Wrap(types.ErrorInvalidAmount, "amount cannot be zero")
	}
	if m.Amount.LT(types.PrecisionLoss) {
		return sdkerrors.Wrapf(types.ErrorInvalidAmount, "amount cannot be less than %s", types.PrecisionLoss)
	}

	return nil
}

func (m *MsgSwapRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}
