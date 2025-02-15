package gorm

import (
	"context"

	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/domain"
	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) repository.TodoRepository {
	return &todoRepository{db: db}
}

func (r *todoRepository) FindAll(ctx context.Context) ([]*domain.Todo, error) {
	var todos []*domain.Todo
	if err := r.db.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *todoRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Todo, error) {
	var todo domain.Todo
	if err := r.db.First(&todo, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *todoRepository) Create(ctx context.Context, todo *domain.Todo) error {
	return r.db.Create(todo).Error
}

func (r *todoRepository) Update(ctx context.Context, todo *domain.Todo) error {
	return r.db.Save(todo).Error
}

func (r *todoRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.Delete(&domain.Todo{}, "id = ?", id).Error
} 