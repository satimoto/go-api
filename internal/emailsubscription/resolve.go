package emailsubscription

import (
	"context"
	"os"

	"github.com/satimoto/go-api/internal/aws/email"
	"github.com/satimoto/go-datastore/pkg/db"
)

type EmailSubscriptionRepository interface {
	CreateEmailSubscription(ctx context.Context, arg db.CreateEmailSubscriptionParams) (db.EmailSubscription, error)
	GetEmailSubscriptionByEmail(ctx context.Context, email string) (db.EmailSubscription, error)
	UpdateEmailSubscription(ctx context.Context, arg db.UpdateEmailSubscriptionParams) (db.EmailSubscription, error)
}

type EmailSubscriptionResolver struct {
	Repository EmailSubscriptionRepository
	Emailer    email.Emailer
}

func NewResolver(repositoryService *db.RepositoryService) *EmailSubscriptionResolver {
	repo := EmailSubscriptionRepository(repositoryService)
	emailer := email.New(os.Getenv("REPLY_TO_EMAIL"))

	return &EmailSubscriptionResolver{
		Repository: repo,
		Emailer:    emailer,
	}
}
