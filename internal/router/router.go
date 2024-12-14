package router

import (
	"encoding/json"
	"net/http"
	"skeleton/internal/model"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type MongoClientUris struct {
	userClientUri string
}

func InitRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/health", func(r chi.Router) {
		r.Get("/check", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("All ok!"))
		})
	})

	// MongoClientUris := MongoClientUris{
	// 	userClientUri: config.LoadConfig("MONGO_DB1_URI"),
	// }
	//
	// user_client, _ := database.GetMongoClient(MongoClientUris.userClientUri)
	// user_handler := handler.NewHandler(user_client)
	//
	r.Route("/user", func(r chi.Router) {
		r.Get("/list", func(w http.ResponseWriter, r *http.Request) {
			results, err := model.UserModel().GetAllUser()
			if err != nil {
				panic(err)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(results)
		})
	})

	return r
}
