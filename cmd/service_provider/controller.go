package service_provider

import (
	"ping_prog/internal/controller/signal_controller"
)

func (sp *ServiceProvider) getSignalController() *signal_controller.Controller {
	if sp.signalController == nil {
		sp.signalController = signal_controller.NewController(sp.GetSignalUseCase())
	}

	return sp.signalController
}
