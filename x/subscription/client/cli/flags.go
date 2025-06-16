package cli

import (
	"github.com/spf13/pflag"

	v1base "github.com/sentinel-official/sentinelhub/v12/types/v1"
)

const (
	flagDenom              = "denom"
	flagRenewalPricePolicy = "renewal-price-policy"
)

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
