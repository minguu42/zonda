package handler

import (
	"context"
	"fmt"

	"github.com/minguu42/zonda/api/usecase"
	"github.com/minguu42/zonda/lib/go/zondaapi"
)

func (h *handler) SignUp(ctx context.Context, req *zondaapi.SignUpReq) (zondaapi.SignUpRes, error) {
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
