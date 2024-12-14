package main

import (
	"log"
	database "skeleton/internal/db"
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

	database.InitDBConnections(logger)

	r := router.InitRouter()

	if err := app.run(r); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
