package main

import (
	"errors"
	"log"
	"os"
	"path"

	"github.com/cosmos/cosmos-sdk/server"
	servercmd "github.com/cosmos/cosmos-sdk/server/cmd"

	base "github.com/sentinel-official/sentinelhub/v12/types"
)

func HomeDir() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	dir = path.Join(dir, ".sentinelhub")

	return dir, nil
}

func main() {
	cfg := base.GetConfig()
	cfg.Seal()

	homeDir, err := HomeDir()
	if err != nil {
		log.Fatal(err)
	}

	cmd := NewRootCmd(homeDir)
	if err = servercmd.Execute(cmd, "SENTINELHUB", homeDir); err != nil {
		var e server.ErrorCode
		switch {
		case errors.As(err, &e):
			os.Exit(e.Code)
		default:
			os.Exit(1)
		}
	}
}
