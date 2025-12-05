package v1

import (
	"slices"
	"strings"
)

func (s Status) String() string {
	switch s {
	case StatusActive:
		return "active"
	case StatusInactivePending:
		return "inactive_pending"
	case StatusInactive:
		return "inactive"
	default:
		return "unspecified"
	}
}

func (s Status) IsValid() bool {
	switch s {
	case StatusActive,
		StatusInactivePending,
		StatusInactive:
		return true
	default:
		return false
	}
}

func (s Status) Equal(v Status) bool {
	return s == v
}

func (s Status) IsOneOf(items ...Status) bool {
	return slices.ContainsFunc(items, s.Equal)
}

func StatusFromString(s string) Status {
	s = strings.ToLower(s)
	switch s {
	case "active":
		return StatusActive
	case "inactive_pending":
		return StatusInactivePending
	case "inactive":
		return StatusInactive
	default:
		return StatusUnspecified
	}
}
