package token

import (
	"os"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/token"
	"github.com/satimoto/go-ocpi-api/pkg/ocpi"
)

type TokenResolver struct {
	Repository  token.TokenRepository
	OcpiService ocpi.Ocpi
}

func NewResolver(repositoryService *db.RepositoryService) *TokenResolver {
	ocpiService := ocpi.NewService(os.Getenv("OCPI_RPC_ADDRESS"))

	return NewResolverWithServices(repositoryService, ocpiService)
}

func NewResolverWithServices(repositoryService *db.RepositoryService, ocpiService ocpi.Ocpi) *TokenResolver {
	return &TokenResolver{
		Repository:  token.NewRepository(repositoryService),
		OcpiService: ocpiService,
	}
}
