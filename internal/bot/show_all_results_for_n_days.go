package bot

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) showAllResultsForNDays(ctx context.Context, message *tgbotapi.Message) {
	var days int
	arg := message.CommandArguments()
	if arg == "" {
		days = 1
	} else {
		numberOfDays, err := strconv.Atoi(arg)
		if err != nil {
			b.sendMessage(int(message.Chat.ID), fmt.Sprintf("️не удалось распознать количество дней: %v", err))
			return
		}

		days = numberOfDays
	}

	// По userName получаем user
	user, err := b.userUseCase.GetByUserName(ctx, message.Chat.UserName)
	if err != nil {
		b.sendMessage(int(message.Chat.ID), fmt.Sprintf("Пользователь не найден: %v", err))
		return
	}

	results, err := b.resultUseCase.ShowAllResultsForNDays(ctx, user.ID, days)
	if err != nil {
		b.sendMessage(int(message.Chat.ID), fmt.Sprintf("Ошибка при получении данных: %v", err))
		return
	}

	if len(results) == 0 {
		b.sendMessage(int(message.Chat.ID), "📭 Нет результатов за указанный период")
		return
	}

	loc, err := time.LoadLocation("Asia/Vladivostok")
	if err != nil {
		loc = time.UTC // fallback
	}

	var builder strings.Builder
	for i, res := range results {
		status := "✅ Успех"
		if !res.Result {
			status = "❌ Ошибка"
		}

		created := res.CreatedAt.In(loc).Format("2006-01-02 15:04:05 MST")

		line := fmt.Sprintf("%d. %s:%d — %s (%s)\n", i+1, res.Address, res.Port, status, created)
		builder.WriteString(line)
	}

	text := builder.String()

	// Telegram ограничение — 4096 символов
	const chunkSize = 4000

	for i := 0; i < len(text); i += chunkSize {
		end := i + chunkSize
		if end > len(text) {
			end = len(text)
		}
		msg := tgbotapi.NewMessage(message.Chat.ID, text[i:end])
		_, err = b.api.Send(msg)
		if err != nil {
			b.sendMessage(int(message.Chat.ID), "❗Ошибка отправки результатов:")
			return
		}
	}
}
