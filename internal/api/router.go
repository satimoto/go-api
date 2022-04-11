package api

import (
	"database/sql"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"

	"github.com/satimoto/go-api/internal/authentication"
	"github.com/satimoto/go-datastore/db"
)

type Router interface {
	Handler() *chi.Mux
}

type RouterService struct {
	*db.RepositoryService
}

func NewRouter(d *sql.DB) Router {
	return &RouterService{
		RepositoryService: db.NewRepositoryService(d),
	}
}

func (rs *RouterService) Handler() *chi.Mux {
	router := chi.NewRouter()

	// Set middleware
	router.Use(render.SetContentType(render.ContentTypeJSON), middleware.Logger, middleware.RedirectSlashes, middleware.Recoverer)
	router.Use(middleware.Timeout(120 * time.Second))
	router.Use(authentication.AuthorizationContext())

	router.Use(cors.Handler(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Mount("/v1", rs.mountV1())

	return router
}
