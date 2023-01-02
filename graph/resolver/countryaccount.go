package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/middleware"
	"github.com/satimoto/go-api/internal/util"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Longitude is the resolver for the longitude field.
func (r *countryAccountResolver) Longitude(ctx context.Context, obj *db.CountryAccount) (*float64, error) {
	return util.NullFloat(obj.Longitude)
}

// Latitude is the resolver for the latitude field.
func (r *countryAccountResolver) Latitude(ctx context.Context, obj *db.CountryAccount) (*float64, error) {
	return util.NullFloat(obj.Latitude)
}

// Zoom is the resolver for the zoom field.
func (r *countryAccountResolver) Zoom(ctx context.Context, obj *db.CountryAccount) (*float64, error) {
	return util.NullFloat(obj.Zoom)
}

// ListCountryAccounts is the resolver for the listCountryAccounts field.
func (r *queryResolver) ListCountryAccounts(ctx context.Context) ([]db.CountryAccount, error) {
	if userID := middleware.GetUserID(ctx); userID != nil {
		return r.CountryAccountResolver.Repository.ListCountryAccounts(ctx)
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

// CountryAccount returns graph.CountryAccountResolver implementation.
func (r *Resolver) CountryAccount() graph.CountryAccountResolver { return &countryAccountResolver{r} }

type countryAccountResolver struct{ *Resolver }
