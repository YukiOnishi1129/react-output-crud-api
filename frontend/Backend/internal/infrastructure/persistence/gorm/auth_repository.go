package gorm

import (
	"context"

	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/domain"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func (r *authRepository) Create(ctx context.Context, input *authdto.RegisterInput) (*domain.User, error) {
	user := input.ToDomain()
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *authRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
} 