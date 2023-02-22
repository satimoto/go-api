package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	accountMocks "github.com/satimoto/go-datastore/pkg/account/mocks"
	"github.com/satimoto/go-api/internal/account"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *account.AccountResolver {
	return &account.AccountResolver{
		Repository: accountMocks.NewRepository(repositoryService),
	}
}
