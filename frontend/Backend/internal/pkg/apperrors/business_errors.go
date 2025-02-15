package apperrors

// ビジネスロジック関連のエラー型
const (
	ValidationError   ErrorType = "VALIDATION_ERROR"
	BusinessRuleError ErrorType = "BUSINESS_RULE_ERROR"
	PermissionDenied  ErrorType = "PERMISSION_DENIED"
)

// ビジネスロジック関連のエラー生成関数
func NewValidationError(message string, err error) *AppError {
	return &AppError{
		Type:    ValidationError,
		Message: message,
		Err:     err,
	}
}

func NewBusinessRuleError(message string, err error) *AppError {
	return &AppError{
		Type:    BusinessRuleError,
		Message: message,
		Err:     err,
	}
}

func NewPermissionDeniedError(message string, err error) *AppError {
	return &AppError{
		Type:    PermissionDenied,
		Message: message,
		Err:     err,
	}
} 