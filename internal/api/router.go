package api

import (
	"database/sql"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"

	"github.com/satimoto/go-api/graph/resolver"
	"github.com/satimoto/go-api/internal/authentication"
	"github.com/satimoto/go-api/internal/lnurl/auth"
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
	router.Use(authentication.Middleware())

	router.Use(cors.Handler(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Adds routes
	router.Route("/v1", func(r chi.Router) {
		r.Mount("/", rs.grapqlRouter())
		r.Mount("/lnurl", rs.lnurlRouter())
	})

	return router
}

func (rs *RouterService) grapqlRouter() *chi.Mux {
	r := resolver.NewResolver(rs.RepositoryService)
	router := chi.NewRouter()

	router.Post("/query", r.GraphQLHandler())
	router.Get("/playground", r.PlaygroundQLHandler("/v1/query"))

	return router
}

func (rs *RouterService) lnurlRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Mount("/auth", rs.lnurlAuthRouter())

	return router
}

func (rs *RouterService) lnurlAuthRouter() *chi.Mux {
	r := auth.NewResolver(rs.RepositoryService)

	router := chi.NewRouter()
	router.Get("/", r.GetHandler)

	return router
}
