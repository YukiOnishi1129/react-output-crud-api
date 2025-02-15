package authdto

import (
	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/domain"
	"github.com/google/uuid"
)

type AuthOutput struct {
	Token string     `json:"token"`
	User  UserOutput `json:"user"`
}

type UserOutput struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

// FromDomain ドメインモデルからDTOへの変換
func FromDomain(user *domain.User) *UserOutput {
	return &UserOutput{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

// ToDomain DTOからドメインモデルへの変換
func (i *RegisterInput) ToDomain() *domain.User {
	return &domain.User{
		Name:     i.Name,
		Email:    i.Email,
		Password: i.Password,
	}
} 