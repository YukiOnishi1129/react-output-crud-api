package repository

import (
	"context"

	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/domain"
	"github.com/google/uuid"
)

type UserRepository interface {
	FindAll(ctx context.Context) ([]*domain.User, error)
	FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	Create(ctx context.Context, user *domain.User) error
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id uuid.UUID) error
} 