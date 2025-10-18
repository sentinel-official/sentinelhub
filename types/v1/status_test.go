package v1

import (
	"testing"
)

func TestStatusFromString(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want Status
	}{
		{
			"empty",
			args{
				s: "",
			},
			StatusUnspecified,
		},
		{
			"invalid",
			args{
				s: "invalid",
			},
			StatusUnspecified,
		},
		{
			"unspecified",
			args{
				s: "unspecified",
			},
			StatusUnspecified,
		},
		{
			"active",
			args{
				s: "active",
			},
			StatusActive,
		},
		{
			"inactive pending",
			args{
				s: "inactive_pending",
			},
			StatusInactivePending,
		},
		{
			"inactive",
			args{
				s: "inactive",
			},
			StatusInactive,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StatusFromString(tt.args.s); got != tt.want {
				t.Errorf("StatusFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatus_Equal(t *testing.T) {
	type args struct {
		v Status
	}

	tests := []struct {
		name string
		s    Status
		args args
		want bool
	}{
		{
			"unspecified and unspecified",
			StatusUnspecified,
			args{
				v: StatusUnspecified,
			},
			true,
		},
		{
			"unspecified and active",
			StatusUnspecified,
			args{
				v: StatusActive,
			},
			false,
		},
		{
			"unspecified and inactive pending",
			StatusUnspecified,
			args{
				v: StatusInactivePending,
			},
			false,
		},
		{
			"unspecified and inactive",
			StatusUnspecified,
			args{
				v: StatusInactive,
			},
			false,
		},
		{
			"active and unspecified",
			StatusActive,
			args{
				v: StatusUnspecified,
			},
			false,
		},
		{
			"active and active",
			StatusActive,
			args{
				v: StatusActive,
			},
			true,
		},
		{
			"active and inactive pending",
			StatusActive,
			args{
				v: StatusInactivePending,
			},
			false,
		},
		{
			"active and inactive",
			StatusActive,
			args{
				v: StatusInactive,
			},
			false,
		},
		{
			"inactive pending and unspecified",
			StatusInactivePending,
			args{
				v: StatusUnspecified,
			},
			false,
		},
		{
			"inactive pending and active",
			StatusInactivePending,
			args{
				v: StatusActive,
			},
			false,
		},
		{
			"inactive pending and inactive pending",
			StatusInactivePending,
			args{
				v: StatusInactivePending,
			},
			true,
		},
		{
			"inactive pending and inactive",
			StatusInactivePending,
			args{
				v: StatusInactive,
			},
			false,
		},
		{
			"inactive and unspecified",
			StatusInactive,
			args{
				v: StatusUnspecified,
			},
			false,
		},
		{
			"inactive and active",
			StatusInactive,
			args{
				v: StatusActive,
			},
			false,
		},
		{
			"inactive and inactive pending",
			StatusInactive,
			args{
				v: StatusInactivePending,
			},
			false,
		},
		{
			"inactive and inactive",
			StatusInactive,
			args{
				v: StatusInactive,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Equal(tt.args.v); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatus_IsValid(t *testing.T) {
	tests := []struct {
		name string
		s    Status
		want bool
	}{
		{
			"unspecified",
			StatusUnspecified,
			false,
		},
		{
			"active",
			StatusActive,
			true,
		},
		{
			"inactive pending",
			StatusInactive,
			true,
		},
		{
			"inactive",
			StatusInactive,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatus_String(t *testing.T) {
	tests := []struct {
		name string
		s    Status
		want string
	}{
		{
			"minus one",
			Status(-1),
			"unspecified",
		},
		{
			"zero",
			Status(0),
			"unspecified",
		},
		{
			"one",
			Status(1),
			"active",
		},
		{
			"two",
			Status(2),
			"inactive_pending",
		},
		{
			"three",
			Status(3),
			"inactive",
		},
		{
			"four",
			Status(4),
			"unspecified",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
