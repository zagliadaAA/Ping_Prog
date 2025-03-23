package service_provider

import (
	"net/http"
)

func (sp *ServiceProvider) GetRoutes() *http.ServeMux {
	// роутер наших ручек (пути для ручек(хэндлеров))
	mux := http.NewServeMux()
	// signal
	mux.HandleFunc("POST /signals", sp.getSignalController().Create)
	mux.HandleFunc("/signals", sp.getSignalController().Create)

	/*mux.HandleFunc("DELETE /clients/{id}", sp.getClientController().Delete)
	mux.HandleFunc("PUT /clients/{id}", sp.getClientController().Update)
	mux.HandleFunc("GET /clients/{id}", sp.getClientController().GetClientByID)*/

	return mux
}
