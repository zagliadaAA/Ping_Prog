package bot

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) handleCommand(ctx context.Context, message *tgbotapi.Message) {
	switch message.Command() {
	case "add": // добавление адреса и порта в бд
		b.addSignal(ctx, message)
	case "show_all": // показать все добавленные адреса
		b.showAllSignals(ctx, message)
	case "status": // статус бота
		b.statusBot(message)
	case "delete": // удаление address ping из таблицы signals
		b.deleteSignal(ctx, message)
	case "start": // приветствие при старте бота
		b.helloUser(message)
	default: // штатное сообщение
		b.sendMessage(int(message.Chat.ID), "Неизвестная команда")
	}
}
