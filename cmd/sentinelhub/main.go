package main

import (
	"log"
	"os"
	"path"

	servercmd "github.com/cosmos/cosmos-sdk/server/cmd"

	base "github.com/sentinel-official/sentinelhub/v13/types"
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
		os.Exit(1)
	}
}
