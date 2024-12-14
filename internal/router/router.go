package router

import (
	"net/http"
	"skeleton/config"
	database "skeleton/internal/db"
	"skeleton/internal/handler"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// type MongoClient interface {
// 	client() string
// }
//
// func GetClient(uri string) (*mongo.Client, error) {
// 	return database.GetMongoClient(uri)
// }

type MongoClientUris struct {
	userClientUri string
}

func InitRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/health", func(r chi.Router) {
		r.Get("/check", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("All ok!"))
		})
	})

	MongoClientUris := MongoClientUris{
		userClientUri: config.LoadConfig("MONGO_DB1_URI"),
	}

	user_client, _ := database.GetMongoClient(MongoClientUris.userClientUri)
	user_handler := handler.NewHandler(user_client)

	r.Route("/user", func(r chi.Router) {
		r.Get("/list", user_handler.GetUsers)
	})

	return r
}
