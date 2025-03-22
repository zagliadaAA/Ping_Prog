package main

import (
	"net/http"

	"ping_prog/cmd/service_provider"
)

func main() {
	sp := service_provider.NewServiceProvider()

	// запуск сервера
	err := http.ListenAndServe(":8080", sp.GetRoutes())
	if err != nil {
		panic(err)
	}

	/*err := godotenv.Load()
	if err != nil { //письмо не отправится
		log.Fatal("Error loading .env file")
	}

	addr := os.Getenv("Address")
	port, err := strconv.Atoi(os.Getenv("Port"))

	// BOT-------------------------------------------------
	var (
		chatID      int64
		chatIDLock  sync.Mutex
		chatIDFile  = os.Getenv("ChatIDFileName")
		botToken    = os.Getenv("TokenTelegramBot")
		pingAddress = os.Getenv("Address")
	)

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	log.Println("🤖 Бот запущен:", bot.Self.UserName)

	// Загружаем сохранённый chatID (если есть)
	chatID = loadChatIDFromFile(chatIDFile)

	// Настраиваем обновления
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	// Запускаем фон для авто-пинга
	go backgroundPing(bot, &chatIDLock, chatID, pingAddress)

	for update := range updates {
		if update.Message == nil || !update.Message.IsCommand() {
			continue
		}

		chatIDLock.Lock()
		if chatID == 0 {
			chatID = update.Message.Chat.ID
			saveChatIDToFile(chatID, chatIDFile)
			log.Println("✅ Сохранил chatID:", chatID)
		}
		chatIDLock.Unlock()

		switch update.Message.Command() {
		case "ping":
			err, pinger := commandPing(pingAddress)
			var msgText string
			if err != nil {
				msgText = "❌ Ошибка пинга: " + err.Error()
			} else {
				stats := pinger.Statistics()
				msgText = fmt.Sprintf("📡 Ping result:\nПотери: %.0f%%\nВремя: %v", stats.PacketLoss, stats.AvgRtt)
				if stats.PacketLoss == 100 {
					alert := tgbotapi.NewMessage(chatID, "🚨 Потеря 100% пакетов! Хост недоступен.")
					bot.Send(alert)
				}
			}
			msg := tgbotapi.NewMessage(chatID, msgText)
			bot.Send(msg)

		case "status":
			msg := tgbotapi.NewMessage(chatID, "📊 Всё работает стабильно.")
			bot.Send(msg)

		default:
			msg := tgbotapi.NewMessage(chatID, "Неизвестная команда 🤖")
			bot.Send(msg)
		}
	}

	//------------------------------------------------------------------

	err, pinger := commandPing(addr)
	if err != nil {
		sendEmail(err, pinger, false)
		os.Exit(0) //завершение успехом
	} else if pinger.Statistics().PacketLoss == 100 {
		sendEmail(fmt.Errorf("пакеты не доставлены"), pinger, false)
		os.Exit(0)
	} else {
		err, portStatus := checkPort(addr, port)
		if err != nil {
			sendEmail(err, pinger, false)
			os.Exit(0)
		}

		if portStatus == false {
			sendEmail(err, pinger, false)
		}
	}*/
}

/*
// Фоновый пинг с тревогой
func backgroundPing(bot *tgbotapi.BotAPI, chatIDLock *sync.Mutex, chatID int64, pingAddress string) {
	for {
		time.Sleep(30 * time.Second)

		chatIDLock.Lock()
		currentChatID := chatID
		chatIDLock.Unlock()

		if currentChatID == 0 {
			continue
		}

		err, pinger := commandPing(pingAddress)
		if err != nil {
			log.Println("Ошибка пинга:", err)
			continue
		}

		if pinger.Statistics().PacketLoss == 100 {
			alert := tgbotapi.NewMessage(currentChatID, "🚨 Потеря 100% пакетов! Хост недоступен.")
			bot.Send(alert)
		}
	}
}

// Загружаем chatID из файла (если есть)
func loadChatIDFromFile(chatIDFile string) int64 {
	data, err := os.ReadFile(chatIDFile)
	if err != nil {
		return 0
	}
	var id int64
	fmt.Sscanf(string(data), "%d", &id)
	return id
}

// Сохраняем chatID в файл
func saveChatIDToFile(id int64, chatIDFile string) {
	f, err := os.Create(chatIDFile)
	if err == nil {
		defer f.Close()
		fmt.Fprintf(f, "%d", id)
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

// Отправка результата ping на электронную почту
func sendEmail(err error, pinger *ping.Pinger, portStatus bool) {
	m := gomail.NewMessage()

	smtpTo := os.Getenv("SMTPTo")
	smtpUser := os.Getenv("SMTPUser")
	smtpPass := os.Getenv("SMTPPass")

	m.SetHeader("From", smtpUser)
	m.SetHeader("To", smtpTo)
	m.SetHeader("Subject", "PING ERROR")

	// Тело письма
	body := fmt.Sprintf("Ошибка ping: %v\n\n Ping статистика:\n Пакетов отправлено = %d, получено = %d, потеряно = %.2f%%\n\n Доступ к серверу: %v",
		err,
		pinger.Statistics().PacketsSent,
		pinger.Statistics().PacketsRecv,
		pinger.Statistics().PacketLoss,
		portStatus,
	)
	m.SetBody("text/plain", body)

	// SMTP-сервер и порт
	d := gomail.NewDialer("smtp.mail.ru", 465, smtpUser, smtpPass)

	// Отправка письма
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Ошибка отправки письма:", err)
	} else {
		fmt.Println("Письмо отправлено")
	}
}
*/
