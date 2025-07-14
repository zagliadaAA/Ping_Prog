package bot

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"ping_prog/internal/usecase/result_usecase"

	"github.com/go-ping/ping"
)

func (b *Bot) backgroundPingTask(ctx context.Context) {
	ticker := time.NewTicker(20 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("фоновый пинг остановлен")
			return
		case <-ticker.C:
			b.checkAllSignals(ctx)
		}
	}
}

func (b *Bot) checkAllSignals(ctx context.Context) {
	userSignalsMap, err := b.signalUseCase.GetActiveSignalsGroupedByUser(ctx)
	if err != nil {
		log.Printf("ошибка при получении списка пользовательских сигналов: %v", err)
		return
	}

	for userID, signals := range userSignalsMap {
		user, err := b.userUseCase.GetByID(ctx, userID)
		if err != nil {
			log.Printf("failed to get user %d: %v", userID, err)
			continue
		}

		for _, s := range signals {
			// Пингуем
			errPing, pinger := commandPing(s.Address)
			portOk := true

			if s.Port != 0 {
				_, portOk = checkPort(s.Address, s.Port)
			}

			if errPing != nil || pinger.Statistics().PacketLoss == 100 || !portOk {
				message := fmt.Sprintf("🚨 Хост %s:%d недоступен!", s.Address, s.Port)
				b.sendMessage(user.ChatID, message)
				b.sendResultToDB(ctx, false, pinger, s.ID, user.ID)
				continue
			}

			b.sendResultToDB(ctx, true, pinger, s.ID, user.ID)

		}
	}
}

func (b *Bot) sendResultToDB(ctx context.Context, result bool, statistic *ping.Pinger, IDSignal int, IDUser int) {
	err := b.resultUseCase.Create(ctx, result_usecase.CreateResultReq{
		Result:    result,
		Statistic: fmt.Sprintf("Потери: %.0f%%, Время: %v", statistic.Statistics().PacketLoss, statistic.Statistics().AvgRtt),
		IDSignal:  IDSignal,
		IDUser:    IDUser,
		CreatedAt: time.Now().UTC(),
	})
	if err != nil {
		log.Printf("ошибка при записи результата в бд: %v", err)
		return
	}
}

// команда ping
func commandPing(addr string) (error, *ping.Pinger) {
	pinger, err := ping.NewPinger(addr)
	if err != nil {
		return fmt.Errorf("ошибка создания pinger"), pinger
	}

	pinger.Count = 4
	pinger.Timeout = 2 * time.Second
	pinger.SetPrivileged(false)

	err = pinger.Run()
	if err != nil {
		return fmt.Errorf("ошибка запуска pinger"), pinger
	}

	return nil, pinger
}

// Функция для проверки доступности порта на сервере
func checkPort(addr string, port int) (error, bool) {
	// Формируем строку с адресом и портом
	address := fmt.Sprintf("%s:%d", addr, port)

	// Пытаемся подключиться по TCP
	conn, err := net.DialTimeout("tcp", address, 2*time.Second)
	if err != nil {
		return fmt.Errorf("ошибка подключения по TCP (для порта): %w", err), false // Возвращаем false, если не удаётся подключиться
	}
	defer conn.Close()

	return nil, true
}
