package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Args:    cobra.MaximumNArgs(0),
	Short:   "list sinks",
	Long:    "list information about sinks currently known",
	PreRun:  preloadConfig,
	Example: "telegramnotify list",
	Run: func(cmd *cobra.Command, args []string) {
		if len(*currentConfig) == 0 {
			fmt.Printf("Is there config in %s?\n", PathToConfig)
		}
		for k, v := range *currentConfig {
			fmt.Printf("Config: Sunc %s (%v) is found!\n", k, v.ChatID)
		}
	},
}
