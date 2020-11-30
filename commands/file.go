package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/cobra"
	"log"
)

var fileCmd = &cobra.Command{
	Use:     "file [pathToFile] [sink]",
	Args:    cobra.RangeArgs(1, 2),
	Short:   "Upload file to sink provided",
	Long:    "Upload file to sink provided",
	Example: "telegramnotify file /etc/passwd work",
	Aliases: []string{
		"document", "upload", "share",
	},
	PreRun: preloadConfig,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := *currentConfig
		var sink Sink
		if len(args) == 1 && len(cfg) == 1 { // only one sink here
			for k, v := range cfg {
				if Verbose {
					fmt.Printf("Sending to default sink %s (%v_...", k, v.ChatID)
				}
				sink = cfg[k]
			}
		} else {
			s, ok := cfg[args[1]]
			if !ok {
				fmt.Printf("Unable to find sink %s in file %s!\n", args[1], PathToConfig)
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
		payload := tgbotapi.NewDocumentUpload(
			sink.ChatID,
			args[0],
		)
		msg, err := bot.Send(payload)
		if err != nil {
			log.Fatalf("%s : while sending test message to channel %s", err, args[0])
		}
		fmt.Printf("File %s is uploaded to chat %s #%v!\n",
			args[0],
			msg.Chat.Title, msg.Chat.ID)
	},
}
