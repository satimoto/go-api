package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/satimoto/go-api/graph/resolver"
)

func (rs *RestService) mountGraphql() *chi.Mux {
	r := resolver.NewResolver(rs.RepositoryService)
	router := chi.NewRouter()

	router.Post("/query", r.GraphQLHandler())
	router.Get("/playground", r.PlaygroundQLHandler("/v1/query"))

	return router
}
