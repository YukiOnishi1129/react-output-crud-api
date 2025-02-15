package output

import (
	"time"

	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/domain"
	"github.com/google/uuid"
)

type TodoOutput struct {
	ID        uuid.UUID  `json:"id"`
	Title     string     `json:"title"`
	Content   *string    `json:"content"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type TodoListOutput struct {
	Todos []TodoOutput `json:"todos"`
	Total int64        `json:"total"`
}

func NewTodoOutput(todo *domain.Todo) *TodoOutput {
	return &TodoOutput{
		ID:        todo.ID,
		Title:     todo.Title,
		Content:   todo.Content,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}
}

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