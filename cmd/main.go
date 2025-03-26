package main

import (
	"log"

	"ping_prog/cmd/service_provider"
)

func main() {
	sp := service_provider.NewServiceProvider()

	go sp.GetTelegramBot().Start()
	log.Println("бот запущен...")

	select {} // это "пустой выбор" без каналов, чтобы программа не завершалась
}
