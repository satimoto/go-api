package resolver

import (
	"context"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/satimoto/go-api/internal/account"
	"github.com/satimoto/go-api/internal/authentication"
	"github.com/satimoto/go-api/internal/aws/email"
	"github.com/satimoto/go-api/internal/ferp"
	"github.com/satimoto/go-api/internal/notification"
	"github.com/satimoto/go-api/internal/token"
	"github.com/satimoto/go-datastore/pkg/businessdetail"
	"github.com/satimoto/go-datastore/pkg/channelrequest"
	"github.com/satimoto/go-datastore/pkg/connector"
	"github.com/satimoto/go-datastore/pkg/credential"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/displaytext"
	"github.com/satimoto/go-datastore/pkg/emailsubscription"
	"github.com/satimoto/go-datastore/pkg/energymix"
	"github.com/satimoto/go-datastore/pkg/evse"
	"github.com/satimoto/go-datastore/pkg/image"
	"github.com/satimoto/go-datastore/pkg/invoicerequest"
	"github.com/satimoto/go-datastore/pkg/location"
	"github.com/satimoto/go-datastore/pkg/node"
	"github.com/satimoto/go-datastore/pkg/openingtime"
	"github.com/satimoto/go-datastore/pkg/party"
	"github.com/satimoto/go-datastore/pkg/pendingnotification"
	"github.com/satimoto/go-datastore/pkg/poi"
	"github.com/satimoto/go-datastore/pkg/promotion"
	"github.com/satimoto/go-datastore/pkg/referral"
	"github.com/satimoto/go-datastore/pkg/session"
	"github.com/satimoto/go-datastore/pkg/tariff"
	"github.com/satimoto/go-datastore/pkg/tokenauthorization"
	"github.com/satimoto/go-datastore/pkg/user"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi/pkg/ocpi"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Repository interface{}

type Resolver struct {
	Repository
	FerpService                   ferp.Ferp
	NotificationService           notification.Notification
	OcpiService                   ocpi.Ocpi
	Emailer                       email.Emailer
	AccountResolver               *account.AccountResolver
	AuthenticationResolver        *authentication.AuthenticationResolver
	BusinessDetailRepository      businessdetail.BusinessDetailRepository
	ChannelRequestRepository      channelrequest.ChannelRequestRepository
	ConnectorRepository           connector.ConnectorRepository
	CredentialRepository          credential.CredentialRepository
	DisplayTextRepository         displaytext.DisplayTextRepository
	EmailSubscriptionRepository   emailsubscription.EmailSubscriptionRepository
	EnergyMixRepository           energymix.EnergyMixRepository
	EvseRepository                evse.EvseRepository
	ImageRepository               image.ImageRepository
	InvoiceRequestRepository      invoicerequest.InvoiceRequestRepository
	LocationRepository            location.LocationRepository
	NodeRepository                node.NodeRepository
	OpeningTimeRepository         openingtime.OpeningTimeRepository
	ReferralRepository            referral.ReferralRepository
	PartyRepository               party.PartyRepository
	PendingNotificationRepository pendingnotification.PendingNotificationRepository
	PoiRepository                 poi.PoiRepository
	PromotionRepository           promotion.PromotionRepository
	SessionRepository             session.SessionRepository
	TariffRepository              tariff.TariffRepository
	TokenAuthorizationRepository  tokenauthorization.TokenAuthorizationRepository
	TokenResolver                 *token.TokenResolver
	UserRepository                user.UserRepository
	defaultTaxPercent             float64
}

func NewResolver(repositoryService *db.RepositoryService) *Resolver {
	ferpService := ferp.NewService(os.Getenv("FERP_RPC_ADDRESS"))
	notificationService := notification.NewService(os.Getenv("FCM_API_KEY"))
	ocpiService := ocpi.NewService(os.Getenv("OCPI_RPC_ADDRESS"))

	return NewResolverWithServices(repositoryService, ferpService, notificationService, ocpiService)
}

func NewResolverWithServices(repositoryService *db.RepositoryService, ferpService ferp.Ferp, notificationService notification.Notification, ocpiService ocpi.Ocpi) *Resolver {
	repo := Repository(repositoryService)
	emailer := email.New(os.Getenv("REPLY_TO_EMAIL"))
	defaultTaxPercent := util.GetEnvFloat64("DEFAULT_TAX_PERCENT", 19)

	return &Resolver{
		Repository:                    repo,
		FerpService:                   ferpService,
		NotificationService:           notificationService,
		OcpiService:                   ocpiService,
		Emailer:                       emailer,
		AccountResolver:               account.NewResolver(repositoryService),
		AuthenticationResolver:        authentication.NewResolver(repositoryService),
		BusinessDetailRepository:      businessdetail.NewRepository(repositoryService),
		ChannelRequestRepository:      channelrequest.NewRepository(repositoryService),
		ConnectorRepository:           connector.NewRepository(repositoryService),
		CredentialRepository:          credential.NewRepository(repositoryService),
		DisplayTextRepository:         displaytext.NewRepository(repositoryService),
		EmailSubscriptionRepository:   emailsubscription.NewRepository(repositoryService),
		EnergyMixRepository:           energymix.NewRepository(repositoryService),
		EvseRepository:                evse.NewRepository(repositoryService),
		ImageRepository:               image.NewRepository(repositoryService),
		InvoiceRequestRepository:      invoicerequest.NewRepository(repositoryService),
		LocationRepository:            location.NewRepository(repositoryService),
		NodeRepository:                node.NewRepository(repositoryService),
		OpeningTimeRepository:         openingtime.NewRepository(repositoryService),
		PartyRepository:               party.NewRepository(repositoryService),
		PendingNotificationRepository: pendingnotification.NewRepository(repositoryService),
		PromotionRepository:           promotion.NewRepository(repositoryService),
		PoiRepository:                 poi.NewRepository(repositoryService),
		ReferralRepository:            referral.NewRepository(repositoryService),
		SessionRepository:             session.NewRepository(repositoryService),
		TariffRepository:              tariff.NewRepository(repositoryService),
		TokenAuthorizationRepository:  tokenauthorization.NewRepository(repositoryService),
		TokenResolver:                 token.NewResolver(repositoryService),
		UserRepository:                user.NewRepository(repositoryService),
		defaultTaxPercent:             defaultTaxPercent,
	}
}

func (r *Resolver) calculateTaxPercent(ctx context.Context) (*float64, error) {
	operationCtx := graphql.GetOperationContext(ctx)
	input := operationCtx.Variables["input"]

	if input != nil {
		inputVariables := input.(map[string]interface{})

		if country, ok := inputVariables["country"]; ok {
			taxPercent := r.AccountResolver.GetTaxPercentByCountry(ctx, country.(string), r.defaultTaxPercent)

			return &taxPercent, nil
		}
	}

	return nil, gqlerror.Errorf("Error retrieving tax by country")
}
