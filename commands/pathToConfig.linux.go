// +build linux

package commands

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

// FindConfig finds path to config file in most usual places
func FindConfig() (pathToConfig string) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("%s: while trying to read config from %s", err, pathToConfig)
			os.Exit(0)
		}
	}()

	pathToConfig = filepath.Join("/", "etc", "telegramnotify.json")
	_, err := os.Stat(pathToConfig)
	if err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}
	} else {
		if Verbose {
			fmt.Printf("Reading config from %s...", pathToConfig)
		}
		return
	}
	pathToConfig = "~/.config/telegramnotify.json"

	me, err := user.Current()
	if err != nil {
		panic(err)
	}
	pathToConfig = filepath.Join(me.HomeDir, ".config", "telegramnotify.json")
	_, err = os.Stat(pathToConfig)
	if err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}
	} else {
		if Verbose {
			fmt.Printf("Reading config from %s...", pathToConfig)
		}
		return
	}
	pathToConfig = "./telegramnotify.json"
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	pathToConfig = filepath.Join(dir, "telegramnotify.json")
	if Verbose {
		fmt.Printf("Reading config from %s...", pathToConfig)
	}
	return
}
