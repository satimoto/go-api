package emailsubscription

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type EmailSubscriptionRepository interface {
	CreateEmailSubscription(ctx context.Context, arg db.CreateEmailSubscriptionParams) (db.EmailSubscription, error)
	GetEmailSubscriptionByEmail(ctx context.Context, email string) (db.EmailSubscription, error)
	UpdateEmailSubscription(ctx context.Context, arg db.UpdateEmailSubscriptionParams) (db.EmailSubscription, error)
}

type EmailSubscriptionResolver struct {
	Repository EmailSubscriptionRepository
}

func NewResolver(repositoryService *db.RepositoryService) *EmailSubscriptionResolver {
	repo := EmailSubscriptionRepository(repositoryService)
	return &EmailSubscriptionResolver{repo}
}
