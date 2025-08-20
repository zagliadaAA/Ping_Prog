package service_provider

import (
	"log"
	"os"

	"ping_prog/internal/bot"
)

func (sp *ServiceProvider) GetTelegramBot() *bot.Bot {
	/*err := godotenv.Load()
	if err != nil {
		log.Fatalf("ошибка загрузки .env файла")
	}*/

	telegramToken := os.Getenv("TokenTelegramBot")
	if telegramToken == "" {
		log.Fatal("не удалось получить телеграм токен из .env файла")
	}

	if sp.telegramBot == nil {
		sp.telegramBot = bot.NewBot(telegramToken, sp.GetSignalUseCase(), sp.GetUserUseCase(), sp.GetResultUsecase())
	}

	return sp.telegramBot
}
