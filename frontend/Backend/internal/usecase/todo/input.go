package usecase

import (
	"errors"

	"github.com/google/uuid"
)

// CreateTodoInput は、Todoの作成用DTOです
type CreateTodoInput struct {
	Title     string  `json:"title" validate:"required,min=1,max=100"`
	Content   *string `json:"content" validate:"omitempty,max=1000"`
	UserID    uuid.UUID `json:"user_id" validate:"required"`
}

// UpdateTodoInput は、Todoの更新用DTOです
type UpdateTodoInput struct {
	ID        uuid.UUID `json:"id" validate:"required"`
	Title     string    `json:"title" validate:"required,min=1,max=100"`
	Content   *string   `json:"content" validate:"omitempty,max=1000"`
}

// DeleteTodoInput は、Todoの削除用DTOです
type DeleteTodoInput struct {
	ID     uuid.UUID `json:"id" validate:"required"`
	UserID uuid.UUID `json:"user_id" validate:"required"`
}

// Validate は、CreateTodoInputのバリデーションを行います
func (i *CreateTodoInput) Validate() error {
	if i.Title == "" {
		return errors.New("title is required")
	}
	if len(i.Title) > 100 {
		return errors.New("title must be less than 100 characters")
	}
	if i.Content != nil && len(*i.Content) > 1000 {
		return errors.New("content must be less than 1000 characters")
	}
	if i.UserID == uuid.Nil {
		return errors.New("user_id is required")
	}
	return nil
}

// Validate は、UpdateTodoInputのバリデーションを行います
func (i *UpdateTodoInput) Validate() error {
	if i.ID == uuid.Nil {
		return errors.New("id is required")
	}
	if i.Title == "" {
		return errors.New("title is required")
	}
	if len(i.Title) > 100 {
		return errors.New("title must be less than 100 characters")
	}
	if i.Content != nil && len(*i.Content) > 1000 {
		return errors.New("content must be less than 1000 characters")
	}
	return nil
}

// Validate は、DeleteTodoInputのバリデーションを行います
func (i *DeleteTodoInput) Validate() error {
	if i.ID == uuid.Nil {
		return errors.New("id is required")
	}
	if i.UserID == uuid.Nil {
		return errors.New("user_id is required")
	}
	return nil
}

// GetTodoInput は、Todo取得用DTOです
type GetTodoInput struct {
	ID     uuid.UUID `json:"id" validate:"required"`
	UserID uuid.UUID `json:"user_id" validate:"required"`
}

// GetTodosInput は、Todoリスト取得用DTOです
type GetTodosInput struct {
	UserID uuid.UUID `json:"user_id" validate:"required"`
	Limit  int       `json:"limit" validate:"required,min=1,max=100"`
	Offset int       `json:"offset" validate:"min=0"`
}

// Validate は、GetTodoInputのバリデーションを行います
func (i *GetTodoInput) Validate() error {
	if i.ID == uuid.Nil {
		return errors.New("id is required")
	}
	if i.UserID == uuid.Nil {
		return errors.New("user_id is required")
	}
	return nil
}

// Validate は、GetTodosInputのバリデーションを行います
func (i *GetTodosInput) Validate() error {
	if i.UserID == uuid.Nil {
		return errors.New("user_id is required")
	}
	if i.Limit < 1 || i.Limit > 100 {
		return errors.New("limit must be between 1 and 100")
	}
	if i.Offset < 0 {
		return errors.New("offset must be greater than or equal to 0")
	}
	return nil
} 