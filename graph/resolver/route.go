package resolver

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-datastore/db"
)


func New(repositoryService *db.RepositoryService) *chi.Mux {
	r := NewResolver(repositoryService)

	return r.routes()
}

// Routes initializes the handlers for the router
func (r *Resolver) routes() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/query", r.graphQLHandler())
	router.Get("/playground", r.playgroundQLHandler("/v1/query"))

	return router
}

func (r *Resolver) graphQLHandler() http.HandlerFunc {
	graphQLServer := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: r}))

	return graphQLServer.ServeHTTP
}

func (r *Resolver) playgroundQLHandler(endpoint string) http.HandlerFunc {
	playgroundHandler := playground.Handler("GraphQL Playground", endpoint)

	return playgroundHandler
}
