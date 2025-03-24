package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) sendMessage(chatID int, text string) {
	msg := tgbotapi.NewMessage(int64(chatID), text)
	_, err := b.api.Send(msg)
	if err != nil {
		log.Printf("ошибка отправки сообщения: %v", err)
	}
}
