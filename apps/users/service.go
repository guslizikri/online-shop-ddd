package users

import (
	"context"
	"database/sql"
	"online-shop-ddd/infra/response"
	"online-shop-ddd/internal/config"
)

type Repository interface {
	GetUserByEmail(ctx context.Context, email string) (model UserEntity, err error)
	CreateUser(ctx context.Context, model UserEntity) (err error)
}

// ini tidak memakai capital, karena tidak perlu di export
type service struct {
	repo Repository
}

func newService(repo Repository) service {
	return service{
		repo: repo,
	}
}

func (s service) register(ctx context.Context, req RegisterRequestPayload) (err error) {
	authEntity := NewFromRegisterRequest(req)

	if err = authEntity.Validate(); err != nil {
		return
	}

	if err = authEntity.EncryptPassword(int(config.Cfg.App.Encryption.Salt)); err != nil {
		return
	}

	// validasi apakah email sudah terdaftar/belum
	model, err := s.repo.GetUserByEmail(ctx, authEntity.Email)
	if err != nil {
		// jika terdapat error not found maka lanjut melakukan create user
		// jika terjadi error selain error not found maka akan me return error tersebut
		if err != response.ErrNotFound {
			return err
		}
	}

	if model.IsExists() {
		return response.ErrEmailExist
	}

	err = s.repo.CreateUser(ctx, authEntity)
	return err
}

func (s service) login(ctx context.Context, req LoginRequestPayload) (token string, err error) {
	authEntity := NewFromLoginRequest(req)

	if err = authEntity.ValidateEmail(); err != nil {
		return
	}
	if err = authEntity.ValidatePassword(); err != nil {
		return
	}
	// model ini berisi data dari database
	model, err := s.repo.GetUserByEmail(ctx, authEntity.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			err = response.ErrNotFound
			return
		}
		return
	}

	err = authEntity.VerifyPasswordFromPlain(model.Password)
	if err != nil {
		err = response.ErrPasswordNotMatch
		return
	}

	token, err = model.GenerateToken(config.Cfg.App.Encryption.JwtSecret)

	return
}
