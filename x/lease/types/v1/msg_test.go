package v1

import (
	"math"
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/stretchr/testify/require"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
)

func TestMsgEndLeaseRequest_ValidateBasic(t *testing.T) {
	tests := []struct {
		name   string
		msg    *MsgEndLeaseRequest
		expErr bool
	}{
		{
			"ValidMessage",
			&MsgEndLeaseRequest{From: base.TestBech32ProvAddr20Bytes, ID: 1},
			false,
		},
		{
			"EmptyFromAddress",
			&MsgEndLeaseRequest{From: base.TestAddrEmpty, ID: 1},
			true,
		},
		{
			"InvalidIDZero",
			&MsgEndLeaseRequest{From: base.TestBech32ProvAddr20Bytes, ID: 0},
			true,
		},
		{
			"InvalidFromAddressFormat",
			&MsgEndLeaseRequest{From: base.TestAddrInvalid, ID: 1},
			true,
		},
		{
			"InvalidFromAddressPrefix",
			&MsgEndLeaseRequest{From: base.TestAddrInvalidPrefix, ID: 1},
			true,
		},
		{
			"MaximumUint64ID",
			&MsgEndLeaseRequest{From: base.TestBech32ProvAddr20Bytes, ID: math.MaxUint64},
			false,
		},
		{
			"ValidFromAddress10Bytes",
			&MsgEndLeaseRequest{From: base.TestBech32ProvAddr10Bytes, ID: 123},
			false,
		},
		{
			"ValidFromAddress30Bytes",
			&MsgEndLeaseRequest{From: base.TestBech32ProvAddr30Bytes, ID: 123},
			false,
		},
		{
			"InvalidFromAddressAccount",
			&MsgEndLeaseRequest{From: base.TestBech32AccAddr20Bytes, ID: 123},
			true,
		},
		{
			"InvalidFromAddressNode",
			&MsgEndLeaseRequest{From: base.TestBech32NodeAddr20Bytes, ID: 123},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgRenewLeaseRequest_ValidateBasic(t *testing.T) {
	tests := []struct {
		name   string
		msg    *MsgRenewLeaseRequest
		expErr bool
	}{
		{
			"ValidMessage",
			&MsgRenewLeaseRequest{From: base.TestBech32ProvAddr20Bytes, ID: 1, Hours: 24, Denom: base.TestDenomOne},
			false,
		},
		{
			"EmptyFromAddress",
			&MsgRenewLeaseRequest{From: base.TestAddrEmpty, ID: 1, Hours: 24, Denom: base.TestDenomOne},
			true,
		},
		{
			"InvalidIDZero",
			&MsgRenewLeaseRequest{From: base.TestBech32ProvAddr20Bytes, ID: 0, Hours: 24, Denom: base.TestDenomOne},
			true,
		},
		{
			"HoursCannotBeZero",
			&MsgRenewLeaseRequest{From: base.TestBech32ProvAddr20Bytes, ID: 1, Hours: 0, Denom: base.TestDenomOne},
			true,
		},
		{
			"NegativeHours",
			&MsgRenewLeaseRequest{From: base.TestBech32ProvAddr20Bytes, ID: 1, Hours: -5, Denom: base.TestDenomOne},
			true,
		},
		{
			"EmptyDenom",
			&MsgRenewLeaseRequest{From: base.TestBech32ProvAddr20Bytes, ID: 1, Hours: 24, Denom: base.TestDenomEmpty},
			false,
		},
		{
			"InvalidDenom",
			&MsgRenewLeaseRequest{From: base.TestBech32ProvAddr20Bytes, ID: 1, Hours: 24, Denom: base.TestDenomInvalid},
			true,
		},
		{
			"ValidFromAddress10Bytes",
			&MsgRenewLeaseRequest{From: base.TestBech32ProvAddr10Bytes, ID: 1, Hours: 24, Denom: base.TestDenomOne},
			false,
		},
		{
			"ValidFromAddress30Bytes",
			&MsgRenewLeaseRequest{From: base.TestBech32ProvAddr30Bytes, ID: 1, Hours: 24, Denom: base.TestDenomOne},
			false,
		},
		{
			"MaximumUint64ID",
			&MsgRenewLeaseRequest{From: base.TestBech32ProvAddr20Bytes, ID: math.MaxUint64, Hours: 24, Denom: base.TestDenomOne},
			false,
		},
		{
			"InvalidFromAddressAccount",
			&MsgRenewLeaseRequest{From: base.TestBech32AccAddr20Bytes, ID: 1, Hours: 24, Denom: base.TestDenomOne},
			true,
		},
		{
			"InvalidFromAddressNode",
			&MsgRenewLeaseRequest{From: base.TestBech32NodeAddr20Bytes, ID: 1, Hours: 24, Denom: base.TestDenomOne},
			true,
		},
		{
			"ValidMessageWithLargeHours",
			&MsgRenewLeaseRequest{From: base.TestBech32ProvAddr20Bytes, ID: 1, Hours: 1000, Denom: base.TestDenomOne},
			false,
		},
		{
			"InvalidAddressFormatAndZeroID",
			&MsgRenewLeaseRequest{From: base.TestAddrInvalid, ID: 0, Hours: 24, Denom: base.TestDenomOne},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgStartLeaseRequest_ValidateBasic(t *testing.T) {
	tests := []struct {
		name   string
		msg    *MsgStartLeaseRequest
		expErr bool
	}{
		{
			"ValidMessage",
			&MsgStartLeaseRequest{From: base.TestBech32ProvAddr20Bytes, NodeAddress: base.TestBech32NodeAddr20Bytes, Hours: 24, Denom: base.TestDenomOne, RenewalPricePolicy: v1base.RenewalPricePolicyAlways},
			false,
		},
		{
			"EmptyFromAddress",
			&MsgStartLeaseRequest{From: base.TestAddrEmpty, NodeAddress: base.TestBech32NodeAddr20Bytes, Hours: 24, Denom: base.TestDenomOne, RenewalPricePolicy: v1base.RenewalPricePolicyAlways},
			true,
		},
		{
			"EmptyNodeAddress",
			&MsgStartLeaseRequest{From: base.TestBech32ProvAddr20Bytes, NodeAddress: base.TestAddrEmpty, Hours: 24, Denom: base.TestDenomOne, RenewalPricePolicy: v1base.RenewalPricePolicyAlways},
			true,
		},
		{
			"InvalidNodeAddressFormat",
			&MsgStartLeaseRequest{From: base.TestBech32ProvAddr20Bytes, NodeAddress: base.TestAddrInvalid, Hours: 24, Denom: base.TestDenomOne, RenewalPricePolicy: v1base.RenewalPricePolicyAlways},
			true,
		},
		{
			"HoursCannotBeZero",
			&MsgStartLeaseRequest{From: base.TestBech32ProvAddr20Bytes, NodeAddress: base.TestBech32NodeAddr20Bytes, Hours: 0, Denom: base.TestDenomOne, RenewalPricePolicy: v1base.RenewalPricePolicyAlways},
			true,
		},
		{
			"NegativeHours",
			&MsgStartLeaseRequest{From: base.TestBech32ProvAddr20Bytes, NodeAddress: base.TestBech32NodeAddr20Bytes, Hours: -5, Denom: base.TestDenomOne, RenewalPricePolicy: v1base.RenewalPricePolicyAlways},
			true,
		},
		{
			"EmptyDenom",
			&MsgStartLeaseRequest{From: base.TestBech32ProvAddr20Bytes, NodeAddress: base.TestBech32NodeAddr20Bytes, Hours: 24, Denom: base.TestDenomEmpty, RenewalPricePolicy: v1base.RenewalPricePolicyAlways},
			false,
		},
		{
			"InvalidDenom",
			&MsgStartLeaseRequest{From: base.TestBech32ProvAddr20Bytes, NodeAddress: base.TestBech32NodeAddr20Bytes, Hours: 24, Denom: base.TestDenomInvalid, RenewalPricePolicy: v1base.RenewalPricePolicyAlways},
			true,
		},
		{
			"ValidFromAddress10Bytes",
			&MsgStartLeaseRequest{From: base.TestBech32ProvAddr10Bytes, NodeAddress: base.TestBech32NodeAddr20Bytes, Hours: 24, Denom: base.TestDenomOne, RenewalPricePolicy: v1base.RenewalPricePolicyAlways},
			false,
		},
		{
			"ValidNodeAddress10Bytes",
			&MsgStartLeaseRequest{From: base.TestBech32ProvAddr20Bytes, NodeAddress: base.TestBech32NodeAddr10Bytes, Hours: 24, Denom: base.TestDenomOne, RenewalPricePolicy: v1base.RenewalPricePolicyAlways},
			false,
		},
		{
			"ValidFromAddress30Bytes",
			&MsgStartLeaseRequest{From: base.TestBech32ProvAddr30Bytes, NodeAddress: base.TestBech32NodeAddr20Bytes, Hours: 24, Denom: base.TestDenomOne, RenewalPricePolicy: v1base.RenewalPricePolicyAlways},
			false,
		},
		{
			"MaximumUint64Hours",
			&MsgStartLeaseRequest{From: base.TestBech32ProvAddr20Bytes, NodeAddress: base.TestBech32NodeAddr20Bytes, Hours: math.MaxInt64, Denom: base.TestDenomOne, RenewalPricePolicy: v1base.RenewalPricePolicyAlways},
			false,
		},
		{
			"InvalidFromAddressAccount",
			&MsgStartLeaseRequest{From: base.TestBech32AccAddr20Bytes, NodeAddress: base.TestBech32NodeAddr20Bytes, Hours: 24, Denom: base.TestDenomOne, RenewalPricePolicy: v1base.RenewalPricePolicyAlways},
			true,
		},
		{
			"InvalidFromAddressEmptyNode",
			&MsgStartLeaseRequest{From: base.TestBech32ProvAddr20Bytes, NodeAddress: base.TestAddrEmpty, Hours: 24, Denom: base.TestDenomOne, RenewalPricePolicy: v1base.RenewalPricePolicyAlways},
			true,
		},
		{
			"LargeValidHours",
			&MsgStartLeaseRequest{From: base.TestBech32ProvAddr20Bytes, NodeAddress: base.TestBech32NodeAddr20Bytes, Hours: 100000, Denom: base.TestDenomOne, RenewalPricePolicy: v1base.RenewalPricePolicyAlways},
			false,
		},
		{
			"InvalidAddressFormatAndZeroHours",
			&MsgStartLeaseRequest{From: base.TestAddrInvalid, NodeAddress: base.TestBech32NodeAddr20Bytes, Hours: 0, Denom: base.TestDenomOne, RenewalPricePolicy: v1base.RenewalPricePolicyAlways},
			true,
		},
		{
			"InvalidFromAndNodeAddress",
			&MsgStartLeaseRequest{From: base.TestAddrInvalid, NodeAddress: base.TestAddrInvalidPrefix, Hours: 24, Denom: base.TestDenomOne, RenewalPricePolicy: v1base.RenewalPricePolicyAlways},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgUpdateLeaseRequest_ValidateBasic(t *testing.T) {
	tests := []struct {
		name   string
		msg    *MsgUpdateLeaseRequest
		expErr bool
	}{
		{
			"ValidMessage",
			&MsgUpdateLeaseRequest{From: base.TestBech32ProvAddr20Bytes, ID: 1, RenewalPricePolicy: v1base.RenewalPricePolicyAlways},
			false,
		},
		{
			"EmptyFromAddress",
			&MsgUpdateLeaseRequest{From: base.TestAddrEmpty, ID: 1, RenewalPricePolicy: v1base.RenewalPricePolicyAlways},
			true,
		},
		{
			"InvalidIDZero",
			&MsgUpdateLeaseRequest{From: base.TestBech32ProvAddr20Bytes, ID: 0, RenewalPricePolicy: v1base.RenewalPricePolicyAlways},
			true,
		},
		{
			"InvalidFromAddressFormat",
			&MsgUpdateLeaseRequest{From: base.TestAddrInvalid, ID: 1, RenewalPricePolicy: v1base.RenewalPricePolicyAlways},
			true,
		},
		{
			"InvalidFromAddressPrefix",
			&MsgUpdateLeaseRequest{From: base.TestAddrInvalidPrefix, ID: 1, RenewalPricePolicy: v1base.RenewalPricePolicyAlways},
			true,
		},
		{
			"ValidFromAddress10Bytes",
			&MsgUpdateLeaseRequest{From: base.TestBech32ProvAddr10Bytes, ID: 1, RenewalPricePolicy: v1base.RenewalPricePolicyAlways},
			false,
		},
		{
			"ValidFromAddress30Bytes",
			&MsgUpdateLeaseRequest{From: base.TestBech32ProvAddr30Bytes, ID: 1, RenewalPricePolicy: v1base.RenewalPricePolicyAlways},
			false,
		},
		{
			"MaximumUint64ID",
			&MsgUpdateLeaseRequest{From: base.TestBech32ProvAddr20Bytes, ID: math.MaxUint64, RenewalPricePolicy: v1base.RenewalPricePolicyAlways},
			false,
		},
		{
			"InvalidFromAddressAccount",
			&MsgUpdateLeaseRequest{From: base.TestBech32AccAddr20Bytes, ID: 1, RenewalPricePolicy: v1base.RenewalPricePolicyAlways},
			true,
		},
		{
			"InvalidFromAddressNode",
			&MsgUpdateLeaseRequest{From: base.TestBech32NodeAddr20Bytes, ID: 1, RenewalPricePolicy: v1base.RenewalPricePolicyAlways},
			true,
		},
		{
			"ValidFromAddressNonRenewable",
			&MsgUpdateLeaseRequest{From: base.TestBech32ProvAddr20Bytes, ID: 1, RenewalPricePolicy: v1base.RenewalPricePolicyUnspecified},
			false,
		},
		{
			"ValidFromAddressSmallValidID",
			&MsgUpdateLeaseRequest{From: base.TestBech32ProvAddr20Bytes, ID: 2, RenewalPricePolicy: v1base.RenewalPricePolicyAlways},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgUpdateParamsRequest_ValidateBasic(t *testing.T) {
	tests := []struct {
		name   string
		msg    *MsgUpdateParamsRequest
		expErr bool
	}{
		{
			"ValidMessage",
			&MsgUpdateParamsRequest{From: base.TestBech32AccAddr20Bytes, Params: DefaultParams()},
			false,
		},
		{
			"EmptyFromAddress",
			&MsgUpdateParamsRequest{From: base.TestAddrEmpty, Params: DefaultParams()},
			true,
		},
		{
			"InvalidFromAddressFormat",
			&MsgUpdateParamsRequest{From: base.TestAddrInvalid, Params: DefaultParams()},
			true,
		},
		{
			"InvalidFromAddressPrefix",
			&MsgUpdateParamsRequest{From: base.TestAddrInvalidPrefix, Params: DefaultParams()},
			true,
		},
		{
			"EmptyParams",
			&MsgUpdateParamsRequest{From: base.TestBech32AccAddr20Bytes, Params: Params{}},
			true,
		},
		{
			"InvalidMaxHoursNegative",
			&MsgUpdateParamsRequest{From: base.TestBech32AccAddr20Bytes, Params: Params{MaxHours: -1}},
			true,
		},
		{
			"InvalidMaxHoursZero",
			&MsgUpdateParamsRequest{From: base.TestBech32AccAddr20Bytes, Params: Params{MaxHours: 0}},
			true,
		},
		{
			"InvalidMinHoursNegative",
			&MsgUpdateParamsRequest{From: base.TestBech32AccAddr20Bytes, Params: Params{MaxHours: 10, MinHours: -1}},
			true,
		},
		{
			"InvalidMinHoursZero",
			&MsgUpdateParamsRequest{From: base.TestBech32AccAddr20Bytes, Params: Params{MinHours: 0}},
			true,
		},
		{
			"ValidMinAndMaxHours",
			&MsgUpdateParamsRequest{From: base.TestBech32AccAddr20Bytes, Params: Params{MaxHours: 10, MinHours: 1, StakingShare: sdkmath.LegacyNewDecWithPrec(5, 1)}},
			false,
		},
		{
			"InvalidStakingShareGreaterThanOne",
			&MsgUpdateParamsRequest{From: base.TestBech32AccAddr20Bytes, Params: Params{MaxHours: 10, MinHours: 1, StakingShare: sdkmath.LegacyNewDec(2)}},
			true,
		},
		{
			"InvalidStakingShareNegative",
			&MsgUpdateParamsRequest{From: base.TestBech32AccAddr20Bytes, Params: Params{MaxHours: 10, MinHours: 1, StakingShare: sdkmath.LegacyNewDec(-1)}},
			true,
		},
		{
			"StakingShareIsNil",
			&MsgUpdateParamsRequest{From: base.TestBech32AccAddr20Bytes, Params: Params{StakingShare: sdkmath.LegacyDec{}}},
			true,
		},
		{
			"ValidStakingShare",
			&MsgUpdateParamsRequest{From: base.TestBech32AccAddr20Bytes, Params: Params{MaxHours: 10, MinHours: 1, StakingShare: sdkmath.LegacyNewDecWithPrec(5, 1)}},
			false,
		},
		{
			"MaxHoursLessThanMinHours",
			&MsgUpdateParamsRequest{From: base.TestBech32AccAddr20Bytes, Params: Params{MaxHours: 5, MinHours: 10}},
			true,
		},
		{
			"ValidMaxHoursAndStakingShare",
			&MsgUpdateParamsRequest{From: base.TestBech32AccAddr20Bytes, Params: Params{MaxHours: 100, MinHours: 10, StakingShare: sdkmath.LegacyNewDecWithPrec(3, 1)}},
			false,
		},
		{
			"InvalidMaxHoursExceedsLimit",
			&MsgUpdateParamsRequest{From: base.TestBech32AccAddr20Bytes, Params: Params{MaxHours: 1000000000}},
			true,
		},
		{
			"InvalidMinHoursGreaterThanMax",
			&MsgUpdateParamsRequest{From: base.TestBech32AccAddr20Bytes, Params: Params{MaxHours: 20, MinHours: 30}},
			true,
		},
		{
			"ValidStakingShareAtOne",
			&MsgUpdateParamsRequest{From: base.TestBech32AccAddr20Bytes, Params: Params{MaxHours: 15, MinHours: 5, StakingShare: sdkmath.LegacyNewDec(1)}},
			false,
		},
		{
			"ValidStakingShareAtZero",
			&MsgUpdateParamsRequest{From: base.TestBech32AccAddr20Bytes, Params: Params{MaxHours: 15, MinHours: 5, StakingShare: sdkmath.LegacyNewDec(0)}},
			false,
		},
		{
			"InvalidStakingShareExceedsOneWithPrecision",
			&MsgUpdateParamsRequest{From: base.TestBech32AccAddr20Bytes, Params: Params{StakingShare: sdkmath.LegacyNewDecWithPrec(15, 1)}},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
