package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) helloUser(message *tgbotapi.Message) {
	mes := fmt.Sprintf("🎉Привет, %s🎉\nМоя задача пинговать какой-либо сервер или сервис на нем\n"+
		"1) Чтобы я запомнил адрес компьютера, введи команду /add адрес порт\n"+
		"2) Если хочешь пинговать только компьютер, в поле порт поставь 0 ", message.From.UserName)

	b.sendMessage(int(message.Chat.ID), mes)
}
