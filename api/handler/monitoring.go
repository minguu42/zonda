package handler

import (
	"context"

	"github.com/minguu42/zonda/lib/go/zondaapi"
)

func (h *handler) CheckHealth(_ context.Context) (*zondaapi.CheckHealthOK, error) {
	out := h.monitoring.CheckHealth()
	return &zondaapi.CheckHealthOK{Revision: out.Revision}, nil
}
