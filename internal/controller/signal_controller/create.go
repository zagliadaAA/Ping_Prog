package signal_controller

import (
	"fmt"
	"net/http"
)

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	// парсим json в нашу структуру
	// валидируем тело запроса или парамерты
	// вызываем юзкейс
	// обрабатываем ошибки если есть
	// возвращаем результат
	//
	fmt.Println("Hello")
}
