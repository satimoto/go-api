package token

import (
	"context"
	"os"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi-api/pkg/ocpi"
)

type TokenRepository interface {
	CreateToken(ctx context.Context, arg db.CreateTokenParams) (db.Token, error)
	GetToken(ctx context.Context, id int64) (db.Token, error)
	GetTokenByAuthID(ctx context.Context, authID string) (db.Token, error)
}

type TokenResolver struct {
	Repository  TokenRepository
	OcpiService ocpi.Ocpi
}

func NewResolver(repositoryService *db.RepositoryService) *TokenResolver {
	ocpiService := ocpi.NewService(os.Getenv("OCPI_RPC_ADDRESS"))

	return NewResolverWithServices(repositoryService, ocpiService)
}

func NewResolverWithServices(repositoryService *db.RepositoryService, ocpiService ocpi.Ocpi) *TokenResolver {
	repo := TokenRepository(repositoryService)

	return &TokenResolver{
		Repository:  repo,
		OcpiService: ocpiService,
	}
}
