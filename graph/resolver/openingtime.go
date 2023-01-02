package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-datastore/pkg/db"
)

// RegularHours is the resolver for the regularHours field.
func (r *openingTimeResolver) RegularHours(ctx context.Context, obj *db.OpeningTime) ([]db.RegularHour, error) {
	return r.OpeningTimeRepository.ListRegularHours(ctx, obj.ID)
}

// ExceptionalOpenings is the resolver for the exceptionalOpenings field.
func (r *openingTimeResolver) ExceptionalOpenings(ctx context.Context, obj *db.OpeningTime) ([]db.ExceptionalPeriod, error) {
	return r.OpeningTimeRepository.ListExceptionalOpeningPeriods(ctx, obj.ID)
}

// ExceptionalClosings is the resolver for the exceptionalClosings field.
func (r *openingTimeResolver) ExceptionalClosings(ctx context.Context, obj *db.OpeningTime) ([]db.ExceptionalPeriod, error) {
	return r.OpeningTimeRepository.ListExceptionalClosingPeriods(ctx, obj.ID)
}

// OpeningTime returns graph.OpeningTimeResolver implementation.
func (r *Resolver) OpeningTime() graph.OpeningTimeResolver { return &openingTimeResolver{r} }

type openingTimeResolver struct{ *Resolver }
