package resolver

import (
	"github.com/satimoto/go-api/internal/authentication"
	"github.com/satimoto/go-api/internal/businessdetail"
	"github.com/satimoto/go-api/internal/channelrequest"
	"github.com/satimoto/go-api/internal/credential"
	"github.com/satimoto/go-api/internal/emailsubscription"
	"github.com/satimoto/go-api/internal/energymix"
	"github.com/satimoto/go-api/internal/evse"
	"github.com/satimoto/go-api/internal/image"
	"github.com/satimoto/go-api/internal/location"
	"github.com/satimoto/go-api/internal/node"
	"github.com/satimoto/go-api/internal/openingtime"
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
	*businessdetail.BusinessDetailResolver
	*channelrequest.ChannelRequestResolver
	*credential.CredentialResolver
	*emailsubscription.EmailSubscriptionResolver
	*energymix.EnergyMixResolver
	*evse.EvseResolver
	*image.ImageResolver
	*location.LocationResolver
	*node.NodeResolver
	*openingtime.OpeningTimeResolver
	*user.UserResolver
}

func NewResolver(repositoryService *db.RepositoryService) *Resolver {
	repo := Repository(repositoryService)
	return &Resolver{
		Repository:                repo,
		AuthenticationResolver:    authentication.NewResolver(repositoryService),
		BusinessDetailResolver:    businessdetail.NewResolver(repositoryService),
		ChannelRequestResolver:    channelrequest.NewResolver(repositoryService),
		CredentialResolver:        credential.NewResolver(repositoryService),
		EmailSubscriptionResolver: emailsubscription.NewResolver(repositoryService),
		EnergyMixResolver:         energymix.NewResolver(repositoryService),
		EvseResolver:              evse.NewResolver(repositoryService),
		ImageResolver:             image.NewResolver(repositoryService),
		LocationResolver:          location.NewResolver(repositoryService),
		NodeResolver:              node.NewResolver(repositoryService),
		OpeningTimeResolver:       openingtime.NewResolver(repositoryService),
		UserResolver:              user.NewResolver(repositoryService),
	}
}
