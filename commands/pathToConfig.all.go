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
		log.Fatalf("%s : while finding config directory for user via os.UserConfigDir", err)
	}
	pathToConfig = filepath.Join(cp, "telegramnotify.json")
	if Verbose {
		fmt.Printf("Reading config from %s...", pathToConfig)
	}
	return
}
