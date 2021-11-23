# go-api
Satimoto public API using golang

## On schema changes
Whenever the SQL or GraphQL schema files change the following should be run

### Generate the GraphQL resolvers
Generates the GraphQL models and resolvers from schema files
```bash
gqlgen
```

## Development

### Run
```bash
go run ./cmd/locald
```
