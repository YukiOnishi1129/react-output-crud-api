package authdto

import "github.com/google/uuid"

type AuthOutput struct {
	Token string     `json:"token"`
	User  UserOutput `json:"user"`
}

type UserOutput struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

// ToUserOutput ドメインモデルからDTOへの変換
func ToUserOutput(user *domain.User) *UserOutput {
	return &UserOutput{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
} 