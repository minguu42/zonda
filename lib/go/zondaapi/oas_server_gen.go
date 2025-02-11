// Code generated by ogen, DO NOT EDIT.

package zondaapi

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// CheckHealth implements checkHealth operation.
	//
	// GET /health
	CheckHealth(ctx context.Context) (*CheckHealthOK, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}
