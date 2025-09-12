package v1

import (
	"errors"
	"fmt"
	"strings"

	sdkmath "cosmossdk.io/math"
)

// String converts a RenewalPricePolicy to its string representation.
func (r RenewalPricePolicy) String() string {
	switch r {
	case RenewalPricePolicyIfLesser:
		return "if_lesser"
	case RenewalPricePolicyIfLesserOrEqual:
		return "if_lesser_or_equal"
	case RenewalPricePolicyIfEqual:
		return "if_equal"
	case RenewalPricePolicyIfNotEqual:
		return "if_not_equal"
	case RenewalPricePolicyIfGreater:
		return "if_greater"
	case RenewalPricePolicyIfGreaterOrEqual:
		return "if_greater_or_equal"
	case RenewalPricePolicyAlways:
		return "always"
	default:
		return "unspecified"
	}
}

// Equal checks if two RenewalPricePolicy values are equal.
func (r RenewalPricePolicy) Equal(v RenewalPricePolicy) bool {
	return r == v
}

// IsValid checks whether the RenewalPricePolicy is a valid value.
func (r RenewalPricePolicy) IsValid() bool {
	switch r {
	case RenewalPricePolicyUnspecified,
		RenewalPricePolicyIfLesser,
		RenewalPricePolicyIfLesserOrEqual,
		RenewalPricePolicyIfEqual,
		RenewalPricePolicyIfNotEqual,
		RenewalPricePolicyIfGreater,
		RenewalPricePolicyIfGreaterOrEqual,
		RenewalPricePolicyAlways:
		return true
	default:
		return false
	}
}

// RenewalPricePolicyFromString converts a string to a RenewalPricePolicy.
func RenewalPricePolicyFromString(s string) RenewalPricePolicy {
	s = strings.ToLower(s)
	switch s {
	case "if_lesser":
		return RenewalPricePolicyIfLesser
	case "if_lesser_or_equal":
		return RenewalPricePolicyIfLesserOrEqual
	case "if_equal":
		return RenewalPricePolicyIfEqual
	case "if_not_equal":
		return RenewalPricePolicyIfNotEqual
	case "if_greater":
		return RenewalPricePolicyIfGreater
	case "if_greater_or_equal":
		return RenewalPricePolicyIfGreaterOrEqual
	case "always":
		return RenewalPricePolicyAlways
	default:
		return RenewalPricePolicyUnspecified
	}
}

// Validate validates whether a subscription can be renewed based on the policy and given DecCoin conditions.
// Returns an error if the renewal is not allowed or invalid.
func (r RenewalPricePolicy) validate(curr, prev sdkmath.LegacyDec) error {
	switch r {
	case RenewalPricePolicyUnspecified:
		return errors.New("renewal policy unspecified")
	case RenewalPricePolicyIfLesser:
		if !curr.LT(prev) {
			return fmt.Errorf("current price %s is not less than previous price %s", curr, prev)
		}
	case RenewalPricePolicyIfLesserOrEqual:
		if !curr.LTE(prev) {
			return fmt.Errorf("current price %s is not less than or equal to previous price %s", curr, prev)
		}
	case RenewalPricePolicyIfEqual:
		if !curr.Equal(prev) {
			return fmt.Errorf("current price %s is not equal to previous price %s", curr, prev)
		}
	case RenewalPricePolicyIfNotEqual:
		if curr.Equal(prev) {
			return fmt.Errorf("current price %s is equal to previous price %s", curr, prev)
		}
	case RenewalPricePolicyIfGreater:
		if !curr.GT(prev) {
			return fmt.Errorf("current price %s is not greater than previous price %s", curr, prev)
		}
	case RenewalPricePolicyIfGreaterOrEqual:
		if !curr.GTE(prev) {
			return fmt.Errorf("current price %s is not greater than or equal to previous price %s", curr, prev)
		}
	case RenewalPricePolicyAlways:
		return nil
	default:
		return errors.New("invalid renewal policy")
	}

	return nil
}

func (r RenewalPricePolicy) Validate(curr, prev Price) error {
	if r.Equal(RenewalPricePolicyAlways) {
		return nil
	}

	if curr.Denom != prev.Denom {
		return fmt.Errorf("current price denom %s does not match previous price denom %s", curr.Denom, prev.Denom)
	}

	if prev.BaseValue.IsZero() && prev.QuoteValue.IsZero() && curr.BaseValue.IsZero() && curr.QuoteValue.IsZero() {
		return r.validate(prev.BaseValue, curr.BaseValue)
	}

	if prev.BaseValue.IsZero() && prev.QuoteValue.IsZero() && curr.BaseValue.IsZero() && !curr.QuoteValue.IsZero() {
		return r.validate(prev.QuoteValue.ToLegacyDec(), curr.QuoteValue.ToLegacyDec())
	}

	if prev.BaseValue.IsZero() && prev.QuoteValue.IsZero() && !curr.BaseValue.IsZero() {
		return r.validate(prev.BaseValue, curr.BaseValue)
	}

	if prev.BaseValue.IsZero() && !prev.QuoteValue.IsZero() {
		return r.validate(prev.QuoteValue.ToLegacyDec(), curr.QuoteValue.ToLegacyDec())
	}

	if !prev.BaseValue.IsZero() && curr.BaseValue.IsZero() && curr.QuoteValue.IsZero() {
		return r.validate(prev.BaseValue, curr.BaseValue)
	}

	if !prev.BaseValue.IsZero() && curr.BaseValue.IsZero() && !curr.QuoteValue.IsZero() {
		return r.validate(prev.QuoteValue.ToLegacyDec(), curr.QuoteValue.ToLegacyDec())
	}

	if !prev.BaseValue.IsZero() && !curr.BaseValue.IsZero() {
		return r.validate(prev.BaseValue, curr.BaseValue)
	}

	return nil
}
