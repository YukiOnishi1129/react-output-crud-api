package auth

import (
	"context"
)

type AuthUseCase interface {
	Register(ctx context.Context, input *authdto.RegisterInput) (*authdto.AuthOutput, error)
	Login(ctx context.Context, input *authdto.LoginInput) (*authdto.AuthOutput, error)
	// ...
} 