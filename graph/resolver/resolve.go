package resolver

import (
	"github.com/satimoto/go-api/internal/authentication"
	"github.com/satimoto/go-api/internal/businessdetail"
	"github.com/satimoto/go-api/internal/channelrequest"
	"github.com/satimoto/go-api/internal/emailsubscription"
	"github.com/satimoto/go-api/internal/energymix"
	"github.com/satimoto/go-api/internal/evse"
	"github.com/satimoto/go-api/internal/image"
	"github.com/satimoto/go-api/internal/location"
	"github.com/satimoto/go-api/internal/node"
	"github.com/satimoto/go-api/internal/openingtime"
	"github.com/satimoto/go-api/internal/token"
	"github.com/satimoto/go-api/internal/user"
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/pkg/ocpi"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Repository interface{}

type Resolver struct {
	Repository
	OcpiService ocpi.Ocpi
	*authentication.AuthenticationResolver
	*businessdetail.BusinessDetailResolver
	*channelrequest.ChannelRequestResolver
	*emailsubscription.EmailSubscriptionResolver
	*energymix.EnergyMixResolver
	*evse.EvseResolver
	*image.ImageResolver
	*location.LocationResolver
	*node.NodeResolver
	*openingtime.OpeningTimeResolver
	*token.TokenResolver
	*user.UserResolver
}

func NewResolver(repositoryService *db.RepositoryService) *Resolver {
	repo := Repository(repositoryService)
	return &Resolver{
		Repository:                repo,
		OcpiService:               ocpi.NewService(),
		AuthenticationResolver:    authentication.NewResolver(repositoryService),
		BusinessDetailResolver:    businessdetail.NewResolver(repositoryService),
		ChannelRequestResolver:    channelrequest.NewResolver(repositoryService),
		EmailSubscriptionResolver: emailsubscription.NewResolver(repositoryService),
		EnergyMixResolver:         energymix.NewResolver(repositoryService),
		EvseResolver:              evse.NewResolver(repositoryService),
		ImageResolver:             image.NewResolver(repositoryService),
		LocationResolver:          location.NewResolver(repositoryService),
		NodeResolver:              node.NewResolver(repositoryService),
		OpeningTimeResolver:       openingtime.NewResolver(repositoryService),
		TokenResolver:             token.NewResolver(repositoryService),
		UserResolver:              user.NewResolver(repositoryService),
	}
}
