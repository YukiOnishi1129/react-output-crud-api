package auth

import (
	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/domain"
	"github.com/google/uuid"
)

// Input DTOs
type RegisterInput struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// Output DTOs
type AuthOutput struct {
	Token string     `json:"token"`
	User  UserOutput `json:"user"`
}

type UserOutput struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

// ドメインモデルとDTOの変換関数
func ToUserOutput(user *domain.User) *UserOutput {
	return &UserOutput{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
} 