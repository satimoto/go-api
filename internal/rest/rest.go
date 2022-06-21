package rest

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/edjumacator/chi-prometheus"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/satimoto/go-api/internal/authentication"
	"github.com/satimoto/go-api/internal/ferp"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
)

type Rest interface {
	Handler() *chi.Mux
	StartRest(context.Context, *sync.WaitGroup)
}

type RestService struct {
	RepositoryService *db.RepositoryService
	FerpService       ferp.Ferp
	Server            *http.Server
}

func NewRest(d *sql.DB, ferpService ferp.Ferp) Rest {
	return &RestService{
		RepositoryService: db.NewRepositoryService(d),
		FerpService:       ferpService,
	}
}

func (rs *RestService) Handler() *chi.Mux {
	router := chi.NewRouter()

	// Set middleware
	router.Use(render.SetContentType(render.ContentTypeJSON), middleware.Logger, middleware.RedirectSlashes, middleware.Recoverer)
	router.Use(middleware.Timeout(120 * time.Second))
	router.Use(authentication.AuthorizationContext())
	router.Use(chiprometheus.NewMiddleware("api"))

	router.Use(cors.Handler(cors.Options{
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Mount("/health", rs.mountHealth())
	router.Mount("/metrics", promhttp.Handler())
	router.Mount("/v1", rs.mountV1())

	return router
}

func (rs *RestService) StartRest(ctx context.Context, waitGroup *sync.WaitGroup) {
	if rs.Server == nil {
		rs.Server = &http.Server{
			Addr:    fmt.Sprintf(":%s", os.Getenv("REST_PORT")),
			Handler: rs.Handler(),
		}
	}

	log.Printf("Starting Rest service")
	waitGroup.Add(1)

	go rs.listenAndServe()

	go func() {
		<-ctx.Done()
		log.Printf("Shutting down Rest service")

		rs.shutdown()

		log.Printf("Rest service shut down")
		waitGroup.Done()
	}()
}

func (rs *RestService) listenAndServe() {
	err := rs.Server.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		log.Printf("Error in Rest service: %v", err)
	}
}

func (rs *RestService) shutdown() {
	timeout := util.GetEnvInt32("SHUTDOWN_TIMEOUT", 20)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	err := rs.Server.Shutdown(ctx)

	if err != nil {
		log.Printf("Error shutting down Rest service: %v", err)
	}
}
