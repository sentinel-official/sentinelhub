package v1

import (
	"errors"
	"fmt"
	"time"
)

// Default parameter values for the Params struct.
var (
	DefaultBlockInterval int64 = 100             // Default block interval in block numbers
	DefaultChannelID           = ""              // Default IBC channel ID
	DefaultTimeout             = 1 * time.Minute // Default timeout duration
)

// Validate checks whether the Params fields are valid according to defined rules.
func (m *Params) Validate() error {
	if err := validateBlockInterval(m.BlockInterval); err != nil {
		return fmt.Errorf("invalid block_interval: %w", err)
	}

	if err := validateChannelID(m.ChannelID); err != nil {
		return fmt.Errorf("invalid channel_id: %w", err)
	}

	if err := validateTimeout(m.Timeout); err != nil {
		return fmt.Errorf("invalid timeout: %w", err)
	}

	return nil
}

// NewParams creates a new Params instance with custom values.
func NewParams(blockInterval int64, channelID string, timeout time.Duration) Params {
	return Params{
		BlockInterval: blockInterval,
		ChannelID:     channelID,
		Timeout:       timeout,
	}
}

// DefaultParams returns a Params struct initialized with default values.
func DefaultParams() Params {
	return NewParams(
		DefaultBlockInterval,
		DefaultChannelID,
		DefaultTimeout,
	)
}

// validateBlockInterval checks that the block interval is a positive value.
func validateBlockInterval(v int64) error {
	if v == 0 {
		return errors.New("value cannot not be zero")
	}

	if v < 0 {
		return errors.New("value cannot be negative")
	}

	return nil
}

// validateChannelID checks that the channel ID is non-empty.
func validateChannelID(_ string) error {
	return nil
}

// validateTimeout ensures the timeout duration is positive.
func validateTimeout(v time.Duration) error {
	if v == 0 {
		return errors.New("value cannot be zero")
	}

	if v < 0 {
		return errors.New("value cannot be negative")
	}

	return nil
}
