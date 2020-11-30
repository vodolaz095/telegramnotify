package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/cobra"
	"log"
)

var testCmd = &cobra.Command{
	Use:     "test [sink]",
	Args:    cobra.MaximumNArgs(1),
	Short:   "Send test message to channel provided",
	Long:    "Send test message to channel provided",
	Example: "telegramnotify test work",
	PreRun:  preloadConfig,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := *currentConfig
		var sink Sink
		if len(args) == 0 && len(cfg) == 1 { // only one sink here
			for k, v := range cfg {
				if Verbose {
					fmt.Printf("Sending to default sink %s (%v_...", k, v.ChatID)
				}
				sink = cfg[k]
			}
		} else {
			s, ok := cfg[args[0]]
			if !ok {
				fmt.Printf("Unable to find sink %s in file %s!\n", args[0], PathToConfig)
				return
			}
			sink = s
		}

		bot, err := tgbotapi.NewBotAPI(sink.Token)
		if err != nil {
			log.Panic(err)
		}
		bot.Debug = Verbose
		fmt.Printf("Authorized on account %s...\n", bot.Self.UserName)
		payload := tgbotapi.NewMessage(
			sink.ChatID,
			"*Hello from telegramnotify*\nThis is test notification. We hope you have received it..",
		)
		payload.ParseMode = tgbotapi.ModeMarkdown
		msg, err := bot.Send(payload)
		if err != nil {
			log.Fatalf("%s : while sending test message to channel %s", err, args[0])
		}
		fmt.Printf("Message delivered to chat %s #%v!\n", msg.Chat.Title, msg.Chat.ID)
	},
}
