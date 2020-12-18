package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/cobra"

	"os"
)

// Version is version engraved
var Version string

// Subversion is version engraved
var Subversion string

// Verbose depicts how many messages application sends to STDOUT
var Verbose bool

// PathToConfig is path to config file
var PathToConfig string

var currentConfig *Config

var parseMode string

var rootCmd = &cobra.Command{
	Use:     "telegramnotify",
	Short:   "Console application to send notifications into telegram channels via bot-api",
	Long:    "Console application to send notifications into telegram channels via bot-api",
	PreRun:  preloadConfig,
	Version: fmt.Sprintf("%s. Build: %s. Homepage: https://github.com/vodolaz095/telegramnotify", Version, Subversion),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func preloadConfig(cmd *cobra.Command, args []string) {
	cfg, err := LoadConfigFromFile(PathToConfig)
	currentConfig = &cfg
	if err != nil {
		fmt.Printf("Error: %s - while reading config from %s\n", err, PathToConfig)
		os.Exit(1)
		return
	}
	if Verbose {
		fmt.Printf("Config: reading from file %s\n", PathToConfig)
		if len(*currentConfig) == 0 {
			fmt.Printf("Is there config in %s?\n", PathToConfig)
		}

		for k := range *currentConfig {
			fmt.Printf("Config: Sunc %s is found!\n", k)
		}
	}
}

func init() {
	rootCmd.AddCommand(setupCmd)
	rootCmd.AddCommand(testCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(fileCmd)
	rootCmd.AddCommand(imageCmd)
	rootCmd.AddCommand(textCmd)
	textCmd.PersistentFlags().StringVarP(&parseMode, "mode", "m", tgbotapi.ModeMarkdown, "Message format - can be Markdown, MarkdownV2, HTML")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringVarP(&PathToConfig, "config", "c", FindConfig(), "path to config file")
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
