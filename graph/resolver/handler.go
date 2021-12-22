package resolver

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/satimoto/go-api/graph"
)

func (r *Resolver) GraphQLHandler() http.HandlerFunc {
	graphQLServer := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: r}))

	return graphQLServer.ServeHTTP
}

func (r *Resolver) PlaygroundQLHandler(endpoint string) http.HandlerFunc {
	playgroundHandler := playground.Handler("GraphQL Playground", endpoint)

	return playgroundHandler
}
