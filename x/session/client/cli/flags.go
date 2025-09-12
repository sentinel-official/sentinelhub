package cli

import (
	"encoding/base64"

	"github.com/spf13/pflag"
)

const (
	flagSignature = "signature"
)

func GetSignature(flags *pflag.FlagSet) ([]byte, error) {
	s, err := flags.GetString(flagSignature)
	if err != nil {
		return nil, err
	}

	if s == "" {
		return nil, nil
	}

	return base64.StdEncoding.DecodeString(s)
}
