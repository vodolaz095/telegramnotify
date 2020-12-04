// +build !linux

package commands

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// FindConfig finds path to config file in most usual places
func FindConfig() (pathToConfig string) {
	cp, err := os.UserConfigDir()
	if err != nil {
		log.Fatalf("%s : while reading config file from %s", err, pathToConfig)
	}
	pathToConfig = filepath.Join(cp, "telegramnotify.json")
	if Verbose {
		fmt.Printf("Reading config from %s...", pathToConfig)
	}
	return
}
