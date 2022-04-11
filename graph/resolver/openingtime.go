package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-datastore/db"
)

func (r *openingTimeResolver) RegularHours(ctx context.Context, obj *db.OpeningTime) ([]db.RegularHour, error) {
	return r.OpeningTimeResolver.Repository.ListRegularHours(ctx, obj.ID)
}

func (r *openingTimeResolver) ExceptionalOpenings(ctx context.Context, obj *db.OpeningTime) ([]db.ExceptionalPeriod, error) {
	return r.OpeningTimeResolver.Repository.ListExceptionalOpeningPeriods(ctx, obj.ID)
}

func (r *openingTimeResolver) ExceptionalClosings(ctx context.Context, obj *db.OpeningTime) ([]db.ExceptionalPeriod, error) {
	return r.OpeningTimeResolver.Repository.ListExceptionalClosingPeriods(ctx, obj.ID)
}

// OpeningTime returns graph.OpeningTimeResolver implementation.
func (r *Resolver) OpeningTime() graph.OpeningTimeResolver { return &openingTimeResolver{r} }

type openingTimeResolver struct{ *Resolver }
