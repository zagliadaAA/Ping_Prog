package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) statusBot(message *tgbotapi.Message) {
	mes := fmt.Sprintf("То, что мертво, умереть не может — оно лишь восстает вновь, "+
		"сильнее и крепче, чем прежде.\n ®️Джордж Мартин\n\n%s, я готов работать 💪", message.From.UserName)

	b.sendMessage(int(message.Chat.ID), mes)
}
