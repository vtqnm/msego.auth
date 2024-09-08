package services

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/vtqnm/msego.auth/internal/lib/logger/sl"
	"github.com/vtqnm/msego.auth/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Save(
		ctx context.Context,
		email string,
		passHash []byte,
	)
}

func (usr *UserRepository) FindByEmail(
	ctx context.Context,
	email string,
) (*models.User, error) {

}

type Auth interface {
	Login(
		ctx context.Context,
		email string,
		password string,
	) (token string, err error)
	Register(
		ctx context.Context,
		email string,
		password string,
	)
}

type AuthService struct {
	Auth

	userRep *UserRepository
	log *slog.Logger
}

func (s *AuthService) Login(
	ctx context.Context,
	email string,
	password string,
) (token string, err error) {
	const op = "AuthService.Login"

	log := s.log.With(
		slog.String("op", op),
		slog.String("email", email),
	)

	log.Info("attempting to login user")

	user, err := s.userRep.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			s.log.Warn("user not found", sl.Err(err))
			return "", fmt.Errorf("%s %w", op, ErrInvalidCredentials)
		}

		s.log.Error("failed to get user", sl.Err(err))
		return "", fmt.Errorf("%s %w", op, err)
	}
}

func (s *AuthService) Register(
	ctx context.Context,
	email string,
	password string,
) (uid int64, err error) {
	const op = "AuthService.Register"

	log := s.log.With(
		slog.String("op", op),
		slog.String("email", email),
	)

	log.Info("register user")

	passHash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("failed to generate password hash", sl.Err(err))
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err, := a.userRep.Save(ctx, email, passHash)
}
