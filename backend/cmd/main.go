package main

import (
	"context"
	"errors"
	"github.com/Max425/manager/internal/config"
	"github.com/Max425/manager/internal/httpserver"
	"github.com/Max425/manager/internal/lib/logger/handlers/slogpretty"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

func run() error {
	// read config
	cfg := config.MustLoad()

	logger := setupLogger(cfg.Env)

	// create http server with all handlers & services & repositories
	srv, err := httpserver.NewHttpServer(logger, cfg.Postgres, cfg.HttpAddr)
	if err != nil {
		logger.Error("create http server: %s", err)
		return err
	}

	// listen to OS signals and gracefully shutdown HTTP server
	stopped := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-sigint
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err = srv.Shutdown(ctx); err != nil {
			logger.Error("HTTP Server Shutdown Error: %v", err)
		}
		close(stopped)
	}()

	logger.Info("Starting HTTP server on %s", cfg.HttpAddr)

	// start HTTP server
	if err = srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		logger.Error("HTTP server ListenAndServe Error: %v", err)
	}

	<-stopped

	return nil
}

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case envLocal:
		logger = setupPrettySlog()
	case envDev:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return logger
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}