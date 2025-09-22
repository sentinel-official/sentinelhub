package v1

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/sentinelhub/v12/x/swap/types"
)

func (m *Swap) GetTxHash() (hash types.EthereumHash) {
	return types.BytesToHash(m.TxHash)
}

func (m *Swap) Validate() error {
	if m.TxHash == nil {
		return errors.New("tx_hash cannot be nil")
	}

	if len(m.TxHash) == 0 {
		return errors.New("tx_hash cannot be empty")
	}

	if len(m.TxHash) < types.EthereumHashLength {
		return fmt.Errorf("tx_hash length cannot be less than %d", types.EthereumHashLength)
	}

	if len(m.TxHash) > types.EthereumHashLength {
		return fmt.Errorf("tx_hash length cannot be greater than %d", types.EthereumHashLength)
	}

	if m.Receiver == "" {
		return errors.New("receiver cannot be empty")
	}

	if _, err := sdk.AccAddressFromBech32(m.Receiver); err != nil {
		return fmt.Errorf("invalid receiver %s: %w", m.Receiver, err)
	}

	if m.Amount.IsNegative() {
		return errors.New("amount cannot be negative")
	}

	if m.Amount.IsZero() {
		return errors.New("amount cannot be zero")
	}

	if m.Amount.Amount.LT(types.PrecisionLoss) {
		return fmt.Errorf("amount cannot be less than %s", types.PrecisionLoss)
	}

	if !m.Amount.IsValid() {
		return errors.New("amount must be valid")
	}

	return nil
}

type (
	Swaps []Swap
)
