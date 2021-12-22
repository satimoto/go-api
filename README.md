# go-api
Satimoto public API using golang

## On schema changes
Whenever the SQL or GraphQL schema files change the following should be run

### Generate the GraphQL resolvers
Generates the GraphQL models and resolvers from schema files
```bash
go mod download github.com/99designs/gqlgen
gqlgen
```

## Development

### Run
```bash
go run ./cmd/locald
```

## Build

### Run
```bash
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s -w' -o bin/main cmd/satimotod/main.go
```