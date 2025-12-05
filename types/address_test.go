package types

import (
	"reflect"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestNodeAddressFromBech32(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name    string
		args    args
		want    NodeAddress
		wantErr bool
	}{
		{
			"empty",
			args{
				s: "",
			},
			NodeAddress{},
			true,
		},
		{
			"invalid",
			args{
				s: "invalid",
			},
			nil,
			true,
		},
		{
			"invalid prefix",
			args{
				s: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			nil,
			true,
		},
		{
			"20 bytes",
			args{
				s: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			NodeAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NodeAddressFromBech32(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("NodeAddressFromBech32() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NodeAddressFromBech32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNodeAddress_Bytes(t *testing.T) {
	tests := []struct {
		name string
		n    NodeAddress
		want []byte
	}{
		{
			"nil",
			nil,
			nil,
		},
		{
			"empty",
			NodeAddress{},
			[]byte{},
		},
		{
			"10 bytes",
			NodeAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10},
			[]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10},
		},
		{
			"20 bytes",
			NodeAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
			[]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
		},
		{
			"30 bytes",
			NodeAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x30},
			[]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x30},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Bytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNodeAddress_Empty(t *testing.T) {
	tests := []struct {
		name string
		n    NodeAddress
		want bool
	}{
		{
			"nil",
			nil,
			true},
		{
			"empty",
			NodeAddress{},
			true,
		},
		{
			"10 bytes",
			NodeAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10},
			false,
		},
		{
			"20 bytes",
			NodeAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
			false,
		},
		{
			"30 bytes",
			NodeAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x30},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Empty(); got != tt.want {
				t.Errorf("Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNodeAddress_Equals(t *testing.T) {
	type args struct {
		address sdk.Address
	}

	tests := []struct {
		name string
		n    NodeAddress
		args args
		want bool
	}{
		{
			"nil",
			nil,
			args{
				address: nil,
			},
			true,
		},
		{
			"equal type with 0 bytes",
			NodeAddress{},
			args{
				address: NodeAddress{},
			},
			true,
		},
		{
			"unequal type with 0 bytes",
			NodeAddress{},
			args{
				address: ProvAddress{},
			},
			true,
		},
		{
			"unequal type with unequal10 bytes",
			NodeAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10},
			args{
				address: ProvAddress{0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
			},
			false,
		},
		{
			"unequal type with equal 10 bytes",
			NodeAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10},
			args{
				address: ProvAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10},
			},
			true,
		},
		{
			"equal type with unequal 10 bytes",
			NodeAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10},
			args{
				address: NodeAddress{0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
			},
			false,
		},
		{
			"equal type with equal 10 bytes",
			NodeAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10},
			args{
				address: NodeAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Equals(tt.args.address); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNodeAddress_Marshal(t *testing.T) {
	tests := []struct {
		name    string
		n       NodeAddress
		want    []byte
		wantErr bool
	}{
		{
			"nil",
			nil,
			nil,
			false,
		},
		{
			"empty",
			NodeAddress{},
			[]byte{},
			false,
		},
		{
			"10 bytes",
			NodeAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10},
			[]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10},
			false,
		},
		{
			"20 bytes",
			NodeAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
			[]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
			false,
		},
		{
			"30 bytes",
			NodeAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x30},
			[]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x30},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.n.Marshal()
			if (err != nil) != tt.wantErr {
				t.Errorf("Marshal() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Marshal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNodeAddress_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		n       NodeAddress
		want    []byte
		wantErr bool
	}{
		{
			"nil",
			nil,
			[]byte(`""`),
			false,
		},
		{
			"empty",
			NodeAddress{},
			[]byte(`""`),
			false,
		},
		{
			"10 bytes",
			NodeAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10},
			[]byte(`"sentnode1qypqxpq9qcrsszgse4wwrm"`),
			false,
		},
		{
			"20 bytes",
			NodeAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
			[]byte(`"sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey"`),
			false,
		},
		{
			"30 bytes",
			NodeAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x30},
			[]byte(`"sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsxqglcv"`),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.n.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNodeAddress_MarshalYAML(t *testing.T) {
	tests := []struct {
		name    string
		n       NodeAddress
		want    any
		wantErr bool
	}{
		{
			"nil",
			nil,
			"",
			false,
		},
		{
			"empty",
			NodeAddress{},
			"",
			false,
		},
		{
			"10 bytes",
			NodeAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10},
			"sentnode1qypqxpq9qcrsszgse4wwrm",
			false,
		},
		{
			"20 bytes",
			NodeAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
			"sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			false,
		},
		{
			"30 bytes",
			NodeAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x30},
			"sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsxqglcv",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.n.MarshalYAML()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalYAML() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalYAML() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNodeAddress_String(t *testing.T) {
	tests := []struct {
		name string
		n    NodeAddress
		want string
	}{
		{
			"nil",
			nil,
			"",
		},
		{
			"empty",
			NodeAddress{},
			"",
		},
		{
			"10 bytes",
			NodeAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10},
			"sentnode1qypqxpq9qcrsszgse4wwrm",
		},
		{
			"20 bytes",
			NodeAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
			"sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
		},
		{
			"30 bytes",
			NodeAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x30},
			"sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsxqglcv",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNodeAddress_Unmarshal(t *testing.T) {
	type args struct {
		data []byte
	}

	tests := []struct {
		name    string
		n       NodeAddress
		args    args
		wantErr bool
	}{
		{
			"nil",
			nil,
			args{
				data: nil,
			},
			false,
		},
		{
			"empty",
			NodeAddress{},
			args{
				data: []byte{},
			},
			false,
		},
		{
			"10 bytes",
			NodeAddress{},
			args{
				data: []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10},
			},
			false,
		},
		{
			"20 bytes",
			NodeAddress{},
			args{
				data: []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
			},
			false,
		},
		{
			"30 bytes",
			NodeAddress{},
			args{
				data: []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x30},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.Unmarshal(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNodeAddress_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}

	tests := []struct {
		name    string
		n       NodeAddress
		args    args
		wantErr bool
	}{
		{
			"nil",
			nil,
			args{
				data: []byte(`""`),
			},
			true,
		},
		{
			"empty",
			NodeAddress{},
			args{
				data: []byte(`""`),
			},
			true,
		},
		{
			"10 bytes",
			NodeAddress{},
			args{
				data: []byte(`"sentnode1qypqxpq9qcrsszgse4wwrm"`),
			},
			false,
		},
		{
			"20 bytes",
			NodeAddress{},
			args{
				data: []byte(`"sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey"`),
			},
			false,
		},
		{
			"30 bytes",
			NodeAddress{},
			args{
				data: []byte(`"sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsxqglcv"`),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNodeAddress_UnmarshalYAML(t *testing.T) {
	type args struct {
		data []byte
	}

	tests := []struct {
		name    string
		n       NodeAddress
		args    args
		wantErr bool
	}{
		{
			"nil",
			nil,
			args{
				data: nil,
			},
			true,
		},
		{
			"empty",
			NodeAddress{},
			args{
				data: []byte(""),
			},
			true,
		},
		{
			"10 bytes",
			NodeAddress{},
			args{
				data: []byte("sentnode1qypqxpq9qcrsszgse4wwrm"),
			},
			false,
		},
		{
			"20 bytes",
			NodeAddress{},
			args{
				data: []byte("sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey"),
			},
			false,
		},
		{
			"30 bytes",
			NodeAddress{},
			args{
				data: []byte("sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsxqglcv"),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.UnmarshalYAML(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalYAML() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProvAddressFromBech32(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name    string
		args    args
		want    ProvAddress
		wantErr bool
	}{
		{
			"empty",
			args{
				s: "",
			},
			ProvAddress{},
			true,
		},
		{
			"invalid",
			args{
				s: "invalid",
			},
			nil,
			true,
		},
		{
			"invalid prefix",
			args{s: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj"},
			nil,
			true,
		},
		{
			"20 bytes",
			args{
				s: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
			},
			ProvAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ProvAddressFromBech32(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProvAddressFromBech32() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProvAddressFromBech32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProvAddress_Bytes(t *testing.T) {
	tests := []struct {
		name string
		p    ProvAddress
		want []byte
	}{
		{
			"nil",
			nil,
			nil,
		},
		{
			"empty",
			ProvAddress{},
			[]byte{},
		},
		{
			"10 bytes",
			ProvAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10},
			[]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10},
		},
		{
			"20 bytes",
			ProvAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
			[]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
		},
		{
			"30 bytes",
			ProvAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x30},
			[]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x30},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Bytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProvAddress_Empty(t *testing.T) {
	tests := []struct {
		name string
		p    ProvAddress
		want bool
	}{
		{
			"nil",
			nil,
			true,
		},
		{
			"empty",
			ProvAddress{},
			true,
		},
		{
			"10 bytes",
			ProvAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10},
			false,
		},
		{
			"20 bytes",
			ProvAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
			false,
		},
		{
			"30 bytes",
			ProvAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x30},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Empty(); got != tt.want {
				t.Errorf("Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProvAddress_Equals(t *testing.T) {
	type args struct {
		address sdk.Address
	}

	tests := []struct {
		name string
		p    ProvAddress
		args args
		want bool
	}{
		{
			"nil",
			nil,
			args{
				address: nil,
			},
			true,
		},
		{
			"equal type with 0 bytes",
			ProvAddress{},
			args{
				address: ProvAddress{},
			},
			true,
		},
		{
			"unequal type with 0 bytes",
			ProvAddress{},
			args{
				address: NodeAddress{},
			},
			true,
		},
		{
			"unequal type with unequal10 bytes",
			ProvAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10},
			args{
				address: NodeAddress{0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
			},
			false,
		},
		{
			"unequal type with equal 10 bytes",
			ProvAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10},
			args{
				address: NodeAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10},
			},
			true,
		},
		{
			"equal type with unequal 10 bytes",
			ProvAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10},
			args{
				address: ProvAddress{0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
			},
			false,
		},
		{
			"equal type with equal 10 bytes",
			ProvAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10},
			args{
				address: ProvAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Equals(tt.args.address); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProvAddress_Marshal(t *testing.T) {
	tests := []struct {
		name    string
		p       ProvAddress
		want    []byte
		wantErr bool
	}{
		{
			"nil",
			nil,
			nil,
			false,
		},
		{
			"empty",
			ProvAddress{},
			[]byte{},
			false,
		},
		{
			"10 bytes",
			ProvAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10},
			[]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10},
			false,
		},
		{
			"20 bytes",
			ProvAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
			[]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
			false,
		},
		{
			"30 bytes",
			ProvAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x30},
			[]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x30},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.Marshal()
			if (err != nil) != tt.wantErr {
				t.Errorf("Marshal() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Marshal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProvAddress_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		p       ProvAddress
		want    []byte
		wantErr bool
	}{
		{
			"nil",
			nil,
			[]byte(`""`),
			false,
		},
		{
			"empty",
			ProvAddress{},
			[]byte(`""`),
			false,
		},
		{
			"10 bytes",
			ProvAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10},
			[]byte(`"sentprov1qypqxpq9qcrsszgsutj8xr"`),
			false,
		},
		{
			"20 bytes",
			ProvAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
			[]byte(`"sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82"`),
			false,
		},
		{
			"30 bytes",
			ProvAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x30},
			[]byte(`"sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsh33zgx"`),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProvAddress_MarshalYAML(t *testing.T) {
	tests := []struct {
		name    string
		p       ProvAddress
		want    any
		wantErr bool
	}{
		{
			"nil",
			nil,
			"",
			false,
		},
		{
			"empty",
			ProvAddress{},
			"",
			false,
		},
		{
			"10 bytes",
			ProvAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10},
			"sentprov1qypqxpq9qcrsszgsutj8xr",
			false,
		},
		{
			"20 bytes",
			ProvAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
			"sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
			false,
		},
		{
			"30 bytes",
			ProvAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x30},
			"sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsh33zgx",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.MarshalYAML()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalYAML() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalYAML() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProvAddress_String(t *testing.T) {
	tests := []struct {
		name string
		p    ProvAddress
		want string
	}{
		{
			"nil",
			nil,
			"",
		},
		{
			"empty",
			ProvAddress{},
			"",
		},
		{
			"10 bytes",
			ProvAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10},
			"sentprov1qypqxpq9qcrsszgsutj8xr",
		},
		{
			"20 bytes",
			ProvAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
			"sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
		},
		{
			"30 bytes",
			ProvAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x30},
			"sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsh33zgx",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProvAddress_Unmarshal(t *testing.T) {
	type args struct {
		data []byte
	}

	tests := []struct {
		name    string
		p       ProvAddress
		args    args
		wantErr bool
	}{
		{
			"nil",
			nil,
			args{
				data: nil,
			},
			false,
		},
		{
			"empty",
			ProvAddress{},
			args{
				data: []byte{},
			},
			false,
		},
		{
			"10 bytes",
			ProvAddress{},
			args{
				data: []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10},
			},
			false,
		},
		{
			"20 bytes",
			ProvAddress{},
			args{
				data: []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
			},
			false,
		},
		{
			"30 bytes",
			ProvAddress{},
			args{
				data: []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x30},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.Unmarshal(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProvAddress_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}

	tests := []struct {
		name    string
		p       ProvAddress
		args    args
		wantErr bool
	}{
		{
			"nil",
			nil,
			args{
				data: []byte(`""`),
			},
			true,
		},
		{
			"empty",
			ProvAddress{},
			args{
				data: []byte(`""`),
			},
			true,
		},
		{
			"10 bytes",
			ProvAddress{},
			args{
				data: []byte(`"sentprov1qypqxpq9qcrsszgsutj8xr"`),
			},
			false,
		},
		{
			"20 bytes",
			ProvAddress{},
			args{
				data: []byte(`"sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82"`),
			},
			false,
		},
		{
			"30 bytes",
			ProvAddress{},
			args{
				data: []byte(`"sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsh33zgx"`),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProvAddress_UnmarshalYAML(t *testing.T) {
	type args struct {
		data []byte
	}

	tests := []struct {
		name    string
		p       ProvAddress
		args    args
		wantErr bool
	}{
		{
			"nil",
			nil,
			args{
				data: nil,
			},
			true,
		},
		{
			"empty",
			ProvAddress{},
			args{
				data: []byte(""),
			},
			true,
		},
		{
			"10 bytes",
			ProvAddress{},
			args{
				data: []byte("sentprov1qypqxpq9qcrsszgsutj8xr"),
			},
			false,
		},
		{
			"20 bytes",
			ProvAddress{},
			args{
				data: []byte("sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82"),
			},
			false,
		},
		{
			"30 bytes",
			ProvAddress{},
			args{
				data: []byte("sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsh33zgx"),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.UnmarshalYAML(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalYAML() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
