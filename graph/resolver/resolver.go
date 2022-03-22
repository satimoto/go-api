package resolver

import (
	"github.com/satimoto/go-api/internal/authentication"
	"github.com/satimoto/go-api/internal/channelrequest"
	"github.com/satimoto/go-api/internal/emailsubscription"
	"github.com/satimoto/go-api/internal/node"
	"github.com/satimoto/go-api/internal/user"
	"github.com/satimoto/go-datastore/db"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Repository interface{}

type Resolver struct {
	Repository
	*authentication.AuthenticationResolver
	*channelrequest.ChannelRequestResolver
	*emailsubscription.EmailSubscriptionResolver
	*node.NodeResolver
	*user.UserResolver
}

func NewResolver(repositoryService *db.RepositoryService) *Resolver {
	repo := Repository(repositoryService)
	return &Resolver{
		Repository:                repo,
		ChannelRequestResolver:    channelrequest.NewResolver(repositoryService),
		AuthenticationResolver:    authentication.NewResolver(repositoryService),
		EmailSubscriptionResolver: emailsubscription.NewResolver(repositoryService),
		NodeResolver:              node.NewResolver(repositoryService),
		UserResolver:              user.NewResolver(repositoryService),
	}
}
