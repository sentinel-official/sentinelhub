package cli

import (
	"github.com/spf13/pflag"

	v1base "github.com/sentinel-official/sentinelhub/v12/types/v1"
)

const (
	flagMaxPrice           = "max-price"
	flagRenewalPricePolicy = "renewal-price-policy"
)

func GetMaxPrice(flags *pflag.FlagSet) (v1base.Price, error) {
	s, err := flags.GetString(flagMaxPrice)
	if err != nil {
		return v1base.Price{}, err
	}
	if s == "" {
		return v1base.ZeroPrice(""), nil
	}

	return v1base.NewPriceFromString(s)
}

func GetRenewalPricePolicy(flags *pflag.FlagSet) (v1base.RenewalPricePolicy, error) {
	s, err := flags.GetString(flagRenewalPricePolicy)
	if err != nil {
		return v1base.RenewalPricePolicyUnspecified, err
	}
	if s == "" {
		return v1base.RenewalPricePolicyUnspecified, nil
	}

	return v1base.RenewalPricePolicyFromString(s), nil
}
