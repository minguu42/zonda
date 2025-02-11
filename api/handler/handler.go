package handler

import (
	"net/http"

	"github.com/minguu42/zonda/api/oapi"
	"github.com/minguu42/zonda/api/usecase"
)

type handler struct {
	monitoring usecase.Monitoring
}

func New() (http.Handler, error) {
	return oapi.NewServer(&handler{monitoring: usecase.Monitoring{}})
}
