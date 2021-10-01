package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	chiproxy "github.com/awslabs/aws-lambda-go-api-proxy/chi"
	"github.com/satimoto/go-api/router"
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-datastore/util"
)

var (
	database  *sql.DB
	chiLambda *chiproxy.ChiLambda

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

	repo := db.NewRepository(database)

	chiLambda = chiproxy.New(router.Initialize(repo))
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return chiLambda.ProxyWithContext(ctx, req)
}

func main() {
	defer database.Close()

	lambda.Start(Handler)
}
