package rest

import (
	"context"
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
	"github.com/satimoto/go-api/internal/ferp"
	metrics "github.com/satimoto/go-api/internal/metric"
	apiMiddleware "github.com/satimoto/go-api/internal/middleware"
	syncronizer "github.com/satimoto/go-api/internal/sync"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/user"

	"github.com/satimoto/go-datastore/pkg/util"
)

type Rest interface {
	Handler() *chi.Mux
	StartRest(context.Context, *sync.WaitGroup)
}

type RestService struct {
	RepositoryService *db.RepositoryService
	UserRepository    user.UserRepository
	FerpService       ferp.Ferp
	SyncService       syncronizer.Sync
	Server            *http.Server
}

func NewRest(repositoryService *db.RepositoryService, ferpService ferp.Ferp, syncService syncronizer.Sync) Rest {
	return &RestService{
		RepositoryService: repositoryService,
		UserRepository:    user.NewRepository(repositoryService),
		FerpService:       ferpService,
		SyncService:       syncService,
	}
}

func (rs *RestService) Handler() *chi.Mux {
	router := chi.NewRouter()

	// Set middleware
	router.Use(render.SetContentType(render.ContentTypeJSON), middleware.RedirectSlashes, middleware.Recoverer)
	router.Use(middleware.Timeout(120 * time.Second))
	router.Use(apiMiddleware.AuthorizationContext())
	router.Use(chiprometheus.NewMiddleware("api"))

	router.Use(cors.Handler(cors.Options{
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Mount("/health", rs.mountHealth())
	router.Mount("/sync", rs.mountSync())
	router.Mount("/v1", rs.mountV1())
	router.Mount("/.well-known", rs.mountWellKnown())

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
		metrics.RecordError("API023", "Error in Rest service", err)
	}
}

func (rs *RestService) shutdown() {
	timeout := util.GetEnvInt32("SHUTDOWN_TIMEOUT", 20)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	err := rs.Server.Shutdown(ctx)

	if err != nil {
		metrics.RecordError("API024", "Error shutting down Rest service", err)
	}
}
