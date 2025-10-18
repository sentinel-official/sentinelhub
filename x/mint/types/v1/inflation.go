package v1

import (
	"errors"

	sdkmath "cosmossdk.io/math"
)

func (i *Inflation) Validate() error {
	if i.Max.IsNegative() {
		return errors.New("max cannot be negative")
	}

	if i.Max.GT(sdkmath.LegacyOneDec()) {
		return errors.New("max cannot be greater than one")
	}

	if i.Min.IsNegative() {
		return errors.New("min cannot be negative")
	}

	if i.Min.GT(sdkmath.LegacyOneDec()) {
		return errors.New("min cannot be greater than one")
	}

	if i.Min.GT(i.Max) {
		return errors.New("min cannot be greater than max")
	}

	if i.RateChange.IsNegative() {
		return errors.New("rate_change cannot be negative")
	}

	if i.RateChange.GT(sdkmath.LegacyOneDec()) {
		return errors.New("rate_change cannot be greater than one")
	}

	if i.Timestamp.IsZero() {
		return errors.New("timestamp cannot be zero")
	}

	return nil
}
