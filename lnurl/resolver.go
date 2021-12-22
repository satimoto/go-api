package lnurl

import "github.com/satimoto/go-datastore/db"

type LnUrlRepository interface {
}

type LnUrlResolver struct {
	Repository LnUrlRepository
}

func NewResolver(repositoryService *db.RepositoryService) *LnUrlResolver {
	repo := LnUrlRepository(repositoryService)

	return &LnUrlResolver{
		Repository: repo,
	}
}
