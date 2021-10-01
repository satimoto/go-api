package handler

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/graph/resolver"
	"github.com/satimoto/go-datastore/db"
)

// Routes initializes the handlers for the router
func Routes(repo db.Repository) *chi.Mux {
	router := chi.NewRouter()

	router.Post("/query", graphQLHandler(repo))

	router.Get("/playground", playgroundQLHandler("/v1/query"))

	return router
}

func graphQLHandler(repo db.Repository) http.HandlerFunc {
	graphQLServer := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{
		Repository: repo,
	}}))

	return graphQLServer.ServeHTTP
}

func playgroundQLHandler(endpoint string) http.HandlerFunc {
	playgroundHandler := playground.Handler("GraphQL Playground", endpoint)

	return playgroundHandler
}
