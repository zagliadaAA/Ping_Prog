package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) showAllSignals(message *tgbotapi.Message) {
	signals, err := b.signalUC.GetAllSignals()
	if err != nil {
		b.sendMessage(message.Chat.ID, fmt.Sprintf("❗Ошибка при выводе списка адресов: %v", err))
		return
	}

	for _, val := range signals {
		b.sendMessage(message.Chat.ID, fmt.Sprintf("ID: %v адрес: %s порт: %v\n", val.ID, val.Address, val.Port))
	}
}
