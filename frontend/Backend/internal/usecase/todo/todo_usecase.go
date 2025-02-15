package usecase

import (
	"context"

	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/domain"
	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/repository"
	"github.com/google/uuid"
)

type TodoUseCase interface {
	GetTodos(ctx context.Context) (*TodoListOutput, error)
	GetTodo(ctx context.Context, id uuid.UUID) (*TodoWithUserOutput, error)
	CreateTodo(ctx context.Context, input *CreateTodoInput) (*domain.Todo, error)
	UpdateTodo(ctx context.Context, input *UpdateTodoInput) (*domain.Todo, error)
	DeleteTodo(ctx context.Context, id uuid.UUID) error
}

type todoUseCase struct {
	todoRepo repository.TodoRepository
}

func NewTodoUseCase(todoRepo repository.TodoRepository) TodoUseCase {
	return &todoUseCase{todoRepo: todoRepo}
}

// 入力DTO
type CreateTodoInput struct {
	Title   string  `json:"title"`
	Content *string `json:"content"`
}

type UpdateTodoInput struct {
	ID      uuid.UUID `json:"id"`
	Title   string    `json:"title"`
	Content *string   `json:"content"`
}

// 実装
func (u *todoUseCase) GetTodos(ctx context.Context) (*TodoListOutput, error) {
	todos, err := u.todoRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	
	total, err := u.todoRepo.Count(ctx)
	if err != nil {
		return nil, err
	}
	
	return NewTodoListOutput(todos, total), nil
}

func (u *todoUseCase) GetTodo(ctx context.Context, id uuid.UUID) (*TodoWithUserOutput, error) {
	todo, err := u.todoRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	
	return NewTodoWithUserOutput(todo), nil
}

func (u *todoUseCase) CreateTodo(ctx context.Context, input *CreateTodoInput) (*domain.Todo, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	todo := &domain.Todo{
		Title:   input.Title,
		Content: input.Content,
		UserID:  input.UserID,
	}

	if err := u.todoRepo.Create(ctx, todo); err != nil {
		return nil, err
	}

	return todo, nil
}

// 他のメソッドも同様に実装... 