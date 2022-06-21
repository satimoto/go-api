package rest

import (
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/satimoto/go-api/graph/resolver"
	"github.com/satimoto/go-ocpi-api/pkg/ocpi"
)

func (rs *RestService) mountGraphql() *chi.Mux {
	ocpiService := ocpi.NewService(os.Getenv("OCPI_RPC_ADDRESS"))
	r := resolver.NewResolverWithServices(rs.RepositoryService, rs.FerpService, ocpiService)
	router := chi.NewRouter()

	router.Post("/query", r.GraphQLHandler())
	router.Get("/playground", r.PlaygroundQLHandler("/v1/query"))

	return router
}
