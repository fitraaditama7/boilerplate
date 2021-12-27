package service

import (
	"arka/cmd/entity"
	"arka/pkg/auth"
	"context"
)

type AuthService interface {
	Login(ctx context.Context, data string, password string) (*entity.Authorization, error)
	Logout(ctx context.Context, accessDetail *auth.AccessDetails) error
}
