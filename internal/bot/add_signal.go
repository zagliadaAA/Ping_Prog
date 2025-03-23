package bot

import (
	"fmt"
	"strconv"
	"strings"

	"ping_prog/internal/usecase/signal_usecase"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) addSignal(message *tgbotapi.Message) {
	err := validateAddSignal(message)
	if err != nil {
		b.sendMessage(message.Chat.ID, fmt.Sprintf("Ошибка: %v", err))
		return
	}

	args := message.CommandArguments()
	parts := strings.Fields(args)
	address := parts[0]
	port, _ := strconv.Atoi(parts[1])

	// Добавляем адрес в бд
	err = b.signalUC.Create(signal_usecase.CreateSignalReq{
		Address: address,
		Port:    port,
	})
	if err != nil {
		b.sendMessage(message.Chat.ID, fmt.Sprintf("❗Ошибка при добавлении адреса: %v", err))
		return
	}

	b.sendMessage(message.Chat.ID, fmt.Sprintf("✅Адрес %s с портом %v добавлен!", address, port))
}

func validateAddSignal(message *tgbotapi.Message) error {
	args := message.CommandArguments()
	if args == "" {
		return fmt.Errorf("❗️не переданы адрес и порт, необходимо /signals <адрес> <порт>")
	}

	parts := strings.Fields(args)
	if len(parts) != 2 {
		return fmt.Errorf("❗️недостаточно аргументов, необходимо /signals <адрес> <порт>")
	}

	_, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("❗️не удалось преобразовать порт в int")
	}

	return nil
}
