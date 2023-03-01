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
	dbUtil "github.com/satimoto/go-datastore/pkg/util"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Type is the resolver for the type field.
func (r *locationResolver) Type(ctx context.Context, obj *db.Location) (string, error) {
	return string(obj.Type), nil
}

// Name is the resolver for the name field.
func (r *locationResolver) Name(ctx context.Context, obj *db.Location) (*string, error) {
	return util.NullString(obj.Name)
}

// RelatedLocations is the resolver for the relatedLocations field.
func (r *locationResolver) RelatedLocations(ctx context.Context, obj *db.Location) ([]graph.AddtionalGeoLocation, error) {
	list := []graph.AddtionalGeoLocation{}

	if additionalGeoLocations, err := r.LocationRepository.ListAdditionalGeoLocations(ctx, obj.ID); err == nil {
		for _, additionalGeoLocation := range additionalGeoLocations {
			agl := graph.AddtionalGeoLocation{
				Latitude:  additionalGeoLocation.LatitudeFloat,
				Longitude: additionalGeoLocation.LongitudeFloat,
			}

			if additionalGeoLocation.DisplayTextID.Valid {
				if displayText, err := r.DisplayTextRepository.GetDisplayText(ctx, additionalGeoLocation.DisplayTextID.Int64); err == nil {
					agl.Name = &displayText
				}
			}

			list = append(list, agl)
		}
	}

	return list, nil
}

// Evses is the resolver for the evses field.
func (r *locationResolver) Evses(ctx context.Context, obj *db.Location) ([]db.Evse, error) {
	return r.LocationRepository.ListActiveEvses(ctx, obj.ID)
}

// IsExperimental is the resolver for the isExperimental field.
func (r *locationResolver) IsExperimental(ctx context.Context, obj *db.Location) (bool, error) {
	return !obj.IsIntermediateCdrCapable, nil
}

// Directions is the resolver for the directions field.
func (r *locationResolver) Directions(ctx context.Context, obj *db.Location) ([]db.DisplayText, error) {
	return r.LocationRepository.ListLocationDirections(ctx, obj.ID)
}

// Operator is the resolver for the operator field.
func (r *locationResolver) Operator(ctx context.Context, obj *db.Location) (*db.BusinessDetail, error) {
	if obj.OperatorID.Valid {
		if businessDetail, err := r.BusinessDetailRepository.GetBusinessDetail(ctx, obj.OperatorID.Int64); err == nil {
			return &businessDetail, nil
		}
	}

	return nil, nil
}

// Suboperator is the resolver for the suboperator field.
func (r *locationResolver) Suboperator(ctx context.Context, obj *db.Location) (*db.BusinessDetail, error) {
	if obj.SuboperatorID.Valid {
		if businessDetail, err := r.BusinessDetailRepository.GetBusinessDetail(ctx, obj.SuboperatorID.Int64); err == nil {
			return &businessDetail, nil
		}
	}

	return nil, nil
}

// Owner is the resolver for the owner field.
func (r *locationResolver) Owner(ctx context.Context, obj *db.Location) (*db.BusinessDetail, error) {
	if obj.OwnerID.Valid {
		if businessDetail, err := r.BusinessDetailRepository.GetBusinessDetail(ctx, obj.OwnerID.Int64); err == nil {
			return &businessDetail, nil
		}
	}

	return nil, nil
}

// Facilities is the resolver for the facilities field.
func (r *locationResolver) Facilities(ctx context.Context, obj *db.Location) ([]graph.TextDescription, error) {
	list := []graph.TextDescription{}

	if facilities, err := r.LocationRepository.ListLocationFacilities(ctx, obj.ID); err == nil {
		for _, facility := range facilities {
			list = append(list, graph.TextDescription{
				Text:        facility.Text,
				Description: facility.Description,
			})
		}
	}

	return list, nil
}

// TimeZone is the resolver for the timeZone field.
func (r *locationResolver) TimeZone(ctx context.Context, obj *db.Location) (*string, error) {
	return util.NullString(obj.TimeZone)
}

// OpeningTime is the resolver for the openingTime field.
func (r *locationResolver) OpeningTime(ctx context.Context, obj *db.Location) (*db.OpeningTime, error) {
	if obj.OpeningTimeID.Valid {
		if openingTime, err := r.OpeningTimeRepository.GetOpeningTime(ctx, obj.OpeningTimeID.Int64); err == nil {
			return &openingTime, nil
		}
	}

	return nil, nil
}

