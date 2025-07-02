package custom_error

import(
	"net/http"
)

type CustomError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *CustomError) Error() string {
	return e.Message
}

// 400 Bad Request
func NewBadRequestError(message string) *CustomError {
	return &CustomError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

// 409 Conflict
func NewConflictError(message string) *CustomError {
	return &CustomError{
		Code:    http.StatusConflict,
		Message: message,
	}
}

// 500 Internal Server Error
func NewInternalServerError(message string) *CustomError {
	return &CustomError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}