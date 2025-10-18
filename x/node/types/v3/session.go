package v3

import (
	"fmt"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/sentinelhub/v12/types"
	sessiontypes "github.com/sentinel-official/sentinelhub/v12/x/session/types/v3"
)

// Ensure that Session implements the sessiontypes.Session interface.
var _ sessiontypes.Session = (*Session)(nil)

// depositAmount calculates the deposit amount for the session.
func (m *Session) depositAmount() sdkmath.Int {
	amount := sdkmath.ZeroInt()

	// Calculate deposit based on max bytes if applicable
	if !m.MaxBytes.IsZero() {
		gigabytes := m.MaxBytes.Quo(base.Gigabyte)
		amount = m.Price.QuoteValue.Mul(gigabytes)
	}

	// Calculate deposit based on max duration if applicable
	if m.MaxDuration != 0 {
		hours := int64(m.MaxDuration / time.Hour)
		amount = m.Price.QuoteValue.MulRaw(hours)
	}

	return amount
}

// paymentAmountForBytes calculates the payment amount based on data usage.
func (m *Session) paymentAmountForBytes() sdkmath.Int {
	decPrice := m.Price.QuoteValue.ToLegacyDec()
	bytePrice := decPrice.QuoInt(base.Gigabyte)      // Price per byte
	totalBytes := m.DownloadBytes.Add(m.UploadBytes) // Total data used

	return bytePrice.MulInt(totalBytes).Ceil().TruncateInt()
}

// paymentAmountForDuration calculates the payment amount based on session duration.
func (m *Session) paymentAmountForDuration() sdkmath.Int {
	decPrice := m.Price.QuoteValue.ToLegacyDec()
	nsPrice := decPrice.QuoInt64(time.Hour.Nanoseconds()) // Price per nanosecond
	nsDuration := m.Duration.Nanoseconds()                // Total duration in nanoseconds

	return nsPrice.MulInt64(nsDuration).Ceil().TruncateInt()
}

// paymentAmount calculates the total payment amount for the session.
func (m *Session) paymentAmount() sdkmath.Int {
	amount := sdkmath.ZeroInt()

	// Calculate payment based on data usage if max bytes is set
	if !m.MaxBytes.IsZero() {
		amount = m.paymentAmountForBytes()
	}

	// Calculate payment based on duration if max duration is set
	if m.MaxDuration != 0 {
		amount = m.paymentAmountForDuration()
	}

	// Ensure the payment amount does not exceed the deposited amount
	deposit := m.depositAmount()
	if amount.GT(deposit) {
		amount = deposit
	}

	return amount
}

// refundAmount calculates the refund amount for the session.
func (m *Session) refundAmount() sdkmath.Int {
	deposit := m.depositAmount()
	payment := m.paymentAmount()

	// Refund is the difference between deposit and payment
	return deposit.Sub(payment)
}

// DepositAmount returns the total deposit amount as an SDK coin.
func (m *Session) DepositAmount() sdk.Coin {
	return sdk.Coin{Denom: m.Price.Denom, Amount: m.depositAmount()}
}

// PaymentAmount returns the total payment amount as an SDK coin.
func (m *Session) PaymentAmount() sdk.Coin {
	return sdk.Coin{Denom: m.Price.Denom, Amount: m.paymentAmount()}
}

// RefundAmount returns the refund amount as an SDK coin.
func (m *Session) RefundAmount() sdk.Coin {
	return sdk.Coin{Denom: m.Price.Denom, Amount: m.refundAmount()}
}

// Validate performs basic validation on the session.
func (m *Session) Validate() error {
	if err := m.BaseSession.Validate(); err != nil {
		return err
	}

	if err := m.Price.Validate(); err != nil {
		return fmt.Errorf("invalid price: %w", err)
	}

	return nil
}
