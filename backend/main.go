package main

import (
	"context"
	"errors"
	"github.com/Max425/manager/internal/config"
	"github.com/Max425/manager/internal/http-server"
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

// @title Manager API
// @version 1.0
// @description Web application for automatic compilation of project teams

// @host localhost:8000
// @BasePath /
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
	srv, err := http_server.NewHttpServer(logger, cfg.Postgres, cfg.HttpAddr)
	if err != nil {
		logger.Error("create http server", slog.Any("error", err))
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
			logger.Error("HTTP Server Shutdown", slog.Any("error", err))
		}
		close(stopped)
	}()

	logger.Info("Starting HTTP server", slog.String("addr", cfg.HttpAddr))

	// start HTTP server
	if err = srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		logger.Error("HTTP server ListenAndServe", slog.Any("error", err))
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
