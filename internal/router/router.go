package router

import (
	"net/http"
	"skeleton/internal/controller"
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

	r.Route("/user", func(r chi.Router) {
		r.Get("/list", controller.GetAllUsers)
		r.Post("/create", controller.CreateUser)
	})

	r.Route("/student", func(r chi.Router) {
		r.Post("/create", controller.CreateStudentAndUser)
	})

	return r
}
