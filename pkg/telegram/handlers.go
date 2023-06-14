package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const (
	commandStart        = "start"
	replayStartTemplate = "Привіт! Для того щоб зберігати ссилки у своєму Pocket аккаунті тобі потрібно, перейти по ссилці:\n%s"
)

func (b *Bot) handleMessage(message *tgbotapi.Message) {

	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	//msg.ReplyToMessageID = message.MessageID

	b.bot.Send(msg)
}

func (b *Bot) handleCommand(message *tgbotapi.Message) error {

	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {

	authLink, err := b.generateAuthorizationLink(message.Chat.ID)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(replayStartTemplate, authLink))
	_, error := b.bot.Send(msg)
	return error
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "You send command unknown command ["+message.Command()+"]")
	_, error := b.bot.Send(msg)
	return error
}
