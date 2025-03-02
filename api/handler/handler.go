package handler

import (
	"net/http"

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
