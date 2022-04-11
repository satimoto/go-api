package credential

import (
	"context"
	"os"

	"github.com/satimoto/go-api/internal/aws/email"
	"github.com/satimoto/go-datastore/db"
)

type CredentialRepository interface {
	CreateCredential(ctx context.Context, arg db.CreateCredentialParams) (db.Credential, error)
}

type CredentialResolver struct {
	Repository CredentialRepository
	Emailer    email.Emailer
}

func NewResolver(repositoryService *db.RepositoryService) *CredentialResolver {
	repo := CredentialRepository(repositoryService)
	emailer := email.New(os.Getenv("REPLY_TO_EMAIL"))

	return &CredentialResolver{
		Repository: repo,
		Emailer:    emailer,
	}
}
