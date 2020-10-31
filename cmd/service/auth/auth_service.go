package service

import (
	"arka/cmd/entity"
	"arka/cmd/lib/authentication"
	"arka/cmd/lib/customError"
	"arka/cmd/repositories"
	"arka/pkg/auth"
	"context"
	"database/sql"
	"log"

	"github.com/sirupsen/logrus"
)

type authService struct {
	userRepository repositories.UserRepository
	authorization  auth.Auth
	token          auth.Token
}

func NewAuthService(userRepository repositories.UserRepository, authorization auth.Auth, token auth.Token) *authService {
	return &authService{
		userRepository: userRepository,
		authorization:  authorization,
		token:          token,
	}
}

func (s *authService) Login(ctx context.Context, data string, password string) (*entity.Authorization, error) {
	user, err := s.userRepository.GetUserByEmailOrPhone(ctx, data)
	if err == sql.ErrNoRows {
		logrus.Error(err)
		return nil, customError.ErrInvalidLogin
	}
	if err != nil {
		log.Println(err)
		return nil, customError.ErrInternalServerError
	}

	ok := authentication.ComparePassword(password, user.Password)
	if ok == false {
		logrus.Error("password doesn't match")
		return nil, customError.ErrInvalidLogin
	}

	tokenDetail, err := s.token.CreateToken(user.ID, user.RoleID)
	if err != nil {
		log.Println(err)
		return nil, customError.ErrUnProcessableEntity
	}

	err = s.authorization.CreateAuth(user.ID, tokenDetail)
	if err != nil {
		log.Println(err)
		return nil, customError.ErrUnProcessableEntity
	}

	result := &entity.Authorization{
		User:    *user,
		Token:   tokenDetail.AccessToken,
		Refresh: tokenDetail.RefreshToken,
	}
	return result, err
}
