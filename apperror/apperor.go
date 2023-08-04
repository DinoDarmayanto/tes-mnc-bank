package apperror

import "fmt"

type AppError struct {
	ErrorMessage string
	ErrorCode    int
}

func (ae AppError) Error() string {
	return fmt.Sprintf("%v - %v", ae.ErrorCode, ae.ErrorMessage)
}

func NewAppError(errorCode int, errorMessage string) error {
	return AppError{
		ErrorCode:    errorCode,
		ErrorMessage: errorMessage,
	}
}
