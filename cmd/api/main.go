package main

import (
	"skeleton/internal/router"

	"go.uber.org/zap"
)

func main() {
	cfg := config{
		addr: ":3201",
	}

	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	app := &application{
		config: cfg,
		logger: logger,
	}

	r := router.InitRouter()

	logger.Fatal(app.run(r))
}
