package service

import (
	"arka/cmd/entity"
	"context"
)

type UserService interface {
	ListUser(ctx context.Context, page, limit int64) ([]*entity.User, error)
	GetUserByID(ctx context.Context, id string) (*entity.User, error)
	InsertUser(ctx context.Context, user *entity.User) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User, id string) (*entity.User, error)
	DeleteUser(ctx context.Context, id string, userID string) error
}
