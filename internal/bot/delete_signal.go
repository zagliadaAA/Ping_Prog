package bot

import (
	"context"
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) deleteSignal(ctx context.Context, message *tgbotapi.Message) {
	args := message.CommandArguments()
	id, err := strconv.Atoi(args)
	if err != nil {
		b.sendMessage(int(message.Chat.ID), fmt.Sprintf("❗Ошибка при удалении адреса: %v", err))
		b.sendMessage(int(message.Chat.ID), fmt.Sprintf("❗Введите id записи для удаления"))
		return
	}

	// Добавляем адрес в бд
	err = b.signalUseCase.Delete(ctx, id)
	if err != nil {
		b.sendMessage(int(message.Chat.ID), fmt.Sprintf("❗Ошибка при удалении адреса: %v", err))
		return
	}

	b.sendMessage(int(message.Chat.ID), fmt.Sprintf("✅Адрес с портом удалены!"))
}
