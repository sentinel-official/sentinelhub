package v1

import (
	"errors"
	"fmt"
	"strings"
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

// Validate checks whether the current price satisfies the policy conditions
// when compared to the previous price. Returns an error if the policy is not met
// or if the policy is invalid or unspecified.
func (r RenewalPricePolicy) Validate(curr, prev Price) error {
	if curr.Denom != prev.Denom {
		return fmt.Errorf("price denom mismatch; current=%s, previous=%s", curr.Denom, prev.Denom)
	}

	switch r {
	case RenewalPricePolicyUnspecified:
		return errors.New("renewal price policy is unspecified")
	case RenewalPricePolicyIfLesser:
		if !curr.IsLT(prev) {
			return fmt.Errorf("current price %s is not less than previous price %s", curr, prev)
		}
	case RenewalPricePolicyIfLesserOrEqual:
		if !curr.IsLTE(prev) {
			return fmt.Errorf("current price %s is not less than or equal to previous price %s", curr, prev)
		}
	case RenewalPricePolicyIfEqual:
		if !curr.IsEqual(prev) {
			return fmt.Errorf("current price %s is not equal to previous price %s", curr, prev)
		}
	case RenewalPricePolicyIfNotEqual:
		if curr.IsEqual(prev) {
			return fmt.Errorf("current price %s is equal to previous price %s", curr, prev)
		}
	case RenewalPricePolicyIfGreater:
		if !curr.IsGT(prev) {
			return fmt.Errorf("current price %s is not greater than previous price %s", curr, prev)
		}
	case RenewalPricePolicyIfGreaterOrEqual:
		if !curr.IsGTE(prev) {
			return fmt.Errorf("current price %s is not greater than or equal to previous price %s", curr, prev)
		}
	case RenewalPricePolicyAlways:
		return nil
	default:
		return fmt.Errorf("unsupported renewal price policy %v", r)
	}

	return nil
}
