package rest

import (
	"github.com/go-chi/chi/v5"

	"github.com/satimoto/go-api/internal/lnurl/auth"
)

func (rs *RestService) mountLnUrl() *chi.Mux {
	router := chi.NewRouter()
	router.Mount("/auth", rs.mountLnUrlAuth())

	return router
}

func (rs *RestService) mountLnUrlAuth() *chi.Mux {
	r := auth.NewResolver(rs.RepositoryService)

	router := chi.NewRouter()
	router.Get("/", r.GetLnUrlAuth)

	return router
}
