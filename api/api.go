package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/minguu42/zonda/api/applog"
	"github.com/minguu42/zonda/api/handler"
)

func main() {
	ctx := context.Background()
	if err := mainRun(ctx); err != nil {
		applog.Error(ctx, err.Error())
		os.Exit(1)
	}
}

func mainRun(ctx context.Context) error {
	h, err := handler.New()
	if err != nil {
		return fmt.Errorf("failed to create handler: %w", err)
	}
	s := &http.Server{
		Addr:              ":8080",
		Handler:           h,
		ReadTimeout:       2 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}

	serveErr := make(chan error)
	go func() {
		applog.Event(ctx, "Start accepting requests")
		serveErr <- s.ListenAndServe()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM)
	select {
	case err := <-serveErr:
		return fmt.Errorf("failed to listen and serve: %w", err)
	case <-quit:
	}

	if err := s.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown server: %w", err)
	}
	applog.Event(ctx, "Stop accepting requests")
	return nil
}
