package commands

import (
	"bufio"
	"bytes"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

var textCmd = &cobra.Command{
	Use: "text [message] [sink]",
	Aliases: []string{
		"send",
		"plain",
	},
	Args:   cobra.MaximumNArgs(2),
	Short:  "Send plain text to sink provided",
	Long:   "Send plain text to sink provided",
	PreRun: preloadConfig,
	Example: strings.Join([]string{
		"telegramnotify send \"Hello, this is test plain text message to be send to sink >>>work<<<!\" work",
		"telegramnotify send -m Markdown \"*Hello from telegramnotify*\nThis is markdown formatted notification.\" work",
		"telegramnotify send -m HTML '<a href=\"https://www.rt.com/\">Stay up to date with latest news!</a>' work",
		"uptime | telegramnotify send work",
	}, "\n"),
	Run: func(cmd *cobra.Command, args []string) {
		var sinkName, input string
		readInput := false
		cfg := *currentConfig
		var sink Sink
		switch len(args) {
		case 2:
			input = args[0]
			sinkName = args[1]
			break
		case 1:
			if len(cfg) == 1 { // only one sink here
				for k, v := range cfg {
					if Verbose {
						fmt.Printf("Sending to default sink %s (%v_...", k, v.ChatID)
					}
					sinkName = k
				}
				input = args[0]
			} else {
				fmt.Println("There is more than one sink present, refusing to read input")
				return
			}
			break
		case 0:
			if len(cfg) == 1 { // only one sink here
				for k, v := range cfg {
					if Verbose {
						fmt.Printf("Sending to default sink %s (%v)...\n", k, v.ChatID)
					}
					sinkName = k
				}
			} else {
				fmt.Println("There are more than one sink available. Where to send STDIN contents?")
				return
			}
			readInput = true
			break
		}
		sink, found := cfg[sinkName]
		if !found {
			fmt.Printf("Unable to find sink %s in file %s!\n", sinkName, PathToConfig)
			return
		}
		if readInput {
			b := bytes.Buffer{}
			scanner := bufio.NewScanner(cmd.InOrStdin())
			for scanner.Scan() {
				b.WriteString(scanner.Text())
			}
			input = b.String()
			if Verbose {
				fmt.Println("Preparing to send:", input)
			}
			if scanner.Err() != nil {
				log.Fatalf("%s : while reading data", scanner.Err())
			}
		}
		if len(input) == 0 {
			fmt.Println("Unable to send empty message!")
			return
		}
		bot, err := tgbotapi.NewBotAPI(sink.Token)
		if err != nil {
			log.Panic(err)
		}
		bot.Debug = Verbose
		fmt.Printf("Authorized on account %s...\n", bot.Self.UserName)
		payload := tgbotapi.NewMessage(
			sink.ChatID,
			input,
		)
		payload.ParseMode = parseMode
		msg, err := bot.Send(payload)
		if err != nil {
			log.Fatalf("%s : while sending test message to channel %s", err, sinkName)
		}
		fmt.Printf("Message is send to chat %s #%v!\n", msg.Chat.Title, msg.Chat.ID)
	},
}
