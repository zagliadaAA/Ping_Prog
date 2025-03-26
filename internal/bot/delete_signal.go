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

	// удалить все записи из таблицы results
	user, err := b.userUseCase.GetByUserName(ctx, message.From.UserName)
	if err != nil {
		b.sendMessage(int(message.Chat.ID), fmt.Sprintf("❗Ошибка при получении пользователя: %v", err))
		return
	}

	signal, err := b.signalUseCase.GetByID(ctx, id, user.ID)
	if err != nil {
		b.sendMessage(int(message.Chat.ID), fmt.Sprintf("❗Ошибка при получении сигнала: %v", err))
		return
	}

	err = b.resultUseCase.DeleteResultsForSignal(ctx, signal)
	//получить сигнал по id для пользователя
	//удалить все результаты для сигнала

	// удалить адрес из таблицы signals
	err = b.signalUseCase.Delete(ctx, id)
	if err != nil {
		b.sendMessage(int(message.Chat.ID), fmt.Sprintf("❗Ошибка при удалении адреса: %v", err))
		return
	}

	b.sendMessage(int(message.Chat.ID), fmt.Sprintf("✅Адрес с портом удалены!"))
}
