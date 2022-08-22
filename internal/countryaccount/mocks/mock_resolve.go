package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	countryaccountMocks "github.com/satimoto/go-datastore/pkg/countryaccount/mocks"
	"github.com/satimoto/go-api/internal/countryaccount"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *countryaccount.CountryAccountResolver {
	return &countryaccount.CountryAccountResolver{
		Repository: countryaccountMocks.NewRepository(repositoryService),
	}
}
