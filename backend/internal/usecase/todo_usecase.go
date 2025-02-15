package usecase

import (
	"context"

	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/domain"
	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/repository"
	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/usecase/input"
	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/usecase/output"
)

type TodoUseCase interface {
	ListTodo(ctx context.Context) (*output.TodoListOutput, error)
	GetTodo(ctx context.Context, input *input.GetTodoInput) (*output.TodoOutput, error)
	CreateTodo(ctx context.Context, input *input.CreateTodoInput) (*output.TodoOutput, error)
	UpdateTodo(ctx context.Context, input *input.UpdateTodoInput) (*output.TodoOutput, error)
	DeleteTodo(ctx context.Context, input *input.DeleteTodoInput) error
}

type todoUseCase struct {
	todoRepo repository.TodoRepository
}

func NewTodoUseCase(todoRepo repository.TodoRepository) TodoUseCase {
	return &todoUseCase{todoRepo: todoRepo}
}

func (u *todoUseCase)ListTodo(ctx context.Context) (*output.TodoListOutput, error) {
	todos, err := u.todoRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return output.NewTodoListOutput(todos, int64(len(todos))), nil
}

func (u *todoUseCase)GetTodo(ctx context.Context, input *input.GetTodoInput) (*output.TodoOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	todo, err := u.todoRepo.FindByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	return output.NewTodoOutput(todo), nil
}

func (u *todoUseCase)CreateTodo(ctx context.Context, input *input.CreateTodoInput) (*output.TodoOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	todo := domain.Todo{
		Title:   input.Title,
		Content: input.Content,
	}

	if err := u.todoRepo.Create(ctx, &todo); err != nil {
		return nil, err
	}

	return output.NewTodoOutput(&todo), nil
}

func (u *todoUseCase)UpdateTodo(ctx context.Context, input *input.UpdateTodoInput) (*output.TodoOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	todo, err := u.todoRepo.FindByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	todo.Title = input.Title
	todo.Content = input.Content

	if err := u.todoRepo.Update(ctx, todo); err != nil {
		return nil, err
	}

	return output.NewTodoOutput(todo), nil
}

func (u *todoUseCase)DeleteTodo(ctx context.Context, input *input.DeleteTodoInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return u.todoRepo.Delete(ctx, input.ID)
}
