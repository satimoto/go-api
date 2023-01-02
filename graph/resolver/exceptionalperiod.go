package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-datastore/pkg/db"
)

// PeriodBegin is the resolver for the periodBegin field.
func (r *exceptionalPeriodResolver) PeriodBegin(ctx context.Context, obj *db.ExceptionalPeriod) (string, error) {
	return obj.PeriodBegin.Format(time.RFC3339), nil
}

// PeriodEnd is the resolver for the periodEnd field.
func (r *exceptionalPeriodResolver) PeriodEnd(ctx context.Context, obj *db.ExceptionalPeriod) (string, error) {
	return obj.PeriodEnd.Format(time.RFC3339), nil
}

// ExceptionalPeriod returns graph.ExceptionalPeriodResolver implementation.
func (r *Resolver) ExceptionalPeriod() graph.ExceptionalPeriodResolver {
	return &exceptionalPeriodResolver{r}
}

type exceptionalPeriodResolver struct{ *Resolver }
