package countryaccount

import (
	"github.com/satimoto/go-datastore/pkg/countryaccount"
	"github.com/satimoto/go-datastore/pkg/db"
)

type CountryAccountResolver struct {
	Repository countryaccount.CountryAccountRepository
}

func NewResolver(repositoryService *db.RepositoryService) *CountryAccountResolver {
	return &CountryAccountResolver{
		Repository: countryaccount.NewRepository(repositoryService),
	}
}
