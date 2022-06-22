package resolver

import (
	"os"

	"github.com/satimoto/go-api/internal/authentication"
	"github.com/satimoto/go-api/internal/aws/email"
	"github.com/satimoto/go-api/internal/ferp"
	"github.com/satimoto/go-api/internal/token"
	"github.com/satimoto/go-datastore/pkg/businessdetail"
	"github.com/satimoto/go-datastore/pkg/channelrequest"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/displaytext"
	"github.com/satimoto/go-datastore/pkg/emailsubscription"
	"github.com/satimoto/go-datastore/pkg/energymix"
	"github.com/satimoto/go-datastore/pkg/evse"
	"github.com/satimoto/go-datastore/pkg/image"
	"github.com/satimoto/go-datastore/pkg/location"
	"github.com/satimoto/go-datastore/pkg/node"
	"github.com/satimoto/go-datastore/pkg/openingtime"
	"github.com/satimoto/go-datastore/pkg/tariff"
	"github.com/satimoto/go-datastore/pkg/user"
	"github.com/satimoto/go-ocpi-api/pkg/ocpi"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Repository interface{}

type Resolver struct {
	Repository
	FerpService                 ferp.Ferp
	OcpiService                 ocpi.Ocpi
	Emailer                     email.Emailer
	AuthenticationResolver      *authentication.AuthenticationResolver
	BusinessDetailRepository    businessdetail.BusinessDetailRepository
	ChannelRequestRepository    channelrequest.ChannelRequestRepository
	DisplayTextRepository       displaytext.DisplayTextRepository
	EmailSubscriptionRepository emailsubscription.EmailSubscriptionRepository
	EnergyMixRepository         energymix.EnergyMixRepository
	EvseRepository              evse.EvseRepository
	ImageRepository             image.ImageRepository
	LocationRepository          location.LocationRepository
	NodeRepository              node.NodeRepository
	OpeningTimeRepository       openingtime.OpeningTimeRepository
	TariffRepository            tariff.TariffRepository
	TokenResolver               *token.TokenResolver
	UserRepository              user.UserRepository
}

func NewResolver(repositoryService *db.RepositoryService) *Resolver {
	ferpService := ferp.NewService(os.Getenv("FERP_RPC_ADDRESS"))
	ocpiService := ocpi.NewService(os.Getenv("OCPI_RPC_ADDRESS"))

	return NewResolverWithServices(repositoryService, ferpService, ocpiService)
}

func NewResolverWithServices(repositoryService *db.RepositoryService, ferpService ferp.Ferp, ocpiService ocpi.Ocpi) *Resolver {
	repo := Repository(repositoryService)
	emailer := email.New(os.Getenv("REPLY_TO_EMAIL"))

	return &Resolver{
		Repository:                  repo,
		FerpService:                 ferpService,
		OcpiService:                 ocpiService,
		Emailer:                     emailer,
		AuthenticationResolver:      authentication.NewResolver(repositoryService),
		BusinessDetailRepository:    businessdetail.NewRepository(repositoryService),
		ChannelRequestRepository:    channelrequest.NewRepository(repositoryService),
		DisplayTextRepository:       displaytext.NewRepository(repositoryService),
		EmailSubscriptionRepository: emailsubscription.NewRepository(repositoryService),
		EnergyMixRepository:         energymix.NewRepository(repositoryService),
		EvseRepository:              evse.NewRepository(repositoryService),
		ImageRepository:             image.NewRepository(repositoryService),
		LocationRepository:          location.NewRepository(repositoryService),
		NodeRepository:              node.NewRepository(repositoryService),
		OpeningTimeRepository:       openingtime.NewRepository(repositoryService),
		TariffRepository:            tariff.NewRepository(repositoryService),
		TokenResolver:               token.NewResolver(repositoryService),
		UserRepository:              user.NewRepository(repositoryService),
	}
}
