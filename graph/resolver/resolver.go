package resolver

import (
	"github.com/satimoto/go-api/authentication"
	"github.com/satimoto/go-api/emailsubscription"
	"github.com/satimoto/go-api/user"
	"github.com/satimoto/go-datastore/db"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Repository interface{}

type Resolver struct {
	Repository
	*authentication.AuthenticationResolver
	*emailsubscription.EmailSubscriptionResolver
	*user.UserResolver
}

func NewResolver(repositoryService *db.RepositoryService) *Resolver {
	repo := Repository(repositoryService)
	return &Resolver{
		Repository:                repo,
		AuthenticationResolver:    authentication.NewResolver(repositoryService),
		EmailSubscriptionResolver: emailsubscription.NewResolver(repositoryService),
		UserResolver:              user.NewResolver(repositoryService),
	}
}
