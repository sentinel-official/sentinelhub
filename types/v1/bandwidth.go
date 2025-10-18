package v1

import (
	sdkmath "cosmossdk.io/math"
)

func NewBandwidth(upload, download sdkmath.Int) Bandwidth {
	return Bandwidth{
		Upload:   upload,
		Download: download,
	}
}

func (b Bandwidth) IsAnyNil() bool {
	return b.Upload.IsNil() || b.Download.IsNil()
}

func (b Bandwidth) IsAnyZero() bool {
	return b.Upload.IsZero() || b.Download.IsZero()
}

func (b Bandwidth) IsAllZero() bool {
	return b.Upload.IsZero() && b.Download.IsZero()
}

func (b Bandwidth) IsAnyNegative() bool {
	return b.Upload.IsNegative() || b.Download.IsNegative()
}

func (b Bandwidth) IsAllPositive() bool {
	return b.Upload.IsPositive() && b.Download.IsPositive()
}

func (b Bandwidth) Sum() sdkmath.Int {
	return b.Upload.Add(b.Download)
}

func (b Bandwidth) Add(v Bandwidth) Bandwidth {
	b.Upload = b.Upload.Add(v.Upload)
	b.Download = b.Download.Add(v.Download)

	return b
}

func (b Bandwidth) Sub(v Bandwidth) Bandwidth {
	b.Upload = b.Upload.Sub(v.Upload)
	b.Download = b.Download.Sub(v.Download)

	return b
}

func (b Bandwidth) IsAllLTE(v Bandwidth) bool {
	return b.Upload.LTE(v.Upload) && b.Download.LTE(v.Download)
}

func (b Bandwidth) IsAnyGT(v Bandwidth) bool {
	return b.Upload.GT(v.Upload) || b.Download.GT(v.Download)
}

func (b Bandwidth) CeilTo(pre sdkmath.Int) Bandwidth {
	if !pre.IsPositive() {
		return b
	}

	diff := NewBandwidth(
		pre.Sub(b.Upload.Mod(pre)),
		pre.Sub(b.Download.Mod(pre)),
	)

	if diff.Upload.Equal(pre) {
		diff.Upload = sdkmath.ZeroInt()
	}

	if diff.Download.Equal(pre) {
		diff.Download = sdkmath.ZeroInt()
	}

	return b.Add(diff)
}

func NewBandwidthFromInt64(upload, download int64) Bandwidth {
	return NewBandwidth(sdkmath.NewInt(upload), sdkmath.NewInt(download))
}
