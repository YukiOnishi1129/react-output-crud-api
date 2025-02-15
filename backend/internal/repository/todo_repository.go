package repository

import (
	"context"

	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/domain"
	"github.com/google/uuid"
)

type TodoRepository interface {
	FindAll(ctx context.Context) ([]*domain.Todo, error)
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Todo, error)
	Create(ctx context.Context, todo *domain.Todo) error
	Update(ctx context.Context, todo *domain.Todo) error
	Delete(ctx context.Context, id uuid.UUID) error
}