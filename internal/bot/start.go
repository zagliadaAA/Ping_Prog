package bot

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.api.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		username := update.Message.From.UserName
		chatID := update.Message.Chat.ID

		ctx := context.Background()

		user, err := b.userUC.RegisterOrGet(ctx, username, chatID)
		if err != nil {
			log.Printf("failed to register user: %v", err)
			continue
		}

		log.Printf("User: %+v", user)

		if update.Message.IsCommand() {
			b.handleCommand(update.Message)
		}

		// Отправляем ответ с тем же текстом
		/*msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Вы написали: %s", update.Message.Text))
		b.api.Send(msg)*/
	}
}
