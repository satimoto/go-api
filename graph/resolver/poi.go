package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/middleware"
	"github.com/satimoto/go-api/internal/param"
	"github.com/satimoto/go-api/internal/util"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Description is the resolver for the description field.
func (r *poiResolver) Description(ctx context.Context, obj *db.Poi) (*string, error) {
	return util.NullString(obj.Description)
}

// Address is the resolver for the address field.
func (r *poiResolver) Address(ctx context.Context, obj *db.Poi) (*string, error) {
	return util.NullString(obj.Address)
}

// City is the resolver for the city field.
func (r *poiResolver) City(ctx context.Context, obj *db.Poi) (*string, error) {
	return util.NullString(obj.City)
}

// PostalCode is the resolver for the postalCode field.
func (r *poiResolver) PostalCode(ctx context.Context, obj *db.Poi) (*string, error) {
	return util.NullString(obj.PostalCode)
}

// Tags is the resolver for the tags field.
func (r *poiResolver) Tags(ctx context.Context, obj *db.Poi) ([]db.Tag, error) {
	return r.PoiRepository.ListPoiTags(ctx, obj.ID)
}

// PaymentURI is the resolver for the paymentUri field.
func (r *poiResolver) PaymentURI(ctx context.Context, obj *db.Poi) (*string, error) {
	return util.NullString(obj.PaymentUri)
}

// OpeningTimes is the resolver for the opening_times field.
func (r *poiResolver) OpeningTimes(ctx context.Context, obj *db.Poi) (*string, error) {
	return util.NullString(obj.OpeningTimes)
}

// Phone is the resolver for the phone field.
func (r *poiResolver) Phone(ctx context.Context, obj *db.Poi) (*string, error) {
	return util.NullString(obj.Phone)
}

// Website is the resolver for the website field.
func (r *poiResolver) Website(ctx context.Context, obj *db.Poi) (*string, error) {
	return util.NullString(obj.Website)
}

// LastUpdated is the resolver for the lastUpdated field.
func (r *poiResolver) LastUpdated(ctx context.Context, obj *db.Poi) (string, error) {
	return obj.LastUpdated.Format(time.RFC3339), nil
}

// GetPoi is the resolver for the getPoi field.
func (r *queryResolver) GetPoi(ctx context.Context, input graph.GetPoiInput) (*db.Poi, error) {
	backgroundCtx := context.Background()

	if userID := middleware.GetUserID(ctx); userID != nil {
		if input.UID != nil {
			if l, err := r.PoiRepository.GetPoiByUid(backgroundCtx, *input.UID); err == nil {
				return &l, nil
			}
		}

		return nil, gqlerror.Errorf("Poi not found")
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

// ListPois is the resolver for the listPois field.
func (r *queryResolver) ListPois(ctx context.Context, input graph.ListPoisInput) ([]db.Poi, error) {
	backgroundCtx := context.Background()

	if userID := middleware.GetUserID(ctx); userID != nil {
		if input.XMin != nil && input.XMax != nil && input.YMin != nil && input.YMax != nil {
			params := param.NewListPoisByGeomParams(input)

			return r.PoiRepository.ListPoisByGeom(backgroundCtx, params)
		}
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

// Poi returns graph.PoiResolver implementation.
func (r *Resolver) Poi() graph.PoiResolver { return &poiResolver{r} }

type poiResolver struct{ *Resolver }
