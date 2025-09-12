package v1

import (
	"reflect"
	"testing"

	sdkmath "cosmossdk.io/math"
)

func TestBandwidth_Add(t *testing.T) {
	type fields struct {
		Upload   sdkmath.Int
		Download sdkmath.Int
	}

	type args struct {
		v Bandwidth
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   Bandwidth
	}{
		{
			"negative upload and negative download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdkmath.NewInt(-1000), Download: sdkmath.NewInt(-1000)},
			},
			Bandwidth{Upload: sdkmath.NewInt(-1000), Download: sdkmath.NewInt(-1000)},
		},
		{
			"negative upload and zero download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdkmath.NewInt(-1000), Download: sdkmath.NewInt(0)},
			},
			Bandwidth{Upload: sdkmath.NewInt(-1000), Download: sdkmath.NewInt(0)},
		},
		{
			"negative upload and positive download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdkmath.NewInt(-1000), Download: sdkmath.NewInt(1000)},
			},
			Bandwidth{Upload: sdkmath.NewInt(-1000), Download: sdkmath.NewInt(1000)},
		},
		{
			"zero upload and negative download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdkmath.NewInt(0), Download: sdkmath.NewInt(-1000)},
			},
			Bandwidth{Upload: sdkmath.NewInt(0), Download: sdkmath.NewInt(-1000)},
		},
		{
			"zero upload and zero download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdkmath.NewInt(0), Download: sdkmath.NewInt(0)},
			},
			Bandwidth{Upload: sdkmath.NewInt(0), Download: sdkmath.NewInt(0)},
		},
		{
			"zero upload and positive download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdkmath.NewInt(0), Download: sdkmath.NewInt(1000)},
			},
			Bandwidth{Upload: sdkmath.NewInt(0), Download: sdkmath.NewInt(1000)},
		},
		{
			"positive upload and negative download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdkmath.NewInt(1000), Download: sdkmath.NewInt(-1000)},
			},
			Bandwidth{Upload: sdkmath.NewInt(1000), Download: sdkmath.NewInt(-1000)},
		},
		{
			"positive upload and zero download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdkmath.NewInt(1000), Download: sdkmath.NewInt(0)},
			},
			Bandwidth{Upload: sdkmath.NewInt(1000), Download: sdkmath.NewInt(0)},
		},
		{
			"positive upload and positive download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdkmath.NewInt(1000), Download: sdkmath.NewInt(1000)},
			},
			Bandwidth{Upload: sdkmath.NewInt(1000), Download: sdkmath.NewInt(1000)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Bandwidth{
				Upload:   tt.fields.Upload,
				Download: tt.fields.Download,
			}
			if got := b.Add(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBandwidth_CeilTo(t *testing.T) {
	type fields struct {
		Upload   sdkmath.Int
		Download sdkmath.Int
	}

	type args struct {
		precision sdkmath.Int
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   Bandwidth
	}{
		{
			"negative pre",
			fields{Upload: sdkmath.NewInt(10), Download: sdkmath.NewInt(10)},
			args{precision: sdkmath.NewInt(-17)},
			Bandwidth{sdkmath.NewInt(10), sdkmath.NewInt(10)},
		},
		{
			"0 upload 0 download",
			fields{Upload: sdkmath.NewInt(0), Download: sdkmath.NewInt(0)},
			args{precision: sdkmath.NewInt(17)},
			Bandwidth{sdkmath.NewInt(0), sdkmath.NewInt(0)},
		},
		{
			"0 upload 10 download",
			fields{Upload: sdkmath.NewInt(0), Download: sdkmath.NewInt(10)},
			args{precision: sdkmath.NewInt(17)},
			Bandwidth{sdkmath.NewInt(0), sdkmath.NewInt(17)},
		},
		{
			"0 upload 17 download",
			fields{Upload: sdkmath.NewInt(0), Download: sdkmath.NewInt(17)},
			args{precision: sdkmath.NewInt(17)},
			Bandwidth{sdkmath.NewInt(0), sdkmath.NewInt(17)},
		},
		{
			"0 upload 20 download",
			fields{Upload: sdkmath.NewInt(0), Download: sdkmath.NewInt(20)},
			args{precision: sdkmath.NewInt(17)},
			Bandwidth{sdkmath.NewInt(0), sdkmath.NewInt(34)},
		},
		{
			"0 upload 34 download",
			fields{Upload: sdkmath.NewInt(0), Download: sdkmath.NewInt(34)},
			args{precision: sdkmath.NewInt(17)},
			Bandwidth{sdkmath.NewInt(0), sdkmath.NewInt(34)},
		},
		{
			"10 upload 0 download",
			fields{Upload: sdkmath.NewInt(10), Download: sdkmath.NewInt(0)},
			args{precision: sdkmath.NewInt(17)},
			Bandwidth{sdkmath.NewInt(17), sdkmath.NewInt(0)},
		},
		{
			"10 upload 10 download",
			fields{Upload: sdkmath.NewInt(10), Download: sdkmath.NewInt(10)},
			args{precision: sdkmath.NewInt(17)},
			Bandwidth{sdkmath.NewInt(17), sdkmath.NewInt(17)},
		},
		{
			"10 upload 17 download",
			fields{Upload: sdkmath.NewInt(10), Download: sdkmath.NewInt(17)},
			args{precision: sdkmath.NewInt(17)},
			Bandwidth{sdkmath.NewInt(17), sdkmath.NewInt(17)},
		},
		{
			"10 upload 20 download",
			fields{Upload: sdkmath.NewInt(10), Download: sdkmath.NewInt(20)},
			args{precision: sdkmath.NewInt(17)},
			Bandwidth{sdkmath.NewInt(17), sdkmath.NewInt(34)},
		},
		{
			"10 upload 34 download",
			fields{Upload: sdkmath.NewInt(10), Download: sdkmath.NewInt(34)},
			args{precision: sdkmath.NewInt(17)},
			Bandwidth{sdkmath.NewInt(17), sdkmath.NewInt(34)},
		},
		{
			"17 upload 0 download",
			fields{Upload: sdkmath.NewInt(17), Download: sdkmath.NewInt(0)},
			args{precision: sdkmath.NewInt(17)},
			Bandwidth{sdkmath.NewInt(17), sdkmath.NewInt(0)},
		},
		{
			"17 upload 10 download",
			fields{Upload: sdkmath.NewInt(17), Download: sdkmath.NewInt(10)},
			args{precision: sdkmath.NewInt(17)},
			Bandwidth{sdkmath.NewInt(17), sdkmath.NewInt(17)},
		},
		{
			"17 upload 17 download",
			fields{Upload: sdkmath.NewInt(17), Download: sdkmath.NewInt(17)},
			args{precision: sdkmath.NewInt(17)},
			Bandwidth{sdkmath.NewInt(17), sdkmath.NewInt(17)},
		},
		{
			"17 upload 20 download",
			fields{Upload: sdkmath.NewInt(17), Download: sdkmath.NewInt(20)},
			args{precision: sdkmath.NewInt(17)},
			Bandwidth{sdkmath.NewInt(17), sdkmath.NewInt(34)},
		},
		{
			"17 upload 34 download",
			fields{Upload: sdkmath.NewInt(17), Download: sdkmath.NewInt(34)},
			args{precision: sdkmath.NewInt(17)},
			Bandwidth{sdkmath.NewInt(17), sdkmath.NewInt(34)},
		},
		{
			"20 upload 0 download",
			fields{Upload: sdkmath.NewInt(20), Download: sdkmath.NewInt(0)},
			args{precision: sdkmath.NewInt(17)},
			Bandwidth{sdkmath.NewInt(34), sdkmath.NewInt(0)},
		},
		{
			"20 upload 10 download",
			fields{Upload: sdkmath.NewInt(20), Download: sdkmath.NewInt(10)},
			args{precision: sdkmath.NewInt(17)},
			Bandwidth{sdkmath.NewInt(34), sdkmath.NewInt(17)},
		},
		{
			"20 upload 17 download",
			fields{Upload: sdkmath.NewInt(20), Download: sdkmath.NewInt(17)},
			args{precision: sdkmath.NewInt(17)},
			Bandwidth{sdkmath.NewInt(34), sdkmath.NewInt(17)},
		},
		{
			"20 upload 20 download",
			fields{Upload: sdkmath.NewInt(20), Download: sdkmath.NewInt(20)},
			args{precision: sdkmath.NewInt(17)},
			Bandwidth{sdkmath.NewInt(34), sdkmath.NewInt(34)},
		},
		{
			"20 upload 34 download",
			fields{Upload: sdkmath.NewInt(20), Download: sdkmath.NewInt(34)},
			args{precision: sdkmath.NewInt(17)},
			Bandwidth{sdkmath.NewInt(34), sdkmath.NewInt(34)},
		},
		{
			"34 upload 0 download",
			fields{Upload: sdkmath.NewInt(34), Download: sdkmath.NewInt(0)},
			args{precision: sdkmath.NewInt(17)},
			Bandwidth{sdkmath.NewInt(34), sdkmath.NewInt(0)},
		},
		{
			"34 upload 10 download",
			fields{Upload: sdkmath.NewInt(34), Download: sdkmath.NewInt(10)},
			args{precision: sdkmath.NewInt(17)},
			Bandwidth{sdkmath.NewInt(34), sdkmath.NewInt(17)},
		},
		{
			"34 upload 17 download",
			fields{Upload: sdkmath.NewInt(34), Download: sdkmath.NewInt(17)},
			args{precision: sdkmath.NewInt(17)},
			Bandwidth{sdkmath.NewInt(34), sdkmath.NewInt(17)},
		},
		{
			"34 upload 20 download",
			fields{Upload: sdkmath.NewInt(34), Download: sdkmath.NewInt(20)},
			args{precision: sdkmath.NewInt(17)},
			Bandwidth{sdkmath.NewInt(34), sdkmath.NewInt(34)},
		},
		{
			"34 upload 34 download",
			fields{Upload: sdkmath.NewInt(34), Download: sdkmath.NewInt(34)},
			args{precision: sdkmath.NewInt(17)},
			Bandwidth{sdkmath.NewInt(34), sdkmath.NewInt(34)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Bandwidth{
				Upload:   tt.fields.Upload,
				Download: tt.fields.Download,
			}
			if got := b.CeilTo(tt.args.precision); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CeilTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBandwidth_IsAllLTE(t *testing.T) {
	type fields struct {
		Upload   sdkmath.Int
		Download sdkmath.Int
	}

	type args struct {
		v Bandwidth
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"negative upload and negative download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdkmath.NewInt(-1000), Download: sdkmath.NewInt(-1000)},
			},
			false,
		},
		{
			"negative upload and zero download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdkmath.NewInt(-1000), Download: sdkmath.NewInt(0)},
			},
			false,
		},
		{
			"negative upload and positive download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdkmath.NewInt(-1000), Download: sdkmath.NewInt(1000)},
			},
			false,
		},
		{
			"zero upload and negative download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdkmath.NewInt(0), Download: sdkmath.NewInt(-1000)},
			},
			false,
		},
		{
			"zero upload and zero download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdkmath.NewInt(0), Download: sdkmath.NewInt(0)},
			},
			true,
		},
		{
			"zero upload and positive download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdkmath.NewInt(0), Download: sdkmath.NewInt(1000)},
			},
			true,
		},
		{
			"positive upload and negative download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdkmath.NewInt(1000), Download: sdkmath.NewInt(-1000)},
			},
			false,
		},
		{
			"positive upload and zero download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdkmath.NewInt(1000), Download: sdkmath.NewInt(0)},
			},
			true,
		},
		{
			"positive upload and positive download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdkmath.NewInt(1000), Download: sdkmath.NewInt(1000)},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Bandwidth{
				Upload:   tt.fields.Upload,
				Download: tt.fields.Download,
			}
			if got := b.IsAllLTE(tt.args.v); got != tt.want {
				t.Errorf("IsAllLTE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBandwidth_IsAllPositive(t *testing.T) {
	type fields struct {
		Upload   sdkmath.Int
		Download sdkmath.Int
	}

	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"negative upload and negative download",
			fields{
				Upload:   sdkmath.NewInt(-1000),
				Download: sdkmath.NewInt(-1000),
			},
			false,
		},
		{
			"negative upload and zero download",
			fields{
				Upload:   sdkmath.NewInt(-1000),
				Download: sdkmath.NewInt(0),
			},
			false,
		},
		{
			"negative upload and positive download",
			fields{
				Upload:   sdkmath.NewInt(-1000),
				Download: sdkmath.NewInt(1000),
			},
			false,
		},
		{
			"zero upload and negative download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(-1000),
			},
			false,
		},
		{
			"zero upload and zero download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			false,
		},
		{
			"zero upload and positive download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(1000),
			},
			false,
		},
		{
			"positive upload and negative download",
			fields{
				Upload:   sdkmath.NewInt(1000),
				Download: sdkmath.NewInt(-1000),
			},
			false,
		},
		{
			"positive upload and zero download",
			fields{
				Upload:   sdkmath.NewInt(1000),
				Download: sdkmath.NewInt(0),
			},
			false,
		},
		{
			"positive upload and positive download",
			fields{
				Upload:   sdkmath.NewInt(1000),
				Download: sdkmath.NewInt(1000),
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Bandwidth{
				Upload:   tt.fields.Upload,
				Download: tt.fields.Download,
			}
			if got := b.IsAllPositive(); got != tt.want {
				t.Errorf("IsAllPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBandwidth_IsAllZero(t *testing.T) {
	type fields struct {
		Upload   sdkmath.Int
		Download sdkmath.Int
	}

	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"negative upload and negative download",
			fields{
				Upload:   sdkmath.NewInt(-1000),
				Download: sdkmath.NewInt(-1000),
			},
			false,
		},
		{
			"negative upload and zero download",
			fields{
				Upload:   sdkmath.NewInt(-1000),
				Download: sdkmath.NewInt(0),
			},
			false,
		},
		{
			"negative upload and positive download",
			fields{
				Upload:   sdkmath.NewInt(-1000),
				Download: sdkmath.NewInt(1000),
			},
			false,
		},
		{
			"zero upload and negative download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(-1000),
			},
			false,
		},
		{
			"zero upload and zero download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			true,
		},
		{
			"zero upload and positive download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(1000),
			},
			false,
		},
		{
			"positive upload and negative download",
			fields{
				Upload:   sdkmath.NewInt(1000),
				Download: sdkmath.NewInt(-1000),
			},
			false,
		},
		{
			"positive upload and zero download",
			fields{
				Upload:   sdkmath.NewInt(1000),
				Download: sdkmath.NewInt(0),
			},
			false,
		},
		{
			"positive upload and positive download",
			fields{
				Upload:   sdkmath.NewInt(1000),
				Download: sdkmath.NewInt(1000),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Bandwidth{
				Upload:   tt.fields.Upload,
				Download: tt.fields.Download,
			}
			if got := b.IsAllZero(); got != tt.want {
				t.Errorf("IsAllZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBandwidth_IsAnyGT(t *testing.T) {
	type fields struct {
		Upload   sdkmath.Int
		Download sdkmath.Int
	}

	type args struct {
		v Bandwidth
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"negative upload and negative download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdkmath.NewInt(-1000), Download: sdkmath.NewInt(-1000)},
			},
			true,
		},
		{
			"negative upload and zero download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdkmath.NewInt(-1000), Download: sdkmath.NewInt(0)},
			},
			true,
		},
		{
			"negative upload and positive download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdkmath.NewInt(-1000), Download: sdkmath.NewInt(1000)},
			},
			true,
		},
		{
			"zero upload and negative download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdkmath.NewInt(0), Download: sdkmath.NewInt(-1000)},
			},
			true,
		},
		{
			"zero upload and zero download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdkmath.NewInt(0), Download: sdkmath.NewInt(0)},
			},
			false,
		},
		{
			"zero upload and positive download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdkmath.NewInt(0), Download: sdkmath.NewInt(1000)},
			},
			false,
		},
		{
			"positive upload and negative download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdkmath.NewInt(1000), Download: sdkmath.NewInt(-1000)},
			},
			true,
		},
		{
			"positive upload and zero download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdkmath.NewInt(1000), Download: sdkmath.NewInt(0)},
			},
			false,
		},
		{
			"positive upload and positive download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdkmath.NewInt(1000), Download: sdkmath.NewInt(1000)},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Bandwidth{
				Upload:   tt.fields.Upload,
				Download: tt.fields.Download,
			}
			if got := b.IsAnyGT(tt.args.v); got != tt.want {
				t.Errorf("IsAnyGT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBandwidth_IsAnyNegative(t *testing.T) {
	type fields struct {
		Upload   sdkmath.Int
		Download sdkmath.Int
	}

	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"negative upload and negative download",
			fields{
				Upload:   sdkmath.NewInt(-1000),
				Download: sdkmath.NewInt(-1000),
			},
			true,
		},
		{
			"negative upload and zero download",
			fields{
				Upload:   sdkmath.NewInt(-1000),
				Download: sdkmath.NewInt(0),
			},
			true,
		},
		{
			"negative upload and positive download",
			fields{
				Upload:   sdkmath.NewInt(-1000),
				Download: sdkmath.NewInt(1000),
			},
			true,
		},
		{
			"zero upload and negative download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(-1000),
			},
			true,
		},
		{
			"zero upload and zero download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			false,
		},
		{
			"zero upload and positive download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(1000),
			},
			false,
		},
		{
			"positive upload and negative download",
			fields{
				Upload:   sdkmath.NewInt(1000),
				Download: sdkmath.NewInt(-1000),
			},
			true,
		},
		{
			"positive upload and zero download",
			fields{
				Upload:   sdkmath.NewInt(1000),
				Download: sdkmath.NewInt(0),
			},
			false,
		},
		{
			"positive upload and positive download",
			fields{
				Upload:   sdkmath.NewInt(1000),
				Download: sdkmath.NewInt(1000),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Bandwidth{
				Upload:   tt.fields.Upload,
				Download: tt.fields.Download,
			}
			if got := b.IsAnyNegative(); got != tt.want {
				t.Errorf("IsAnyNegative() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBandwidth_IsAnyZero(t *testing.T) {
	type fields struct {
		Upload   sdkmath.Int
		Download sdkmath.Int
	}

	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"negative upload and negative download",
			fields{
				Upload:   sdkmath.NewInt(-1000),
				Download: sdkmath.NewInt(-1000),
			},
			false,
		},
		{
			"negative upload and zero download",
			fields{
				Upload:   sdkmath.NewInt(-1000),
				Download: sdkmath.NewInt(0),
			},
			true,
		},
		{
			"negative upload and positive download",
			fields{
				Upload:   sdkmath.NewInt(-1000),
				Download: sdkmath.NewInt(1000),
			},
			false,
		},
		{
			"zero upload and negative download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(-1000),
			},
			true,
		},
		{
			"zero upload and zero download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			true,
		},
		{
			"zero upload and positive download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(1000),
			},
			true,
		},
		{
			"positive upload and negative download",
			fields{
				Upload:   sdkmath.NewInt(1000),
				Download: sdkmath.NewInt(-1000),
			},
			false,
		},
		{
			"positive upload and zero download",
			fields{
				Upload:   sdkmath.NewInt(1000),
				Download: sdkmath.NewInt(0),
			},
			true,
		},
		{
			"positive upload and positive download",
			fields{
				Upload:   sdkmath.NewInt(1000),
				Download: sdkmath.NewInt(1000),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Bandwidth{
				Upload:   tt.fields.Upload,
				Download: tt.fields.Download,
			}
			if got := b.IsAnyZero(); got != tt.want {
				t.Errorf("IsAnyZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBandwidth_Sub(t *testing.T) {
	type fields struct {
		Upload   sdkmath.Int
		Download sdkmath.Int
	}

	type args struct {
		v Bandwidth
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   Bandwidth
	}{
		{
			"negative upload and negative download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdkmath.NewInt(-1000), Download: sdkmath.NewInt(-1000)},
			},
			Bandwidth{Upload: sdkmath.NewInt(1000), Download: sdkmath.NewInt(1000)},
		},
		{
			"negative upload and zero download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdkmath.NewInt(-1000), Download: sdkmath.NewInt(0)},
			},
			Bandwidth{Upload: sdkmath.NewInt(1000), Download: sdkmath.NewInt(0)},
		},
		{
			"negative upload and positive download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdkmath.NewInt(-1000), Download: sdkmath.NewInt(1000)},
			},
			Bandwidth{Upload: sdkmath.NewInt(1000), Download: sdkmath.NewInt(-1000)},
		},
		{
			"zero upload and negative download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdkmath.NewInt(0), Download: sdkmath.NewInt(-1000)},
			},
			Bandwidth{Upload: sdkmath.NewInt(0), Download: sdkmath.NewInt(1000)},
		},
		{
			"zero upload and zero download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdkmath.NewInt(0), Download: sdkmath.NewInt(0)},
			},
			Bandwidth{Upload: sdkmath.NewInt(0), Download: sdkmath.NewInt(0)},
		},
		{
			"zero upload and positive download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdkmath.NewInt(0), Download: sdkmath.NewInt(1000)},
			},
			Bandwidth{Upload: sdkmath.NewInt(0), Download: sdkmath.NewInt(-1000)},
		},
		{
			"positive upload and negative download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdkmath.NewInt(1000), Download: sdkmath.NewInt(-1000)},
			},
			Bandwidth{Upload: sdkmath.NewInt(-1000), Download: sdkmath.NewInt(1000)},
		},
		{
			"positive upload and zero download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdkmath.NewInt(1000), Download: sdkmath.NewInt(0)},
			},
			Bandwidth{Upload: sdkmath.NewInt(-1000), Download: sdkmath.NewInt(0)},
		},
		{
			"positive upload and positive download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdkmath.NewInt(1000), Download: sdkmath.NewInt(1000)},
			},
			Bandwidth{Upload: sdkmath.NewInt(-1000), Download: sdkmath.NewInt(-1000)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Bandwidth{
				Upload:   tt.fields.Upload,
				Download: tt.fields.Download,
			}
			if got := b.Sub(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBandwidth_Sum(t *testing.T) {
	type fields struct {
		Upload   sdkmath.Int
		Download sdkmath.Int
	}

	tests := []struct {
		name   string
		fields fields
		want   sdkmath.Int
	}{
		{
			"negative upload and negative download",
			fields{
				Upload:   sdkmath.NewInt(-1000),
				Download: sdkmath.NewInt(-1000),
			},
			sdkmath.NewInt(-2000),
		},
		{
			"negative upload and zero download",
			fields{
				Upload:   sdkmath.NewInt(-1000),
				Download: sdkmath.NewInt(0),
			},
			sdkmath.NewInt(-1000),
		},
		{
			"negative upload and positive download",
			fields{
				Upload:   sdkmath.NewInt(-1000),
				Download: sdkmath.NewInt(1000),
			},
			sdkmath.NewInt(1).Add(sdkmath.NewInt(-1)),
		},
		{
			"zero upload and negative download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(-1000),
			},
			sdkmath.NewInt(-1000),
		},
		{
			"zero upload and zero download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(0),
			},
			sdkmath.NewInt(0),
		},
		{
			"zero upload and positive download",
			fields{
				Upload:   sdkmath.NewInt(0),
				Download: sdkmath.NewInt(1000),
			},
			sdkmath.NewInt(1000),
		},
		{
			"positive upload and negative download",
			fields{
				Upload:   sdkmath.NewInt(1000),
				Download: sdkmath.NewInt(-1000),
			},
			sdkmath.NewInt(1).Add(sdkmath.NewInt(-1)),
		},
		{
			"positive upload and zero download",
			fields{
				Upload:   sdkmath.NewInt(1000),
				Download: sdkmath.NewInt(0),
			},
			sdkmath.NewInt(1000),
		},
		{
			"positive upload and positive download",
			fields{
				Upload:   sdkmath.NewInt(1000),
				Download: sdkmath.NewInt(1000),
			},
			sdkmath.NewInt(2000),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Bandwidth{
				Upload:   tt.fields.Upload,
				Download: tt.fields.Download,
			}
			if got := b.Sum(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBandwidth(t *testing.T) {
	type args struct {
		upload   sdkmath.Int
		download sdkmath.Int
	}

	tests := []struct {
		name string
		args args
		want Bandwidth
	}{
		{
			"0 upload 0 download",
			args{upload: sdkmath.NewInt(0), download: sdkmath.NewInt(0)},
			Bandwidth{sdkmath.NewInt(0), sdkmath.NewInt(0)},
		},
		{
			"0 upload 10 download",
			args{upload: sdkmath.NewInt(0), download: sdkmath.NewInt(10)},
			Bandwidth{sdkmath.NewInt(0), sdkmath.NewInt(10)},
		},
		{
			"10 upload 0 download",
			args{upload: sdkmath.NewInt(10), download: sdkmath.NewInt(0)},
			Bandwidth{sdkmath.NewInt(10), sdkmath.NewInt(0)},
		},
		{
			"10 upload 10 download",
			args{upload: sdkmath.NewInt(10), download: sdkmath.NewInt(10)},
			Bandwidth{sdkmath.NewInt(10), sdkmath.NewInt(10)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBandwidth(tt.args.upload, tt.args.download); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBandwidth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBandwidthFromInt64(t *testing.T) {
	type args struct {
		upload   int64
		download int64
	}

	tests := []struct {
		name string
		args args
		want Bandwidth
	}{
		{
			"0 upload 0 download",
			args{upload: 0, download: 0},
			Bandwidth{sdkmath.NewInt(0), sdkmath.NewInt(0)},
		},
		{
			"0 upload 10 download",
			args{upload: 0, download: 10},
			Bandwidth{sdkmath.NewInt(0), sdkmath.NewInt(10)},
		},
		{
			"10 upload 0 download",
			args{upload: 10, download: 0},
			Bandwidth{sdkmath.NewInt(10), sdkmath.NewInt(0)},
		},
		{
			"10 upload 10 download",
			args{upload: 10, download: 10},
			Bandwidth{sdkmath.NewInt(10), sdkmath.NewInt(10)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBandwidthFromInt64(tt.args.upload, tt.args.download); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBandwidthFromInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}
