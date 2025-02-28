package v2

import (
	"testing"

	sdkmath "cosmossdk.io/math"

	base "github.com/sentinel-official/hub/v12/types"
)

func TestAllocation_Validate(t *testing.T) {
	type fields struct {
		Address       string
		GrantedBytes  sdkmath.Int
		UtilisedBytes sdkmath.Int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"empty account address",
			fields{
				Address: base.TestAddrEmpty,
			},
			true,
		},
		{
			"invalid account address",
			fields{
				Address: base.TestAddrInvalid,
			},
			true,
		},
		{
			"invalid prefix account address",
			fields{
				Address: base.TestBech32NodeAddr20Bytes,
			},
			true,
		},
		{
			"10 bytes account address",
			fields{
				Address:       base.TestBech32AccAddr10Bytes,
				GrantedBytes:  sdkmath.NewInt(0),
				UtilisedBytes: sdkmath.NewInt(0),
			},
			false,
		},
		{
			"20 bytes account address",
			fields{
				Address:       base.TestBech32AccAddr20Bytes,
				GrantedBytes:  sdkmath.NewInt(0),
				UtilisedBytes: sdkmath.NewInt(0),
			},
			false,
		},
		{
			"30 bytes account address",
			fields{
				Address:       base.TestBech32AccAddr30Bytes,
				GrantedBytes:  sdkmath.NewInt(0),
				UtilisedBytes: sdkmath.NewInt(0),
			},
			false,
		},
		{
			"nil granted",
			fields{
				Address:      base.TestBech32AccAddr20Bytes,
				GrantedBytes: sdkmath.Int{},
			},
			true,
		},
		{
			"negative granted",
			fields{
				Address:      base.TestBech32AccAddr20Bytes,
				GrantedBytes: sdkmath.NewInt(-1000),
			},
			true,
		},
		{
			"zero granted",
			fields{
				Address:       base.TestBech32AccAddr20Bytes,
				GrantedBytes:  sdkmath.NewInt(0),
				UtilisedBytes: sdkmath.NewInt(0),
			},
			false,
		},
		{
			"positive granted",
			fields{
				Address:       base.TestBech32AccAddr20Bytes,
				GrantedBytes:  sdkmath.NewInt(1000),
				UtilisedBytes: sdkmath.NewInt(0),
			},
			false,
		},
		{
			"nil utilised",
			fields{
				Address:       base.TestBech32AccAddr20Bytes,
				GrantedBytes:  sdkmath.NewInt(1000),
				UtilisedBytes: sdkmath.Int{},
			},
			true,
		},
		{
			"negative utilised",
			fields{
				Address:       base.TestBech32AccAddr20Bytes,
				GrantedBytes:  sdkmath.NewInt(1000),
				UtilisedBytes: sdkmath.NewInt(-1000),
			},
			true,
		},
		{
			"zero utilised",
			fields{
				Address:       base.TestBech32AccAddr20Bytes,
				GrantedBytes:  sdkmath.NewInt(1000),
				UtilisedBytes: sdkmath.NewInt(0),
			},
			false,
		},
		{
			"positive utilised",
			fields{
				Address:       base.TestBech32AccAddr20Bytes,
				GrantedBytes:  sdkmath.NewInt(1000),
				UtilisedBytes: sdkmath.NewInt(1000),
			},
			false,
		},
		{
			"granted less than utilised",
			fields{
				Address:       base.TestBech32AccAddr20Bytes,
				GrantedBytes:  sdkmath.NewInt(1000),
				UtilisedBytes: sdkmath.NewInt(2000),
			},
			true,
		},
		{
			"granted equals to utilised",
			fields{
				Address:       base.TestBech32AccAddr20Bytes,
				GrantedBytes:  sdkmath.NewInt(2000),
				UtilisedBytes: sdkmath.NewInt(2000),
			},
			false,
		},
		{
			"granted greater than utilised",
			fields{
				Address:       base.TestBech32AccAddr20Bytes,
				GrantedBytes:  sdkmath.NewInt(2000),
				UtilisedBytes: sdkmath.NewInt(1000),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Allocation{
				Address:       tt.fields.Address,
				GrantedBytes:  tt.fields.GrantedBytes,
				UtilisedBytes: tt.fields.UtilisedBytes,
			}
			if err := m.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
