package auth

import (
	"context"
	"fmt"

	"github.com/muasx88/library-api/internal/config"
	"github.com/muasx88/library-api/internal/response"
)

type IService interface {
	Register(context.Context, RegisterRequestPayload) error
	Login(context.Context, LoginRequestPayload) (string, error)
}

type service struct {
	repo IRepository
}

func newService(repo IRepository) IService {
	return &service{
		repo: repo,
	}
}

func (s service) Register(ctx context.Context, req RegisterRequestPayload) (err error) {
	userEntity := NewFromRegisterRequest(req)

	salt := int(config.Config.App.Encryption.Salt) // salt from config
	err = userEntity.EncryptPassword(salt)
	if err != nil {
		return fmt.Errorf("error encrypt password. %s", err.Error())
	}

	model, err := s.repo.GetAuthByEmail(ctx, userEntity.Email)
	if err != nil {
		if err != response.ErrUserNotFound {
			return
		}
	}

	if model.IsExists() {
		return response.ErrUserAlreadyExists
	}

	return s.repo.CreateAuth(ctx, userEntity)
}

func (s service) Login(ctx context.Context, req LoginRequestPayload) (token string, err error) {
	userEntity := NewFromLoginRequest(req)

	model, err := s.repo.GetAuthByEmail(ctx, userEntity.Email)
	if err != nil {
		if err == response.ErrUserNotFound {
			err = response.ErrUserPasswordNotMatch
			return
		}

		return
	}

	if err = userEntity.VerifyPasswordFromPlain(model.Password); err != nil {
		err = response.ErrUserPasswordNotMatch
		return
	}

	token, err = model.GenerateToken(config.Config.App.Encryption.JWTSecret)
	return
}
