package factory

import (
	"fmt"

	"github.com/minguu42/zonda/api/config"
	"github.com/minguu42/zonda/api/database"
	"github.com/minguu42/zonda/api/jwtauth"
)

type Factory struct {
	Auth *jwtauth.Authenticator
	DB   *database.Client
}

func New(conf *config.Config) (*Factory, error) {
	auth := &jwtauth.Authenticator{
		AccessTokenExpiry:  conf.Auth.AccessTokenExpiry,
		RefreshTokenExpiry: conf.Auth.RefreshTokenExpiry,
		AccessTokenSecret:  conf.Auth.AccessTokenSecret,
		RefreshTokenSecret: conf.Auth.RefreshTokenSecret,
	}

	db, err := database.NewClient(conf.DB)
	if err != nil {
		return nil, fmt.Errorf("failed to create database client: %w", err)
	}

	return &Factory{
		Auth: auth,
		DB:   db,
	}, nil
}

func (f *Factory) Close() {
	_ = f.DB.Close()
}
