package web

import (
	"errors"
	"log/slog"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/FedotCompot/file-cacher/internal/config"
	"github.com/klauspost/compress/gzhttp"
)

var httpServer *http.Server

func Start() {
	listener, err := net.Listen("tcp", config.Data.Listener)
	if err != nil {
		slog.Error("Failed to listen", "error", err)
		os.Exit(1)
	}

	router := getRouter()
	httpServer = &http.Server{
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       5 * time.Second,
		IdleTimeout:       60 * time.Second,
		Handler:           gzhttp.GzipHandler(http.Handler(router)),
	}
	slog.Info("Starting web server", "address", config.Data.Listener)
	go func() {
		if err := httpServer.Serve(listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("Failed to start web server", "error", err)
			os.Exit(1)
		}
	}()
}

func Stop() error {
	if httpServer == nil {
		return nil
	}
	err := httpServer.Close()
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}
	return err
}
