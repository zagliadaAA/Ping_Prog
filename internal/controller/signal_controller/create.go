package signal_controller

import (
	"net/http"

	"ping_prog/internal/controller"
	"ping_prog/internal/usecase/signal_usecase"
)

type createSignalReq struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	// парсим json в нашу структуру
	// валидируем тело запроса или парамерты
	// вызываем юзкейс
	// обрабатываем ошибки если есть
	// возвращаем результат

	var req createSignalReq
	if err := controller.DecodeRequest(w, r, &req); err != nil {
		return
	}

	validationError := validateCreateSignalReq(&req)
	if validationError != nil {
		controller.RespondValidationError(w, validationError)

		return
	}

	err := c.signalUseCase.Create(signal_usecase.CreateSignalReq{
		Address: req.Address,
		Port:    req.Port,
	})
	if err != nil {
		controller.RespondStatusBadRequestError(w, controller.NewStatusBadRequestError("failed to create signal"))

		return
	}

	w.WriteHeader(http.StatusOK)
}

func validateCreateSignalReq(r *createSignalReq) *controller.ValidationError {
	if r.Address == "" || len(r.Address) > 15 {
		return controller.NewValidationError("address", "address not null, length no more then 15")
	}

	return nil
}
