package resolver

import (
	"github.com/satimoto/go-api/emailsubscription"
	"github.com/satimoto/go-api/node"
	"github.com/satimoto/go-api/user"
	"github.com/satimoto/go-datastore/db"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Repository interface{}

type Resolver struct {
	Repository
	*emailsubscription.EmailSubscriptionResolver
	*node.NodeResolver
	*user.UserResolver
}

func NewResolver(repositoryService *db.RepositoryService) *Resolver {
	repo := Repository(repositoryService)
	return &Resolver{
		Repository:                repo,
		EmailSubscriptionResolver: emailsubscription.NewResolver(repositoryService),
		NodeResolver:              node.NewResolver(repositoryService),
		UserResolver:              user.NewResolver(repositoryService),
	}
}
