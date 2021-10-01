package router

import (
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/satimoto/go-api/handler"
	"github.com/satimoto/go-datastore/db"
)

func Initialize(repo db.Repository) *chi.Mux {
	router := chi.NewRouter()

	// Set middleware
	router.Use(render.SetContentType(render.ContentTypeJSON), middleware.RedirectSlashes, middleware.Recoverer)

	router.Use(middleware.Timeout(30 * time.Second))

	// Adds routes
	router.Route("/v1", func(r chi.Router) {
		r.Mount("/", handler.Routes(repo))
	})

	return router
}
