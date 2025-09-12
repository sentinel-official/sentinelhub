package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/pflag"

	"github.com/sentinel-official/sentinelhub/v12/types/v1"
)

const (
	FlagAccAddr        = "acc-addr"
	FlagNodeAddr       = "node-addr"
	FlagPlanID         = "plan-id"
	FlagProvAddr       = "prov-addr"
	FlagStatus         = "status"
	FlagSubscriptionID = "subscription-id"
)

func AccAddrFromFlags(flags *pflag.FlagSet) (sdk.AccAddress, error) {
	s, err := flags.GetString(FlagAccAddr)
	if err != nil {
		return nil, err
	}

	if s == "" {
		return nil, err
	}

	return sdk.AccAddressFromBech32(s)
}

func NodeAddrFromFlags(flags *pflag.FlagSet) (NodeAddress, error) {
	s, err := flags.GetString(FlagNodeAddr)
	if err != nil {
		return nil, err
	}

	if s == "" {
		return nil, err
	}

	return NodeAddressFromBech32(s)
}

func PlanIDFromFlags(flags *pflag.FlagSet) (uint64, error) {
	return flags.GetUint64(FlagPlanID)
}

func ProvAddrFromFlags(flags *pflag.FlagSet) (ProvAddress, error) {
	s, err := flags.GetString(FlagProvAddr)
	if err != nil {
		return nil, err
	}

	if s == "" {
		return nil, err
	}

	return ProvAddressFromBech32(s)
}

func StatusFromFlags(flags *pflag.FlagSet) (v1.Status, error) {
	s, err := flags.GetString(FlagStatus)
	if err != nil {
		return v1.StatusUnspecified, err
	}

	return v1.StatusFromString(s), nil
}

func SubscriptionIDFromFlags(flags *pflag.FlagSet) (uint64, error) {
	return flags.GetUint64(FlagSubscriptionID)
}
