package controller

import (
	"encoding/json"
	"net/http"
)

// RespondJSON отправляет JSON-ответ клиенту с указанным кодом статуса.
func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if payload != nil {
		json.NewEncoder(w).Encode(payload)
	}
}

// RespondValidationError отправляет JSON-ответ с ошибкой.
func RespondValidationError(w http.ResponseWriter, v *ValidationError) {
	RespondJSON(w, http.StatusBadRequest, v)
}

func RespondStatusBadRequestError(w http.ResponseWriter, s *StatusBadRequestError) {
	RespondJSON(w, http.StatusBadRequest, s)
}

// DecodeRequest считывает тело запроса и помещает его в структуру
func DecodeRequest(w http.ResponseWriter, r *http.Request, req interface{}) error {
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		RespondValidationError(w, NewValidationError("decode", "error reading json"))

		return err
	}

	return nil
}
