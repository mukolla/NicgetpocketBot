package telegram

import (
	"context"
	"fmt"
)

func (b *Bot) generateAuthorizationLink(chatID int64) (string, error) {
	redirectURL := b.generateRedirectURL(chatID, b.redirectURL)

	token, err := b.pocketClient.GetRequestToken(context.Background(), redirectURL)
	if err != nil {
		return "", err
	}

	return b.pocketClient.GetAuthorizationURL(token, redirectURL)
}

func (b *Bot) generateRedirectURL(chatID int64, redirectURL string) string {
	return fmt.Sprintf("%s?chat_id=%d", redirectURL, chatID)
}
