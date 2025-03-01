package main

//go:generate go tool ogen -clean -config ../.ogen.yaml -package zondaapi -target ../lib/go/zondaapi ./openapi.yaml

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/minguu42/zonda/api/applog"
	"github.com/minguu42/zonda/api/config"
	"github.com/minguu42/zonda/api/handler"
)

func init() {
	time.Local = time.FixedZone("JST", 9*60*60)
}

func main() {
	ctx := context.Background()
	if err := mainRun(ctx); err != nil {
		applog.Error(ctx, err.Error())
		os.Exit(1)
	}
}

func mainRun(ctx context.Context) error {
	var conf config.Config
	if err := config.LoadEnv(&conf); err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	h, err := handler.New()
	if err != nil {
		return fmt.Errorf("failed to create handler: %w", err)
	}
	s := &http.Server{
		Addr:         net.JoinHostPort(conf.API.Host, strconv.Itoa(conf.API.Port)),
		Handler:      h,
		ReadTimeout:  conf.API.ReadTimeout,
		WriteTimeout: conf.API.WriteTimeout,
	}

	serveErr := make(chan error)
	go func() {
		applog.Event(ctx, "Start accepting requests")
		serveErr <- s.ListenAndServe()
	}()

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGTERM)
	select {
	case err := <-serveErr:
		return fmt.Errorf("failed to listen and serve: %w", err)
	case <-sigterm:
	}

	ctxWithTimeout, cancel := context.WithTimeout(ctx, conf.API.StopTimeout)
	defer cancel()
	if err := s.Shutdown(ctxWithTimeout); err != nil {
		return fmt.Errorf("failed to shutdown server: %w", err)
	}
	applog.Event(ctx, "Stop accepting requests")
	return nil
}
