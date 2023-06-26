package telegram

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mukolla/nicgetpocketbot/pkg/repository"
)

func (b *Bot) initAuthorizationProcess(message *tgbotapi.Message) error {
	authLink, err := b.generateAuthorizationLink(message.Chat.ID)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(b.messages.Response.Start, authLink))
	_, error := b.bot.Send(msg)
	return error
}

func (b *Bot) getAccessToken(chatID int64) (string, error) {
	return b.tokenRepository.Get(chatID, repository.AccessToken)
}

func (b *Bot) generateAuthorizationLink(chatID int64) (string, error) {
	redirectURL := b.generateRedirectURL(chatID, b.redirectURL)

	token, err := b.pocketClient.GetRequestToken(context.Background(), redirectURL)
	if err != nil {
		return "", err
	}

	if err := b.tokenRepository.Save(chatID, token, repository.RequestToken); err != nil {
		return "", err
	}

	return b.pocketClient.GetAuthorizationURL(token, redirectURL)
}

func (b *Bot) generateRedirectURL(chatID int64, redirectURL string) string {
	return fmt.Sprintf("%s?chat_id=%d", redirectURL, chatID)
}
