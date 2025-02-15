package auth

import (
	"context"

	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/domain"
	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase interface {
	Register(ctx context.Context, input *RegisterInput) (*AuthOutput, error)
	Login(ctx context.Context, input *LoginInput) (*AuthOutput, error)
	RefreshToken(ctx context.Context, token string) (*AuthOutput, error)
	Logout(ctx context.Context) error
}

type authUseCase struct {
	userRepo repository.UserRepository
	jwtService JWTService
}

func NewAuthUseCase(userRepo repository.UserRepository, jwtService JWTService) AuthUseCase {
	return &authUseCase{
		userRepo: userRepo,
		jwtService: jwtService,
	}
}

func (u *authUseCase) Register(ctx context.Context, input *RegisterInput) (*AuthOutput, error) {
	// パスワードのハッシュ化
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

	// JWTトークンの生成
	token, err := u.jwtService.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &AuthOutput{
		Token: token,
		User: UserOutput{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
} 