package businessdetail

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

type BusinessDetailRepository interface {
	CreateBusinessDetail(ctx context.Context, arg db.CreateBusinessDetailParams) (db.BusinessDetail, error)
	GetBusinessDetail(ctx context.Context, id int64) (db.BusinessDetail, error)
}

type BusinessDetailResolver struct {
	Repository BusinessDetailRepository
}

func NewResolver(repositoryService *db.RepositoryService) *BusinessDetailResolver {
	repo := BusinessDetailRepository(repositoryService)
	return &BusinessDetailResolver{repo}
}
