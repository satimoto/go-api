package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/satimoto/go-api/router"
	"github.com/satimoto/go-datastore/util"
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

	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s", dbUser, dbPass, dbHost, dbName, sslMode)
	d, err := sql.Open("postgres", dataSourceName)
	
	if err != nil {
		log.Fatal(err)
	}

	database = d
}

func main() {
	defer database.Close()

	routerService := router.NewRouter(database)
	handler := routerService.Handler()
	
	err := http.ListenAndServe(":9000", handler)

	if err != nil {
		log.Println("Error serving")
	}
}
