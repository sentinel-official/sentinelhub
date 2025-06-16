package v1

import (
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/sentinelhub/v12/x/swap/types"
)

func (m *Swap) GetTxHash() (hash types.EthereumHash) {
	return types.BytesToHash(m.TxHash)
}

func (m *Swap) Validate() error {
	if m.TxHash == nil {
		return fmt.Errorf("tx_hash cannot be nil")
	}
	if len(m.TxHash) == 0 {
		return fmt.Errorf("tx_hash cannot be empty")
	}
	if len(m.TxHash) < types.EthereumHashLength {
		return fmt.Errorf("tx_hash length cannot be less than %d", types.EthereumHashLength)
	}
	if len(m.TxHash) > types.EthereumHashLength {
		return fmt.Errorf("tx_hash length cannot be greater than %d", types.EthereumHashLength)
	}
	if m.Receiver == "" {
		return fmt.Errorf("receiver cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.Receiver); err != nil {
		return sdkerrors.Wrapf(err, "invalid receiver %s", m.Receiver)
	}
	if m.Amount.IsNegative() {
		return fmt.Errorf("amount cannot be negative")
	}
	if m.Amount.IsZero() {
		return fmt.Errorf("amount cannot be zero")
	}
	if m.Amount.Amount.LT(types.PrecisionLoss) {
		return fmt.Errorf("amount cannot be less than %s", types.PrecisionLoss)
	}
	if !m.Amount.IsValid() {
		return fmt.Errorf("amount must be valid")
	}

	return nil
}

type (
	Swaps []Swap
)
