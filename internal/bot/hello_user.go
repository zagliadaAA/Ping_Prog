package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) helloUser(message *tgbotapi.Message) {
	mes := fmt.Sprintf("🎉Привет, %s🎉\nМоя задача пинговать какой-либо сервер или сервис на нем\n"+
		"Чтобы я запомнил адрес компьютера, введи команду /add адрес порт", message.From.UserName)

	b.sendMessage(message.Chat.ID, mes)
}
