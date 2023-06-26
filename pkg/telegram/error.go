package telegram

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	errorInvalidUrl   = errors.New("URL is invalid")
	errorAuthorized   = errors.New("user is not Authorized")
	errorUnableToSave = errors.New("unable to save")
)

func (b *Bot) handleError(chatID int64, err error) error {

	msg := tgbotapi.NewMessage(chatID, "")

	switch err {
	case errorInvalidUrl:
		msg.Text = b.messages.Errors.InvalidURL
	case errorUnableToSave:
		msg.Text = b.messages.Errors.UnableToSave
	case errorAuthorized:
		msg.Text = b.messages.Errors.Unauthorized
	default:
		msg.Text = b.messages.Errors.Default
	}

	b.bot.Send(msg)

	return err
}
