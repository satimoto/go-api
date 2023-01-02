package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
)

// StartTime is the resolver for the startTime field.
func (r *elementRestrictionResolver) StartTime(ctx context.Context, obj *db.ElementRestriction) (*string, error) {
	return util.NilString(obj.StartTime), nil
}

// EndTime is the resolver for the endTime field.
func (r *elementRestrictionResolver) EndTime(ctx context.Context, obj *db.ElementRestriction) (*string, error) {
	return util.NilString(obj.EndTime), nil
}

// StartDate is the resolver for the startDate field.
func (r *elementRestrictionResolver) StartDate(ctx context.Context, obj *db.ElementRestriction) (*string, error) {
	return util.NilString(obj.StartDate), nil
}

// EndDate is the resolver for the endDate field.
func (r *elementRestrictionResolver) EndDate(ctx context.Context, obj *db.ElementRestriction) (*string, error) {
	return util.NilString(obj.EndDate), nil
}

// MinKwh is the resolver for the minKwh field.
func (r *elementRestrictionResolver) MinKwh(ctx context.Context, obj *db.ElementRestriction) (*float64, error) {
	return util.NilFloat64(obj.MinKwh), nil
}

// MaxKwh is the resolver for the maxKwh field.
func (r *elementRestrictionResolver) MaxKwh(ctx context.Context, obj *db.ElementRestriction) (*float64, error) {
	return util.NilFloat64(obj.MaxKwh), nil
}

// MinPower is the resolver for the minPower field.
func (r *elementRestrictionResolver) MinPower(ctx context.Context, obj *db.ElementRestriction) (*float64, error) {
	return util.NilFloat64(obj.MinPower), nil
}

// MaxPower is the resolver for the maxPower field.
func (r *elementRestrictionResolver) MaxPower(ctx context.Context, obj *db.ElementRestriction) (*float64, error) {
	return util.NilFloat64(obj.MaxPower), nil
}

// MinDuration is the resolver for the minDuration field.
func (r *elementRestrictionResolver) MinDuration(ctx context.Context, obj *db.ElementRestriction) (*int, error) {
	return util.NilInt(obj.MinDuration), nil
}

// MaxDuration is the resolver for the maxDuration field.
func (r *elementRestrictionResolver) MaxDuration(ctx context.Context, obj *db.ElementRestriction) (*int, error) {
	return util.NilInt(obj.MaxDuration), nil
}

// DayOfWeek is the resolver for the dayOfWeek field.
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
