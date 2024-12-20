package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	database "skeleton/internal/db"
	"skeleton/internal/router"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {
	appAddr := ":3201"

	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	database.InitDBConnections(logger)
	r := router.InitRouter()

	srv := &http.Server{
		Addr:         appAddr,
		Handler:      r,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	shutdown := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		sig := <-quit
		logger.Infow("Signal received, shutting down", "signal", sig)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		shutdown <- srv.Shutdown(ctx)
	}()

	logger.Infow("Server has started", "addr", appAddr)

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		logger.Infow("ListenAndServe error", "error", err)
	}

	err = <-shutdown
	if err != nil {
		logger.Fatal("Unable to shutdown gracefully")
		panic(err)
	}

	logger.Infow("server has stopped", "addr", appAddr)
}
