package v1

import (
	"errors"
	"time"
)

// Default parameter values for the Params struct
var (
	DefaultBlockInterval int64 = 10              // Default block interval in block numbers
	DefaultChannelID           = ""              // Default IBC channel ID
	DefaultPortID              = ""              // Default IBC port ID
	DefaultTimeout             = 1 * time.Minute // Default timeout duration
)

// Validate checks whether the Params fields are valid according to defined rules.
func (m *Params) Validate() error {
	if err := validateBlockInterval(m.BlockInterval); err != nil {
		return err
	}
	if err := validateChannelID(m.ChannelID); err != nil {
		return err
	}
	if err := validatePortID(m.PortID); err != nil {
		return err
	}
	if err := validateTimeout(m.Timeout); err != nil {
		return err
	}

	return nil
}

// NewParams creates a new Params instance with custom values.
func NewParams(blockInterval int64, channelID, portID string, timeout time.Duration) Params {
	return Params{
		BlockInterval: blockInterval,
		ChannelID:     channelID,
		PortID:        portID,
		Timeout:       timeout,
	}
}

// DefaultParams returns a Params struct initialized with default values.
func DefaultParams() Params {
	return NewParams(
		DefaultBlockInterval,
		DefaultChannelID,
		DefaultPortID,
		DefaultTimeout,
	)
}

// validateBlockInterval checks that the block interval is a positive value.
func validateBlockInterval(v int64) error {
	if v < 0 {
		return errors.New("block_interval cannot be negative")
	}
	if v == 0 {
		return errors.New("block_interval cannot not be zero")
	}

	return nil
}

// validateChannelID checks that the channel ID is non-empty.
func validateChannelID(_ string) error {
	return nil
}

// validatePortID checks that the port ID is non-empty.
func validatePortID(_ string) error {
	return nil
}

// validateTimeout ensures the timeout duration is positive.
func validateTimeout(v time.Duration) error {
	if v < 0 {
		return errors.New("timeout cannot be negative")
	}
	if v == 0 {
		return errors.New("timeout cannot be zero")
	}

	return nil
}
