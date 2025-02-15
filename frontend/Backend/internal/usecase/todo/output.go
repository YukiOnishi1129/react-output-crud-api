package usecase

import (
	"time"

	"github.com/google/uuid"
)

// TodoOutput は、Todoの表示用DTOです
type TodoOutput struct {
	ID        uuid.UUID  `json:"id"`
	Title     string     `json:"title"`
	Content   *string    `json:"content"`
	UserID    uuid.UUID  `json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// TodoWithUserOutput は、ユーザー情報を含むTodoの表示用DTOです
type TodoWithUserOutput struct {
	ID        uuid.UUID      `json:"id"`
	Title     string         `json:"title"`
	Content   *string        `json:"content"`
	User      UserMinOutput  `json:"user"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

// UserMinOutput は、Todo表示時の最小限のユーザー情報DTOです
type UserMinOutput struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

// TodoListOutput は、Todoリストの表示用DTOです
type TodoListOutput struct {
	Todos []TodoOutput `json:"todos"`
	Total int64        `json:"total"`
}

// NewTodoOutput は、domainモデルからDTOを生成します
func NewTodoOutput(todo *domain.Todo) *TodoOutput {
	return &TodoOutput{
		ID:        todo.ID,
		Title:     todo.Title,
		Content:   todo.Content,
		UserID:    todo.UserID,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}
}

// NewTodoWithUserOutput は、ユーザー情報を含むDTOを生成します
func NewTodoWithUserOutput(todo *domain.Todo) *TodoWithUserOutput {
	if todo.User == nil {
		return nil
	}

	return &TodoWithUserOutput{
		ID:      todo.ID,
		Title:   todo.Title,
		Content: todo.Content,
		User: UserMinOutput{
			ID:    todo.User.ID,
			Name:  todo.User.Name,
			Email: todo.User.Email,
		},
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}
}

// NewTodoListOutput は、Todoリストのレスポンスを生成します
func NewTodoListOutput(todos []*domain.Todo, total int64) *TodoListOutput {
	outputs := make([]TodoOutput, len(todos))
	for i, todo := range todos {
		outputs[i] = *NewTodoOutput(todo)
	}
	return &TodoListOutput{
		Todos: outputs,
		Total: total,
	}
} 