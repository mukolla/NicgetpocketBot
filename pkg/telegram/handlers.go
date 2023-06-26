package telegram

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
	"log"
	"net/url"
)

const commandStart = "start"

func (b *Bot) handleMessage(message *tgbotapi.Message) error {

	log.Printf("[%s] %s", message.From.UserName, message.Text)

	_, err := url.ParseRequestURI(message.Text)
	if err != nil {
		return errorInvalidUrl
	}

	accessToken, err := b.getAccessToken(message.Chat.ID)
	if err != nil {
		return errorAuthorized
	}

	if err := b.pocketClient.Add(context.Background(), pocket.AddInput{
		AccessToken: accessToken,
		URL:         message.Text,
	}); err != nil {
		return errorUnableToSave
	}

	return b.sendMessage(message, b.messages.Response.SavedSuccessfully)
}

func (b *Bot) sendMessage(message *tgbotapi.Message, messageText string) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, messageText)
	_, err := b.bot.Send(msg)
	return err
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
	_, err := b.getAccessToken(message.Chat.ID)
	if err != nil {
		return b.initAuthorizationProcess(message)
	}

	return errorAuthorized
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.Response.UnknownCommand+" ["+message.Command()+"]")
	_, error := b.bot.Send(msg)
	return error
}
