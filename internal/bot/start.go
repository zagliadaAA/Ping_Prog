package bot

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) Start() {
	ctx := context.Background()

	// запуск пинга фоном при старте бота
	go b.backgroundPingTask(ctx)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.api.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		userName := update.Message.From.UserName
		chatID := update.Message.Chat.ID

		ctx := context.Background()

		_, err := b.userUseCase.Create(ctx, userName, int(chatID))
		if err != nil {
			log.Println(err)
		}

		//log.Printf("User: %+v", user)

		if update.Message.IsCommand() {
			b.handleCommand(ctx, update.Message)
		}
	}
}
