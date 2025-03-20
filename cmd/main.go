package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/go-ping/ping"
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func main() {
	err := godotenv.Load()
	if err != nil { //письмо не отправится
		log.Fatal("Error loading .env file")
	}

	addr := os.Getenv("Address")
	port, err := strconv.Atoi(os.Getenv("Port"))

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
	}
}

// команда ping
func commandPing(addr string) (error, *ping.Pinger) {
	pinger, err := ping.NewPinger(addr)
	if err != nil {
		return fmt.Errorf("ошибка создания pinger"), pinger
	}

	pinger.Count = 5
	pinger.Timeout = 1 * time.Second
	pinger.SetPrivileged(true)

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
