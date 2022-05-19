package resolver

import (
	"os"

	"github.com/satimoto/go-api/internal/authentication"
	"github.com/satimoto/go-api/internal/businessdetail"
	"github.com/satimoto/go-api/internal/channelrequest"
	"github.com/satimoto/go-api/internal/emailsubscription"
	"github.com/satimoto/go-api/internal/energymix"
	"github.com/satimoto/go-api/internal/evse"
	"github.com/satimoto/go-api/internal/location"
	"github.com/satimoto/go-api/internal/openingtime"
	"github.com/satimoto/go-api/internal/token"
	"github.com/satimoto/go-api/internal/user"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/image"
	"github.com/satimoto/go-datastore/pkg/node"
	"github.com/satimoto/go-ocpi-api/pkg/ocpi"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Repository interface{}

type Resolver struct {
	Repository
	OcpiService               ocpi.Ocpi
	AuthenticationResolver    *authentication.AuthenticationResolver
	BusinessDetailResolver    *businessdetail.BusinessDetailResolver
	ChannelRequestResolver    *channelrequest.ChannelRequestResolver
	EmailSubscriptionResolver *emailsubscription.EmailSubscriptionResolver
	EnergyMixResolver         *energymix.EnergyMixResolver
	EvseResolver              *evse.EvseResolver
	ImageRepository           image.ImageRepository
	LocationResolver          *location.LocationResolver
	NodeRepository            node.NodeRepository
	OpeningTimeResolver       *openingtime.OpeningTimeResolver
	TokenResolver             *token.TokenResolver
	UserResolver              *user.UserResolver
}

func NewResolver(repositoryService *db.RepositoryService) *Resolver {
	ocpiService := ocpi.NewService(os.Getenv("OCPI_RPC_ADDRESS"))

	return NewResolverWithServices(repositoryService, ocpiService)
}

func NewResolverWithServices(repositoryService *db.RepositoryService, ocpiService ocpi.Ocpi) *Resolver {
	repo := Repository(repositoryService)

	return &Resolver{
		Repository:                repo,
		OcpiService:               ocpiService,
		AuthenticationResolver:    authentication.NewResolver(repositoryService),
		BusinessDetailResolver:    businessdetail.NewResolver(repositoryService),
		ChannelRequestResolver:    channelrequest.NewResolver(repositoryService),
		EmailSubscriptionResolver: emailsubscription.NewResolver(repositoryService),
		EnergyMixResolver:         energymix.NewResolver(repositoryService),
		EvseResolver:              evse.NewResolver(repositoryService),
		ImageRepository:           image.NewRepository(repositoryService),
		LocationResolver:          location.NewResolver(repositoryService),
		NodeRepository:            node.NewRepository(repositoryService),
		OpeningTimeResolver:       openingtime.NewResolver(repositoryService),
		TokenResolver:             token.NewResolver(repositoryService),
		UserResolver:              user.NewResolver(repositoryService),
	}
}
