package handler

import (
	"context"

	"github.com/minguu42/zonda/api/oapi"
)

func (h *handler) CheckHealth(_ context.Context) (*oapi.CheckHealthOK, error) {
	out := h.monitoring.CheckHealth()
	return &oapi.CheckHealthOK{Revision: out.Revision}, nil
}
