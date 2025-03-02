package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/minguu42/zonda/api/apperr"
	"github.com/minguu42/zonda/api/database"
	"github.com/minguu42/zonda/api/domain"
	"github.com/minguu42/zonda/api/jwtauth"
	"github.com/minguu42/zonda/lib/go/idgen"
	"golang.org/x/crypto/bcrypt"
)

type Authentication struct {
	auth *jwtauth.Authenticator
	db   *database.Client
}

func NewAuthentication(auth *jwtauth.Authenticator, db *database.Client) Authentication {
	return Authentication{
		auth: auth,
		db:   db,
	}
}

type SignUpInput struct {
	Email    string
	Password string
}

func (in *SignUpInput) User() (*domain.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to generate hashed password: %w", err)
	}
	return &domain.User{
		ID:             domain.UserID(idgen.ULID()),
		Email:          in.Email,
		HashedPassword: string(hashedPassword),
	}, nil
}

type SignUpOutput struct {
	AccessToken  string
	RefreshToken string
}

func (uc *Authentication) SignUp(ctx context.Context, in *SignUpInput) (*SignUpOutput, error) {
	if _, err := uc.db.GetUserByEmail(ctx, in.Email); !errors.Is(err, database.ErrModelNotFound) {
		return nil, apperr.ErrDuplicateUserEmail(err)
	}

	user, err := in.User()
	if err != nil {
		return nil, fmt.Errorf("failed to generate user: %w", err)
	}

	accessToken, err := uc.auth.CreateAccessToken(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to create access token: %w", err)
	}
	refreshToken, err := uc.auth.CreateRefreshToken(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to create refresh token: %w", err)
	}

	if err := uc.db.CreateUser(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return &SignUpOutput{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
