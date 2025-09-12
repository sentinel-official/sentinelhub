package cli

import (
	"github.com/spf13/pflag"

	v1base "github.com/sentinel-official/sentinelhub/v12/types/v1"
)

const (
	flagDenom          = "denom"
	flagGigabytePrices = "gigabyte-prices"
	flagGigabytes      = "gigabytes"
	flagHourlyPrices   = "hourly-prices"
	flagHours          = "hours"
	flagMaxPrice       = "max-price"
	flagRemoteAddrs    = "remote-addrs"
)

func GetGigabytePrices(flags *pflag.FlagSet) (v1base.Prices, error) {
	s, err := flags.GetString(flagGigabytePrices)
	if err != nil {
		return nil, err
	}

	if s == "" {
		return nil, nil
	}

	return v1base.NewPricesFromString(s)
}

func GetHourlyPrices(flags *pflag.FlagSet) (v1base.Prices, error) {
	s, err := flags.GetString(flagHourlyPrices)
	if err != nil {
		return nil, err
	}

	if s == "" {
		return nil, nil
	}

	return v1base.NewPricesFromString(s)
}

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
