package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/minguu42/zonda/api/apperr"
	"github.com/minguu42/zonda/api/factory"
	"github.com/minguu42/zonda/api/usecase"
	"github.com/minguu42/zonda/lib/go/zondaapi"
)

type handler struct {
	authentication usecase.Authentication
	monitoring     usecase.Monitoring
}

func New(f *factory.Factory) (http.Handler, error) {
	return zondaapi.NewServer(&handler{
		authentication: usecase.NewAuthentication(f.Auth, f.DB),
		monitoring:     usecase.Monitoring{},
	})
}

func (h *handler) NewError(_ context.Context, err error) *zondaapi.ErrorStatusCode {
	var appErr apperr.Error
	switch {
	case errors.As(err, &appErr):
	case errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded):
		appErr = apperr.ErrDeadlineExceeded(err)
	default:
		appErr = apperr.ErrUnknown(err)
	}

	return appErr.APIError()
}
