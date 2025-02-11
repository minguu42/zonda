package handler

import (
	"net/http"

	"github.com/minguu42/zonda/api/usecase"
	"github.com/minguu42/zonda/lib/go/zondaapi"
)

type handler struct {
	monitoring usecase.Monitoring
}

func New() (http.Handler, error) {
	return zondaapi.NewServer(&handler{monitoring: usecase.Monitoring{}})
}
