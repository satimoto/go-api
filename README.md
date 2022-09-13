# go-api
Satimoto public API using golang

## On schema changes
Whenever the SQL or GraphQL schema files change the following should be run

## Config
Some dependencies require using bzr repository
```
brew install bzr
go env -w GOVCS="launchpad.net:bzr,public:git|hg,private:all"
```

### Generate the GraphQL resolvers
Generates the GraphQL models and resolvers from schema files
```bash
go mod download github.com/99designs/gqlgen
gqlgen
```

## Development

### Run
```bash
go run ./cmd/api
```

## Build
```bash
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s -w' -o bin/main cmd/api/main.go
```