package bot

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"ping_prog/internal/usecase/signal_usecase"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) addSignal(ctx context.Context, message *tgbotapi.Message) {
	err := validateAddSignal(message)
	if err != nil {
		b.sendMessage(int(message.Chat.ID), fmt.Sprintf("Ошибка: %v", err))
		return
	}

	args := message.CommandArguments()
	parts := strings.Fields(args)
	address := parts[0]
	port, _ := strconv.Atoi(parts[1])

	// Добавляем адрес в бд
	err = b.signalUseCase.Create(ctx, signal_usecase.CreateSignalReq{
		Address:    address,
		Port:       port,
		UserChatId: int(message.Chat.ID),
	})
	if err != nil {
		b.sendMessage(int(message.Chat.ID), fmt.Sprintf("❗Ошибка при добавлении адреса: %v", err))
		return
	}

	b.sendMessage(int(message.Chat.ID), fmt.Sprintf("✅Адрес %s с портом %v добавлен!", address, port))
}

func validateAddSignal(message *tgbotapi.Message) error {
	args := message.CommandArguments()
	if args == "" {
		return fmt.Errorf("❗️не переданы адрес и порт, необходимо /add <адрес> <порт>")
	}

	parts := strings.Fields(args)
	if len(parts) != 2 {
		return fmt.Errorf("❗️недостаточно аргументов, необходимо /add <адрес> <порт>")
	}

	_, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("❗️не удалось преобразовать порт в int")
	}

	return nil
}
