package main

import (
	"github.com/boltdb/bolt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mukolla/nicgetpocketbot/pkg/repository"
	"github.com/mukolla/nicgetpocketbot/pkg/repository/boltdb"
	"github.com/mukolla/nicgetpocketbot/pkg/server"
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

	db, err := initDb()
	if err != nil {
		log.Fatal(err)
	}

	tokenRepository := boltdb.NewTokenRepository(db)

	telegramBot := telegram.NewBot(bot, pocketClient, tokenRepository, "http://localhost:8183/")
	authorizationServer := server.NewAuthorizationServer(pocketClient, tokenRepository, "https://t.me/Nicgetpocket_bot")

	go func() {
		if err := telegramBot.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := authorizationServer.Start(); err != nil {
		log.Fatal(err)
	}
}

func initDb() (*bolt.DB, error) {
	db, err := bolt.Open("pocketbook.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(repository.AccessToken))
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists([]byte(repository.RequestToken))
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return db, err
}
