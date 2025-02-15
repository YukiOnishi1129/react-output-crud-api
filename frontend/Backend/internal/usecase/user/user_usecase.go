package usecase

import (
	"context"

	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/domain"
	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	GetUsers(ctx context.Context) ([]*domain.User, error)
	GetUser(ctx context.Context, id uuid.UUID) (*domain.User, error)
	CreateUser(ctx context.Context, input *CreateUserInput) (*domain.User, error)
	UpdateUser(ctx context.Context, input *UpdateUserInput) (*domain.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type userUseCase struct {
	userRepo repository.UserRepository
	todoRepo repository.TodoRepository
}

func NewUserUseCase(userRepo repository.UserRepository, todoRepo repository.TodoRepository) UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
		todoRepo: todoRepo,
	}
}

func (u *userUseCase) CreateUser(ctx context.Context, input *CreateUserInput) (*domain.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	if err := u.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

// 他のメソッドも同様に実装... 