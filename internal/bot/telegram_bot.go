package bot

import (
	"context"

	"ping_prog/internal/domain"
	"ping_prog/internal/usecase/result_usecase"
	"ping_prog/internal/usecase/signal_usecase"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	api           *tgbotapi.BotAPI
	signalUseCase signalUseCase
	userUseCase   userUseCase
	resultUseCase resultUseCase
}

type signalUseCase interface {
	Create(ctx context.Context, req signal_usecase.CreateSignalReq) error
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, signalID int, userID int) (*domain.Signal, error)
	GetAllSignals(ctx context.Context, userName string) ([]domain.Signal, error)
	GetActiveSignalsGroupedByUser(ctx context.Context) (map[int][]domain.Signal, error)
}

type userUseCase interface {
	Create(ctx context.Context, userName string, chatID int) (*domain.User, error)
	GetByUserName(ctx context.Context, userName string) (*domain.User, error)
	GetByID(ctx context.Context, userID int) (*domain.User, error)
}

type resultUseCase interface {
	Create(ctx context.Context, req result_usecase.CreateResultReq) error
	DeleteResultsForSignal(ctx context.Context, signal *domain.Signal) error
	ShowAllResultsForNDays(ctx context.Context, userID int, days int) ([]*result_usecase.ResultView, error)
}

func NewBot(token string, signalUseCase signalUseCase, userUseCase userUseCase, resultUseCase resultUseCase) *Bot {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil
	}

	return &Bot{
		api:           api,
		signalUseCase: signalUseCase,
		userUseCase:   userUseCase,
		resultUseCase: resultUseCase,
	}
}
