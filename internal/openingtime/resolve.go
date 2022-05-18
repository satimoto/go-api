package openingtime

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

type OpeningTimeRepository interface {
	GetOpeningTime(ctx context.Context, id int64) (db.OpeningTime, error)
	ListExceptionalOpeningPeriods(ctx context.Context, openingTimeID int64) ([]db.ExceptionalPeriod, error)
	ListExceptionalClosingPeriods(ctx context.Context, openingTimeID int64) ([]db.ExceptionalPeriod, error)
	ListRegularHours(ctx context.Context, openingTimeID int64) ([]db.RegularHour, error)
}

type OpeningTimeResolver struct {
	Repository OpeningTimeRepository
}

func NewResolver(repositoryService *db.RepositoryService) *OpeningTimeResolver {
	repo := OpeningTimeRepository(repositoryService)
	return &OpeningTimeResolver{repo}
}
