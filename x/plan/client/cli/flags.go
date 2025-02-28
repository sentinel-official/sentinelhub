package cli

import (
	"github.com/spf13/pflag"

	v1base "github.com/sentinel-official/hub/v12/types/v1"
)

const (
	flagPrices = "prices"
)

func GetPrices(flags *pflag.FlagSet) (v1base.Prices, error) {
	s, err := flags.GetString(flagPrices)
	if err != nil {
		return nil, err
	}
	if s == "" {
		return nil, nil
	}

	return v1base.NewPricesFromString(s)
}
