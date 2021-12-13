package router

import (
	"database/sql"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/satimoto/go-api/graph/resolver"
	"github.com/satimoto/go-datastore/db"
)

func Initialize(d *sql.DB) *chi.Mux {
	repositoryService := db.NewRepositoryService(d)
	router := chi.NewRouter()

	// Set middleware
	router.Use(render.SetContentType(render.ContentTypeJSON), middleware.RedirectSlashes, middleware.Recoverer)
	router.Use(middleware.Timeout(30 * time.Second))

	router.Use(cors.Handler(cors.Options{
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge: 300,
	}))

	// Adds routes
	router.Route("/v1", func(r chi.Router) {
		r.Mount("/", resolver.New(repositoryService))
	})

	return router
}
