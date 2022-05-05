package token

import (
	"context"

	"github.com/satimoto/go-datastore/db"
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
	repo := TokenRepository(repositoryService)

	return &TokenResolver{
		Repository:  repo,
		OcpiService: ocpi.NewService(),
	}
}
