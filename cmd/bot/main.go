package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mukolla/nicgetpocketbot/pkg/telegram"
	"github.com/zhashkevych/go-pocket-sdk"
	"log"
)

func main() {
	//Nicgetpocket_bot
	//http://t.me/Nicgetpocket_bot

	bot, err := tgbotapi.NewBotAPI("6003962422:AAHicNfe9XlhAmg7I8wsdTgOaLX9fUvQGAk")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	pocketClient, err := pocket.NewClient("107726-fa8cfa6223e05b23dcbf901")
	if err != nil {
		log.Fatal(err)
	}

	telegramBot := telegram.NewBot(bot, pocketClient, "http://localhost/")
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}
