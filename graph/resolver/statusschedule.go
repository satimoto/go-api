package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/util"
	"github.com/satimoto/go-datastore/pkg/db"
)

func (r *statusScheduleResolver) PeriodBegin(ctx context.Context, obj *db.StatusSchedule) (string, error) {
	return obj.PeriodBegin.Format(time.RFC3339Nano), nil
}

func (r *statusScheduleResolver) PeriodEnd(ctx context.Context, obj *db.StatusSchedule) (*string, error) {
	return util.NullTime(obj.PeriodEnd, time.RFC3339Nano)
}
func (r *statusScheduleResolver) Status(ctx context.Context, obj *db.StatusSchedule) (string, error) {
	return string(obj.Status), nil
}

// StatusSchedule returns graph.StatusScheduleResolver implementation.
func (r *Resolver) StatusSchedule() graph.StatusScheduleResolver { return &statusScheduleResolver{r} }

type statusScheduleResolver struct{ *Resolver }
