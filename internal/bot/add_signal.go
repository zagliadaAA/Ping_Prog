package bot

import (
	"context"
	"errors"
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

	signals, err := b.signalUseCase.GetAllSignals(ctx, message.From.UserName)
	if err != nil {
		b.sendMessage(int(message.Chat.ID), fmt.Sprintf("❗Ошибка при получении адресов пользователя: %v", err))
		return
	}
	for _, signal := range signals {
		if signal.Address == address && signal.Port == port {
			err = errors.New("адрес уже существует")
			b.sendMessage(int(message.Chat.ID), fmt.Sprintf("❗Ошибка при добавлении адреса: %v", err))
			return
		}
	}

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

	port, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("❗️не удалось преобразовать порт в int")
	}

	if port < 0 || port > 65535 {
		return fmt.Errorf("❗Порт должен быть в диапазоне 0–65535")
	}

	return nil
}
