package auth

import (
	"github.com/satimoto/go-api/internal/authentication"
	"github.com/satimoto/go-datastore/pkg/db"
)

type LnUrlAuthRepository interface {
}

type LnUrlAuthResolver struct {
	Repository             LnUrlAuthRepository
	AuthenticationResolver *authentication.AuthenticationResolver
}

func NewResolver(repositoryService *db.RepositoryService) *LnUrlAuthResolver {
	repo := LnUrlAuthRepository(repositoryService)

	return &LnUrlAuthResolver{
		Repository:             repo,
		AuthenticationResolver: authentication.NewResolver(repositoryService),
	}
}
