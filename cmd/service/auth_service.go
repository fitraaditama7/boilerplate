package service

import (
	"arka/cmd/entity"
	"context"
)

type AuthService interface {
	Login(ctx context.Context, data string, password string) (*entity.Authorization, error)
}
