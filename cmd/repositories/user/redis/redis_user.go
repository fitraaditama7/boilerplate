package redis

import (
	"arka/cmd/entity"
	"arka/cmd/repositories"
	"arka/pkg/cache"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/sirupsen/logrus"
)

type Redis struct {
	conn   cache.RedisCommand
	repo   repositories.UserRepository
	prefix string
	module string
}

func NewUser(conn cache.RedisCommand, repo repositories.UserRepository, prefix string) *Redis {
	return &Redis{conn, repo, prefix, "user"}
}

func (r *Redis) ListUser(ctx context.Context, page, limit int64) ([]*entity.User, error) {
	return r.repo.ListUser(ctx, page, limit)
}

func (r *Redis) GetUserByEmailOrPhone(ctx context.Context, data string) (*entity.User, error) {
	var key = fmt.Sprintf("%s:%s:%s", r.prefix, r.module, data)
	var user = new(entity.User)

	value, err := r.conn.Get(key)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	if value == nil {
		user, err = r.repo.GetUserByEmailOrPhone(ctx, data)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		data, err := json.Marshal(user)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		err = r.conn.SetEx(key, cache.ONEHOUR, string(data))
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return user, err
	}

	err = json.Unmarshal(value, user)
	if err != nil {
		logrus.Println(err)
		return nil, err
	}
	return user, err
}
