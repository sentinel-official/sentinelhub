package v1

import (
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/sentinel-official/sentinelhub/v12/x/swap/types"
)

var (
	DefaultSwapEnabled = false
	DefaultSwapDenom   = sdk.DefaultBondDenom
	DefaultApproveBy   = authtypes.NewModuleAddress(types.ModuleName).String()
)

var (
	KeySwapEnabled = []byte("SwapEnabled")
	KeySwapDenom   = []byte("SwapDenom")
	KeyApproveBy   = []byte("ApproveBy")
)

var (
	_ paramstypes.ParamSet = (*Params)(nil)
)

func (m *Params) Validate() error {
	if m.SwapDenom == "" {
		return fmt.Errorf("swap_denom cannot be empty")
	}
	if err := sdk.ValidateDenom(m.SwapDenom); err != nil {
		return sdkerrors.Wrapf(err, "invalid swap_denom %s", m.SwapDenom)
	}
	if m.ApproveBy == "" {
		return fmt.Errorf("approve_by cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.ApproveBy); err != nil {
		return sdkerrors.Wrapf(err, "invalid approve_by %s", m.ApproveBy)
	}

	return nil
}

func (m *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		{
			Key:   KeySwapEnabled,
			Value: &m.SwapEnabled,
			ValidatorFn: func(v interface{}) error {
				_, ok := v.(bool)
				if !ok {
					return fmt.Errorf("invalid parameter type %T", v)
				}

				return nil
			},
		},
		{
			Key:   KeySwapDenom,
			Value: &m.SwapDenom,
			ValidatorFn: func(v interface{}) error {
				value, ok := v.(string)
				if !ok {
					return fmt.Errorf("invalid parameter type %T", v)
				}

				if value == "" {
					return fmt.Errorf("value cannot be empty")
				}
				if err := sdk.ValidateDenom(value); err != nil {
					return sdkerrors.Wrapf(err, "invalid value %s", value)
				}

				return nil
			},
		},
		{
			Key:   KeyApproveBy,
			Value: &m.ApproveBy,
			ValidatorFn: func(v interface{}) error {
				value, ok := v.(string)
				if !ok {
					return fmt.Errorf("invalid parameter type %T", v)
				}

				if value == "" {
					return fmt.Errorf("value cannot be empty")
				}
				if _, err := sdk.AccAddressFromBech32(value); err != nil {
					return sdkerrors.Wrapf(err, "invalid value %s", value)
				}

				return nil
			},
		},
	}
}

func NewParams(swapEnabled bool, swapDenom, approveBy string) Params {
	return Params{
		SwapEnabled: swapEnabled,
		SwapDenom:   swapDenom,
		ApproveBy:   approveBy,
	}
}

func DefaultParams() Params {
	return Params{
		SwapEnabled: DefaultSwapEnabled,
		SwapDenom:   DefaultSwapDenom,
		ApproveBy:   DefaultApproveBy,
	}
}

func ParamsKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}
