package controller

type ValidationError struct {
	Param   string `json:"param"`
	Message string `json:"message"`
}

func NewValidationError(param string, message string) *ValidationError {
	return &ValidationError{
		Param:   param,
		Message: message,
	}

}

type StatusBadRequestError struct {
	Message string `json:"message"`
}

func NewStatusBadRequestError(message string) *StatusBadRequestError {
	return &StatusBadRequestError{
		Message: message,
	}
}
