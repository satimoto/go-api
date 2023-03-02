package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	_ "github.com/joho/godotenv/autoload"
	"github.com/satimoto/go-api/internal/ferp"
	metrics "github.com/satimoto/go-api/internal/metric"
	"github.com/satimoto/go-api/internal/rest"
	syncronizer "github.com/satimoto/go-api/internal/sync"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
)

var (
	database *sql.DB

	dbHost  = os.Getenv("DB_HOST")
	dbName  = os.Getenv("DB_NAME")
	dbPass  = os.Getenv("DB_PASS")
	dbUser  = os.Getenv("DB_USER")
	sslMode = util.GetEnv("SSL_MODE", "disable")
)

func init() {
	if len(dbHost) == 0 || len(dbName) == 0 || len(dbPass) == 0 || len(dbUser) == 0 {
		log.Fatalf("Database env variables not defined")
	}

	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s/%s?binary_parameters=yes&sslmode=%s", dbUser, dbPass, dbHost, dbName, sslMode)
	d, err := sql.Open("postgres", dataSourceName)

	if err != nil {
		log.Fatal(err)
	}

	database = d
}

func main() {
	defer database.Close()

	log.Printf("Starting up API server")
	shutdownCtx, cancelFunc := context.WithCancel(context.Background())
	waitGroup := &sync.WaitGroup{}
	
	repositoryService := db.NewRepositoryService(database)

	ferpService := ferp.NewService(os.Getenv("FERP_RPC_ADDRESS"))
	ferpService.Start(shutdownCtx, waitGroup)

	metricsService := metrics.NewMetrics()
	metricsService.StartMetrics(shutdownCtx, waitGroup)

	restService := rest.NewRest(repositoryService, ferpService)
	restService.StartRest(shutdownCtx, waitGroup)

	syncService := syncronizer.NewService(repositoryService)
	syncService.StartService(shutdownCtx, waitGroup)

	sigtermChan := make(chan os.Signal)
	signal.Notify(sigtermChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigtermChan

	log.Printf("Shutting down API server")

	cancelFunc()
	waitGroup.Wait()

	log.Printf("API server shut down")
}
