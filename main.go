package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "embed"

	"github.com/FedotCompot/file-cacher/internal/cache"
	"github.com/FedotCompot/file-cacher/internal/web"
)

func init() {
	if os.Getenv("DEBUG") == "true" {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}
}

func main() {
	startTime := time.Now()
	mainCtx := context.Background()

	slog.Info("Starting...")

	// Cache
	if err := cache.Initialize(mainCtx); err != nil {
		slog.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}
	slog.Info("Connected to redis", "url", os.Getenv("REDIS_URL"))

	// HTTP Server
	web.Start()

	slog.Info("Started!", "in", time.Since(startTime))
	waitExitSignal()
	slog.Info("Shutting down...")

	if err := web.Stop(); err != nil {
		slog.Error("Failed to stop web server", "error", err)
	}

	if err := cache.Close(); err != nil {
		slog.Error("Failed to close cache", "error", err)
	}
}

func waitExitSignal() os.Signal {
	ch := make(chan os.Signal, 3)
	signal.Notify(
		ch,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	return <-ch
}