// Images is the resolver for the images field.
func (r *locationResolver) Images(ctx context.Context, obj *db.Location) ([]db.Image, error) {
	return r.LocationRepository.ListLocationImages(ctx, obj.ID)
}

// EnergyMix is the resolver for the energyMix field.
func (r *locationResolver) EnergyMix(ctx context.Context, obj *db.Location) (*db.EnergyMix, error) {
	if obj.EnergyMixID.Valid {
		if energyMix, err := r.EnergyMixRepository.GetEnergyMix(ctx, obj.EnergyMixID.Int64); err == nil {
			return &energyMix, nil
		}
	}

	return nil, nil
}

// LastUpdated is the resolver for the lastUpdated field.
func (r *locationResolver) LastUpdated(ctx context.Context, obj *db.Location) (string, error) {
	return obj.LastUpdated.Format(time.RFC3339), nil
}

// PublishLocation is the resolver for the publishLocation field.
func (r *mutationResolver) PublishLocation(ctx context.Context, input graph.PublishLocationInput) (*graph.ResultOk, error) {
	backgroundCtx := context.Background()

	if user := middleware.GetCtxUser(ctx, r.UserRepository); user != nil && user.IsAdmin {
		if input.ID != nil {
			updateLocationPublishedParams := db.UpdateLocationPublishedParams{
				ID:          *input.ID,
				IsPublished: input.IsPublished,
			}

			if err := r.LocationRepository.UpdateLocationPublished(backgroundCtx, updateLocationPublishedParams); err == nil {
				return &graph.ResultOk{Ok: true}, nil
			}
		} else if input.CredentialID != nil {
			updateLocationsPublishedByCredentialParams := db.UpdateLocationsPublishedByCredentialParams{
				CredentialID: *input.CredentialID,
				IsPublished:  input.IsPublished,
			}

			if err := r.LocationRepository.UpdateLocationsPublishedByCredential(backgroundCtx, updateLocationsPublishedByCredentialParams); err == nil {
				return &graph.ResultOk{Ok: true}, nil
			}
		} else if input.PartyID != nil && input.CountryCode != nil {
			updateLocationsPublishedByPartyAndCountryCodeParams := db.UpdateLocationsPublishedByPartyAndCountryCodeParams{
				CountryCode: dbUtil.SqlNullString(input.CountryCode),
				PartyID:     dbUtil.SqlNullString(input.PartyID),
				IsPublished: input.IsPublished,
			}

			if err := r.LocationRepository.UpdateLocationsPublishedByPartyAndCountryCode(backgroundCtx, updateLocationsPublishedByPartyAndCountryCodeParams); err == nil {
				return &graph.ResultOk{Ok: true}, nil
			}
		}
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

// GetLocation is the resolver for the getLocation field.
func (r *queryResolver) GetLocation(ctx context.Context, input graph.GetLocationInput) (*db.Location, error) {
	backgroundCtx := context.Background()

	if userID := middleware.GetUserID(ctx); userID != nil {
		if input.ID != nil {
			if l, err := r.LocationRepository.GetLocation(backgroundCtx, *input.ID); err == nil {
				return &l, nil
			}
		} else if input.UID != nil {
			if l, err := r.LocationRepository.GetLocationByUid(backgroundCtx, *input.UID); err == nil {
				return &l, nil
			}
		}

		return nil, gqlerror.Errorf("Location not found")
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

// ListLocations is the resolver for the listLocations field.
func (r *queryResolver) ListLocations(ctx context.Context, input graph.ListLocationsInput) ([]graph.ListLocation, error) {
	backgroundCtx := context.Background()

	if userID := middleware.GetUserID(ctx); userID != nil {
		var list []graph.ListLocation
		var locations []db.Location
		var err error

		if input.Country != nil && len(*input.Country) > 0 {
			locations, err = r.LocationRepository.ListLocationsByCountry(backgroundCtx, *input.Country)
		} else if input.XMin != nil && input.XMax != nil && input.YMin != nil && input.YMax != nil {
			params := param.NewListLocationsByGeomParams(input)
			locations, err = r.LocationRepository.ListLocationsByGeom(backgroundCtx, params)
		} else if user := middleware.GetCtxUser(ctx, r.UserRepository); user.IsAdmin {
			locations, err = r.LocationRepository.ListLocations(backgroundCtx)
		}

		if err == nil {
			for _, l := range locations {
				list = append(list, param.NewListLocation(l))
			}
		}

		return list, nil
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

// Location returns graph.LocationResolver implementation.
func (r *Resolver) Location() graph.LocationResolver { return &locationResolver{r} }

type locationResolver struct{ *Resolver }
