package main

import (
	"errors"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type config struct {
	addr string
}

type application struct {
	config config
	logger *zap.SugaredLogger
}

func (app *application) run(mux http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	// shutdown := make(chan error)

	// go func() {
	// 	quit := make(chan os.Signal, 1)
	//
	// 	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//
	// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// 	defer cancel()
	//
	// 	shutdown <- srv.Shutdown(ctx)
	// }()

	app.logger.Infow("Server has started", "addr", app.config.addr)

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	// err = <-shutdown
	// if err != nil {
	// 	return err
	// }

	app.logger.Infow("server has stopped", "addr", app.config.addr)
	return nil
}
