package rest

import (
	"os"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/satimoto/go-api/graph/resolver"
	"github.com/satimoto/go-api/internal/notification"
	"github.com/satimoto/go-ocpi/pkg/ocpi"
)

func (rs *RestService) mountGraphql() *chi.Mux {
	ocpiService := ocpi.NewService(os.Getenv("OCPI_RPC_ADDRESS"))
	notificationService := notification.NewService(os.Getenv("FCM_API_KEY"))
	r := resolver.NewResolverWithServices(rs.RepositoryService, rs.FerpService, notificationService, ocpiService)
	
	router := chi.NewRouter()
	router.Use(middleware.Timeout(120 * time.Second))

	router.Post("/query", r.GraphQLHandler())
	router.Get("/playground", r.PlaygroundQLHandler("/v1/query"))

	return router
}
