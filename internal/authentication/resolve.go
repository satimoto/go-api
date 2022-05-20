package authentication

import (
	"github.com/satimoto/go-datastore/pkg/authentication"
	"github.com/satimoto/go-datastore/pkg/db"
)

type AuthenticationResolver struct {
	Repository authentication.AuthenticationRepository
}

func NewResolver(repositoryService *db.RepositoryService) *AuthenticationResolver {
	return &AuthenticationResolver{
		Repository: authentication.NewRepository(repositoryService),
	}
}
