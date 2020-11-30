package commands

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/cobra"
	"log"
)

var setupCmd = &cobra.Command{
	Use:    "setup [botToken]",
	Args:   cobra.MaximumNArgs(1),
	Short:  "Perform initial setup",
	Long:   "Reveal latest messages send to bot to extract channel ID being used",
	PreRun: preloadConfig,
	Run: func(cmd *cobra.Command, args []string) {
		var token string
		if len(args) == 0 {
			fmt.Println("Please, provide bot token, received from https://t.me/BotFather official bot.")
			fmt.Println("It will be something like 238222314:BAjcF4IKGAIiL.")
			fmt.Println("Press ENTER when you are ready:")
			fmt.Print("> ")
			_, err := fmt.Scanln(&token)
			if err != nil {
				log.Fatalf("%s : while reading bot token string", err)
			}
		} else {
			token = args[0]
		}
		bot, err := tgbotapi.NewBotAPI(token)
		if err != nil {
			log.Fatalf("%s : while dialing api", err)
		}
		bot.Debug = false
		fmt.Printf(
			"We have authorized as bot %s #%v!\nNow please invite your bot %s to groups, where it should post notifications...\n",
			bot.Self.UserName,
			bot.Self.ID,
			bot.Self.UserName,
		)

		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60

		updates, err := bot.GetUpdatesChan(u)
		for update := range updates {
			if update.Message.NewChatMembers == nil {
				continue
			}

			members := *update.Message.NewChatMembers
			for _, newChatMember := range members {
				if newChatMember.IsBot {
					if newChatMember.ID == bot.Self.ID {
						fmt.Printf("Bot %s (#%v) was invited to chat %s (#%v) of type %s. It as us! Lets send message!\n",
							newChatMember.String(),
							newChatMember.ID,
							update.Message.Chat.Title,
							update.Message.Chat.ID,
							update.Message.Chat.Type,
						)
						greeting := tgbotapi.NewMessage(
							update.Message.Chat.ID,
							fmt.Sprintf("*Hello from telegramnotify!*\nWe can now send messages to %s *\"%s\"* (#%v)!!!\n",
								update.Message.Chat.Type,
								update.Message.Chat.Title,
								update.Message.Chat.ID,
							),
						)
						greeting.ParseMode = tgbotapi.ModeMarkdown
						_, err := bot.Send(greeting)
						if err != nil {
							fmt.Printf(
								"%s : while sending notification to chat %s %s (#%v)",
								err,
								update.Message.Chat.Type,
								update.Message.Chat.Title,
								update.Message.Chat.ID,
							)
							return
						}
						fmt.Printf("Confirmation is send to %s %s (#%v)",
							update.Message.Chat.Type,
							update.Message.Chat.Title,
							update.Message.Chat.ID,
						)

						cfg := *currentConfig
						cfg[update.Message.Chat.Title] = Sink{
							Token:  token,
							ChatID: update.Message.Chat.ID,
						}
						err = cfg.Save(PathToConfig)
						if err != nil {
							fmt.Printf("%s : while saving config to %s\n", err, PathToConfig)
							return
						}
						fmt.Printf("Sink %s is added to config at %s\n",
							update.Message.Chat.Title, PathToConfig)
						fmt.Printf("If you want to send messages to other groups, you can invite your bot %s into them.\n",
							bot.Self.UserName,
						)
						fmt.Printf("HINT: you can copy file %s into /etc/telegramnotify.json on linux machine to make config global",
							PathToConfig)
						fmt.Println("Press CTRL+C when you have finished adding groups to this bot!")
					}
				}
			}
		}
	},
}
