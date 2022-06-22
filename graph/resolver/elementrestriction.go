package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
)

func (r *elementRestrictionResolver) StartTime(ctx context.Context, obj *db.ElementRestriction) (*string, error) {
	return util.NilString(obj.StartTime), nil
}

func (r *elementRestrictionResolver) EndTime(ctx context.Context, obj *db.ElementRestriction) (*string, error) {
	return util.NilString(obj.EndTime), nil
}

func (r *elementRestrictionResolver) StartDate(ctx context.Context, obj *db.ElementRestriction) (*string, error) {
	return util.NilString(obj.StartDate), nil
}

func (r *elementRestrictionResolver) EndDate(ctx context.Context, obj *db.ElementRestriction) (*string, error) {
	return util.NilString(obj.EndDate), nil
}

func (r *elementRestrictionResolver) MinKwh(ctx context.Context, obj *db.ElementRestriction) (*float64, error) {
	return util.NilFloat64(obj.MinKwh), nil
}

func (r *elementRestrictionResolver) MaxKwh(ctx context.Context, obj *db.ElementRestriction) (*float64, error) {
	return util.NilFloat64(obj.MaxKwh), nil
}

func (r *elementRestrictionResolver) MinPower(ctx context.Context, obj *db.ElementRestriction) (*float64, error) {
	return util.NilFloat64(obj.MinPower), nil
}

func (r *elementRestrictionResolver) MaxPower(ctx context.Context, obj *db.ElementRestriction) (*float64, error) {
	return util.NilFloat64(obj.MaxPower), nil
}

func (r *elementRestrictionResolver) MinDuration(ctx context.Context, obj *db.ElementRestriction) (*int, error) {
	return util.NilInt(obj.MinDuration), nil
}

func (r *elementRestrictionResolver) MaxDuration(ctx context.Context, obj *db.ElementRestriction) (*int, error) {
	return util.NilInt(obj.MaxDuration), nil
}

func (r *elementRestrictionResolver) DayOfWeek(ctx context.Context, obj *db.ElementRestriction) ([]string, error) {
	list := []string{}

	if weekdays, err := r.TariffRepository.ListElementRestrictionWeekdays(ctx, obj.ID); err == nil {
		for _, weekday := range weekdays {
			list = append(list, weekday.Text)
		}
	}

	return list, nil
}

// ElementRestriction returns graph.ElementRestrictionResolver implementation.
func (r *Resolver) ElementRestriction() graph.ElementRestrictionResolver {
	return &elementRestrictionResolver{r}
}

type elementRestrictionResolver struct{ *Resolver }
