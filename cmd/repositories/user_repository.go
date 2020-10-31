package repositories

import (
	"arka/cmd/entity"
	"context"
)

type UserRepository interface {
	ListUser(ctx context.Context, page, limit int64) ([]*entity.User, error)
	GetUserByEmailOrPhone(ctx context.Context, data string) (*entity.User, error)
}
