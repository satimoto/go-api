package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-datastore/pkg/db"
)

func (r *regularHourResolver) Weekday(ctx context.Context, obj *db.RegularHour) (int, error) {
	return int(obj.Weekday), nil
}

// RegularHour returns graph.RegularHourResolver implementation.
func (r *Resolver) RegularHour() graph.RegularHourResolver { return &regularHourResolver{r} }

type regularHourResolver struct{ *Resolver }
