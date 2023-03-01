package account

import (
	"github.com/satimoto/go-datastore/pkg/account"
	"github.com/satimoto/go-datastore/pkg/db"
)

type AccountResolver struct {
	Repository account.AccountRepository
}

func NewResolver(repositoryService *db.RepositoryService) *AccountResolver {
	return &AccountResolver{
		Repository: account.NewRepository(repositoryService),
	}
}
