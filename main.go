package main

import (
	"context"
	"http-template/api"
	"http-template/middlewares"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		logger.Info("Received shutdown signal")
		cancel()
	}()

	appCtx := middlewares.NewAppContext(ctx, logger)

	err := api.StartServer(appCtx)
	if err != nil {
		logger.Error("failed to start server", "err", err)
		return
	}
}
