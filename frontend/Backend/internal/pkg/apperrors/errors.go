package apperrors

import (
	"fmt"
)

type ErrorType string

const (
	NotFound       ErrorType = "NOT_FOUND"
	AlreadyExists  ErrorType = "ALREADY_EXISTS"
	InvalidInput   ErrorType = "INVALID_INPUT"
	Unauthorized   ErrorType = "UNAUTHORIZED"
	InternalError  ErrorType = "INTERNAL_ERROR"
)

type AppError struct {
	Type    ErrorType
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s (%v)", e.Type, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Type, e.Message)
}

// Is implements error Is interface
func (e *AppError) Is(target error) bool {
	t, ok := target.(*AppError)
	if !ok {
		return false
	}
	return e.Type == t.Type
}

// エラー生成関数
func NewNotFoundError(message string, err error) *AppError {
	return &AppError{
		Type:    NotFound,
		Message: message,
		Err:     err,
	}
}

// 他のエラー関数も同様... 