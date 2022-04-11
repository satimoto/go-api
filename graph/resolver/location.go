package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/authentication"
	"github.com/satimoto/go-api/internal/location"
	"github.com/satimoto/go-api/internal/util"
	"github.com/satimoto/go-datastore/db"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *queryResolver) GetLocation(ctx context.Context, uid string) (*db.Location, error) {
	if userId := authentication.GetUserId(ctx); userId != nil {
		if l, err := r.LocationResolver.Repository.GetLocationByUid(ctx, uid); err == nil {
			return &l, nil
		}
		return nil, gqlerror.Errorf("Location not found")
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

func (r *queryResolver) ListLocations(ctx context.Context, input graph.ListLocationsInput) ([]graph.ListLocation, error) {
	if userId := authentication.GetUserId(ctx); userId != nil {
		var list []graph.ListLocation

		params := location.NewListLocationsByGeomParams(input)

		if locations, err := r.LocationResolver.Repository.ListLocationsByGeom(ctx, params); err == nil {
			for _, l := range locations {
				list = append(list, location.NewListLocation(l))
			}
		}

		return list, nil
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

func (r *locationResolver) Type(ctx context.Context, obj *db.Location) (string, error) {
	return string(obj.Type), nil
}

func (r *locationResolver) Name(ctx context.Context, obj *db.Location) (*string, error) {
	return util.NullString(obj.Name)
}

func (r *locationResolver) RelatedLocations(ctx context.Context, obj *db.Location) ([]graph.Geolocation, error) {
	list := []graph.Geolocation{}

	if relatedLocations, err := r.LocationResolver.Repository.ListRelatedLocations(ctx, obj.ID); err == nil {
		for _, relatedLocation := range relatedLocations {
			name, _ := util.NullString(relatedLocation.Name)
			list = append(list, graph.Geolocation{
				Latitude:  relatedLocation.LatitudeFloat,
				Longitude: relatedLocation.LongitudeFloat,
				Name:      name,
			})
		}
	}

	return list, nil
}

func (r *locationResolver) Evses(ctx context.Context, obj *db.Location) ([]db.Evse, error) {
	return r.LocationResolver.Repository.ListEvses(ctx, obj.ID)
}

func (r *locationResolver) Directions(ctx context.Context, obj *db.Location) ([]db.DisplayText, error) {
	return r.LocationResolver.Repository.ListLocationDirections(ctx, obj.ID)
}

func (r *locationResolver) Operator(ctx context.Context, obj *db.Location) (*db.BusinessDetail, error) {
	if obj.OperatorID.Valid {
		if businessDetail, err := r.BusinessDetailResolver.Repository.GetBusinessDetail(ctx, obj.OperatorID.Int64); err == nil {
			return &businessDetail, nil
		}
	}

	return nil, nil
}

func (r *locationResolver) Suboperator(ctx context.Context, obj *db.Location) (*db.BusinessDetail, error) {
	if obj.SuboperatorID.Valid {
		if businessDetail, err := r.BusinessDetailResolver.Repository.GetBusinessDetail(ctx, obj.SuboperatorID.Int64); err == nil {
			return &businessDetail, nil
		}
	}

	return nil, nil
}

func (r *locationResolver) Owner(ctx context.Context, obj *db.Location) (*db.BusinessDetail, error) {
	if obj.OwnerID.Valid {
		if businessDetail, err := r.BusinessDetailResolver.Repository.GetBusinessDetail(ctx, obj.OwnerID.Int64); err == nil {
			return &businessDetail, nil
		}
	}

	return nil, nil
}

func (r *locationResolver) Facilities(ctx context.Context, obj *db.Location) ([]graph.TextDescription, error) {
	list := []graph.TextDescription{}

	if facilities, err := r.LocationResolver.Repository.ListLocationFacilities(ctx, obj.ID); err == nil {
		for _, facility := range facilities {
			list = append(list, graph.TextDescription{
				Text:        facility.Text,
				Description: facility.Description,
			})
		}
	}

	return list, nil
}

func (r *locationResolver) TimeZone(ctx context.Context, obj *db.Location) (*string, error) {
	return util.NullString(obj.TimeZone)
}

func (r *locationResolver) OpeningTime(ctx context.Context, obj *db.Location) (*db.OpeningTime, error) {
	if obj.OpeningTimeID.Valid {
		if openingTime, err := r.OpeningTimeResolver.Repository.GetOpeningTime(ctx, obj.OpeningTimeID.Int64); err == nil {
			return &openingTime, nil
		}
	}

	return nil, nil
}

func (r *locationResolver) Images(ctx context.Context, obj *db.Location) ([]db.Image, error) {
	return r.LocationResolver.Repository.ListLocationImages(ctx, obj.ID)
}

func (r *locationResolver) EnergyMix(ctx context.Context, obj *db.Location) (*db.EnergyMix, error) {
	if obj.EnergyMixID.Valid {
		if energyMix, err := r.EnergyMixResolver.Repository.GetEnergyMix(ctx, obj.EnergyMixID.Int64); err == nil {
			return &energyMix, nil
		}
	}

	return nil, nil
}

func (r *locationResolver) LastUpdated(ctx context.Context, obj *db.Location) (string, error) {
	return obj.LastUpdated.Format(time.RFC3339Nano), nil
}

// Location returns graph.LocationResolver implementation.
func (r *Resolver) Location() graph.LocationResolver { return &locationResolver{r} }

type locationResolver struct{ *Resolver }
