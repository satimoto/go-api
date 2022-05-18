package authentication

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

type AuthenticationRepository interface {
	CreateAuthentication(ctx context.Context, arg db.CreateAuthenticationParams) (db.Authentication, error)
	DeleteAuthentication(ctx context.Context, id int64) error
	GetAuthenticationByChallenge(ctx context.Context, challenge string) (db.Authentication, error)
	GetAuthenticationByCode(ctx context.Context, code string) (db.Authentication, error)
	UpdateAuthentication(ctx context.Context, arg db.UpdateAuthenticationParams) (db.Authentication, error)
}
type AuthenticationResolver struct {
	Repository AuthenticationRepository
}

func NewResolver(repositoryService *db.RepositoryService) *AuthenticationResolver {
	repo := AuthenticationRepository(repositoryService)
	return &AuthenticationResolver{repo}
}
