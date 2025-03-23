package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) {
	switch message.Command() {
	case "add":
		b.addSignal(message)
	case "show_all":
		b.showAllSignals(message)
	case "start":
		b.helloUser(message)
	default:
		b.sendMessage(message.Chat.ID, "Неизвестная команда")
	}
}
