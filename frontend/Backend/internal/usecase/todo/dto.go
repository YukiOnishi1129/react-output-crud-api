package todo

import (
	"time"

	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/domain"
	"github.com/google/uuid"
)

// Input DTOs
type CreateTodoInput struct {
	Title   string     `json:"title" validate:"required,min=1,max=100"`
	Content *string    `json:"content" validate:"omitempty,max=1000"`
	UserID  uuid.UUID `json:"user_id" validate:"required"`
}

// Output DTOs
type TodoOutput struct {
	ID        uuid.UUID  `json:"id"`
	Title     string     `json:"title"`
	Content   *string    `json:"content"`
	UserID    uuid.UUID  `json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// 変換関数
func ToTodoOutput(todo *domain.Todo) *TodoOutput {
	return &TodoOutput{
		ID:        todo.ID,
		Title:     todo.Title,
		Content:   todo.Content,
		UserID:    todo.UserID,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}
} 