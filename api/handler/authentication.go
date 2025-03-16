package handler

import (
	"context"
	"fmt"

	"github.com/minguu42/zonda/api/usecase"
	"github.com/minguu42/zonda/lib/go/zondaapi"
)

func (h *handler) SignUp(ctx context.Context, req *zondaapi.SignUpReq) (*zondaapi.SignUpOK, error) {
	out, err := h.authentication.SignUp(ctx, &usecase.SignUpInput{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to execute SignUp usecase: %w", err)
	}
	return &zondaapi.SignUpOK{
		AccessToken:  out.AccessToken,
		RefreshToken: out.RefreshToken,
	}, nil
}

func (h *handler) SignIn(ctx context.Context, req *zondaapi.SignInReq) (*zondaapi.SignInOK, error) {
	out, err := h.authentication.SignIn(ctx, &usecase.SignInInput{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to execute SignIn usecase: %w", err)
	}
	return &zondaapi.SignInOK{
		AccessToken:  out.AccessToken,
		RefreshToken: out.RefreshToken,
	}, nil
}

func (h *handler) RefreshToken(ctx context.Context, req *zondaapi.RefreshTokenReq) (*zondaapi.RefreshTokenOK, error) {
	out, err := h.authentication.RefreshToken(ctx, &usecase.RefreshTokenInput{
		RefreshToken: req.RefreshToken,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to execute RefreshToken usecase: %w", err)
	}
	return &zondaapi.RefreshTokenOK{
		AccessToken:  out.AccessToken,
		RefreshToken: out.RefreshToken,
	}, nil
}
