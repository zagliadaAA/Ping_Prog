package bot

import (
	"ping_prog/internal/usecase/signal_usecase"
	"ping_prog/internal/usecase/user_usecase"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	api      *tgbotapi.BotAPI
	userUC   *user_usecase.UseCase
	signalUC *signal_usecase.UseCase
}

type signalUseCase interface {
	Create(req signal_usecase.CreateSignalReq) error
}

func NewBot(token string, userUC *user_usecase.UseCase, signalUC *signal_usecase.UseCase) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &Bot{
		api:      api,
		userUC:   userUC,
		signalUC: signalUC,
	}, nil
}
