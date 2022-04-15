package main

import (
	"log"

	"github.com/SakoDroid/telego/pkg/telegobot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5183055481:AAE7gMZhHWUDBFctFqIH4ip9ec83tW8lT9U")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	telegobot.NewBot()

}
