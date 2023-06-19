package telegram

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
	"log"
	"net/url"
)

const (
	commandStart            = "start"
	replayStartTemplate     = "Привіт! Для того щоб зберігати ссилки у своєму Pocket аккаунті тобі потрібно, перейти по ссилці:\n%s"
	replayAlreadyAuthorized = "Привіт, ти вже авторизований присилай посилання, а я його збережу."
)

func (b *Bot) handleMessage(message *tgbotapi.Message) error {

	log.Printf("[%s] %s", message.From.UserName, message.Text)

	_, err := url.ParseRequestURI(message.Text)
	if err != nil {
		return b.sendMessage(message, "Це посилання не валідне!")
	}

	accessToken, err := b.getAccessToken(message.Chat.ID)
	if err != nil {
		return b.sendMessage(message, "Ви не атворизовані, використання команду /start для авторизації")
	}

	if err := b.pocketClient.Add(context.Background(), pocket.AddInput{
		AccessToken: accessToken,
		URL:         message.Text,
	}); err != nil {
		return b.sendMessage(message, "Халлепа, викиникла помилка при зберженні посилання. Спробуйте пізніше.")
	}

	return b.sendMessage(message, "Посилання успішно збереженно")
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

	msg := tgbotapi.NewMessage(message.Chat.ID, replayAlreadyAuthorized)
	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "You send command unknown command ["+message.Command()+"]")
	_, error := b.bot.Send(msg)
	return error
}
