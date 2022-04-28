package token

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type TokenRepository interface {
	CreateToken(ctx context.Context, arg db.CreateTokenParams) (db.Token, error)
	GetToken(ctx context.Context, id int64) (db.Token, error)
	GetTokenByAuthID(ctx context.Context, authID string) (db.Token, error)
}

type TokenResolver struct {
	Repository TokenRepository
}

func NewResolver(repositoryService *db.RepositoryService) *TokenResolver {
	repo := TokenRepository(repositoryService)
	return &TokenResolver{repo}
}
