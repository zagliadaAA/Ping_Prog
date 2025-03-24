package bot

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) showAllSignals(ctx context.Context, message *tgbotapi.Message) {
	signals, err := b.signalUseCase.GetAllSignals(ctx, message.Chat.UserName)
	if err != nil {
		b.sendMessage(int(message.Chat.ID), fmt.Sprintf("❗Ошибка при выводе списка адресов: %v", err))
		return
	}

	for _, val := range signals {
		b.sendMessage(int(message.Chat.ID), fmt.Sprintf("ID: %v адрес: %s порт: %v\n", val.ID, val.Address, val.Port))
	}
}
