// +build !linux

package commands

import (
	"fmt"
	"os"
	"path/filepath"
)

// FindConfig finds path to config file in most usual places
func FindConfig() (pathToConfig string) {
	pathToConfig, err = filepath.Join(os.UserConfigDir(), "telegramnotify.json")
	if err != nil {
		fmt.Printf("%s : while reading config file from %s", err, pathToConfig)
		os.Exit(1)
	}

	if Verbose {
		fmt.Printf("Reading config from %s...", pathToConfig)
	}
	return
}
