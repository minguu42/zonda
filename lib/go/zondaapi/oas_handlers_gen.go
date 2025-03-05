// Code generated by ogen, DO NOT EDIT.

package zondaapi

import (
	"context"
	"net/http"

	"github.com/go-faster/errors"

	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
)

type codeRecorder struct {
	http.ResponseWriter
	status int
}

func (c *codeRecorder) WriteHeader(status int) {
	c.status = status
	c.ResponseWriter.WriteHeader(status)
}

func recordError(string, error) {}

// handleCheckHealthRequest handles checkHealth operation.
//
// GET /health
func (s *Server) handleCheckHealthRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	ctx := r.Context()

	var (
		err error
	)

	var response *CheckHealthOK
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    CheckHealthOperation,
			OperationSummary: "",
			OperationID:      "checkHealth",
			Body:             nil,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = *CheckHealthOK
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.CheckHealth(ctx)
				return response, err
			},
		)
	} else {
		response, err = s.h.CheckHealth(ctx)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeCheckHealthResponse(response, w); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleRefreshTokenRequest handles refreshToken operation.
//
// POST /refresh-token
func (s *Server) handleRefreshTokenRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	ctx := r.Context()

	var (
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: RefreshTokenOperation,
			ID:   "refreshToken",
		}
	)
	request, close, err := s.decodeRefreshTokenRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response RefreshTokenRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    RefreshTokenOperation,
			OperationSummary: "",
			OperationID:      "refreshToken",
			Body:             request,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *RefreshTokenReq
			Params   = struct{}
			Response = RefreshTokenRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.RefreshToken(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.RefreshToken(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeRefreshTokenResponse(response, w); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleSignInRequest handles signIn operation.
//
// POST /sign-in
func (s *Server) handleSignInRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	ctx := r.Context()

	var (
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: SignInOperation,
			ID:   "signIn",
		}
	)
	request, close, err := s.decodeSignInRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response SignInRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    SignInOperation,
			OperationSummary: "",
			OperationID:      "signIn",
			Body:             request,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *SignInReq
			Params   = struct{}
			Response = SignInRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.SignIn(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.SignIn(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeSignInResponse(response, w); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleSignUpRequest handles signUp operation.
//
// POST /sign-up
func (s *Server) handleSignUpRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	ctx := r.Context()

	var (
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: SignUpOperation,
			ID:   "signUp",
		}
	)
	request, close, err := s.decodeSignUpRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response SignUpRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    SignUpOperation,
			OperationSummary: "",
			OperationID:      "signUp",
			Body:             request,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *SignUpReq
			Params   = struct{}
			Response = SignUpRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.SignUp(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.SignUp(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeSignUpResponse(response, w); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}
