package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/authentication"
	"github.com/satimoto/go-api/internal/param"
	"github.com/satimoto/go-api/internal/util"
	"github.com/satimoto/go-datastore/pkg/db"
	dbUtil "github.com/satimoto/go-datastore/pkg/util"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *queryResolver) GetLocation(ctx context.Context, input graph.GetLocationInput) (*db.Location, error) {
	if userId := authentication.GetUserId(ctx); userId != nil {
		if input.ID != nil {
			if l, err := r.LocationRepository.GetLocation(ctx, *input.ID); err == nil {
				return &l, nil
			}
		} else if input.UID != nil {
			if l, err := r.LocationRepository.GetLocationByUid(ctx, *input.UID); err == nil {
				return &l, nil
			}
		}

		return nil, gqlerror.Errorf("Location not found")
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

func (r *queryResolver) ListLocations(ctx context.Context, input graph.ListLocationsInput) ([]graph.ListLocation, error) {
	if userId := authentication.GetUserId(ctx); userId != nil {
		var list []graph.ListLocation

		params := param.NewListLocationsByGeomParams(input)

		if locations, err := r.LocationRepository.ListLocationsByGeom(ctx, params); err == nil {
			for _, l := range locations {
				list = append(list, param.NewListLocation(l))
			}
		}

		return list, nil
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

func (r *mutationResolver) PublishLocation(ctx context.Context, input graph.PublishLocationInput) (*graph.ResultOk, error) {
	if userId := authentication.GetUserId(ctx); userId != nil {
		if input.ID != nil {
			updateLocationPublishParams := db.UpdateLocationPublishParams{
				ID:      *input.ID,
				Publish: input.Publish,
			}

			if err := r.LocationRepository.UpdateLocationPublish(ctx, updateLocationPublishParams); err == nil {
				return &graph.ResultOk{Ok: true}, nil
			}
		} else if input.CredentialID != nil {
			updateLocationsPublishByCredentialParams := db.UpdateLocationsPublishByCredentialParams{
				CredentialID: *input.CredentialID,
				Publish:      input.Publish,
			}

			if err := r.LocationRepository.UpdateLocationsPublishByCredential(ctx, updateLocationsPublishByCredentialParams); err == nil {
				return &graph.ResultOk{Ok: true}, nil
			}
		} else if input.PartyID != nil && input.CountryCode != nil {
			updateLocationsPublishByPartyAndCountryCodeParams := db.UpdateLocationsPublishByPartyAndCountryCodeParams{
				CountryCode: dbUtil.SqlNullString(input.CountryCode),
				PartyID:     dbUtil.SqlNullString(input.PartyID),
				Publish:     input.Publish,
			}

			if err := r.LocationRepository.UpdateLocationsPublishByPartyAndCountryCode(ctx, updateLocationsPublishByPartyAndCountryCodeParams); err == nil {
				return &graph.ResultOk{Ok: true}, nil
			}
		}
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

func (r *locationResolver) Type(ctx context.Context, obj *db.Location) (string, error) {
	return string(obj.Type), nil
}

func (r *locationResolver) Name(ctx context.Context, obj *db.Location) (*string, error) {
	return util.NullString(obj.Name)
}

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

func (r *locationResolver) Evses(ctx context.Context, obj *db.Location) ([]db.Evse, error) {
	return r.LocationRepository.ListActiveEvses(ctx, obj.ID)
}

func (r *locationResolver) Directions(ctx context.Context, obj *db.Location) ([]db.DisplayText, error) {
	return r.LocationRepository.ListLocationDirections(ctx, obj.ID)
}

func (r *locationResolver) Operator(ctx context.Context, obj *db.Location) (*db.BusinessDetail, error) {
	if obj.OperatorID.Valid {
		if businessDetail, err := r.BusinessDetailRepository.GetBusinessDetail(ctx, obj.OperatorID.Int64); err == nil {
			return &businessDetail, nil
		}
	}

	return nil, nil
}

func (r *locationResolver) Suboperator(ctx context.Context, obj *db.Location) (*db.BusinessDetail, error) {
	if obj.SuboperatorID.Valid {
		if businessDetail, err := r.BusinessDetailRepository.GetBusinessDetail(ctx, obj.SuboperatorID.Int64); err == nil {
			return &businessDetail, nil
		}
	}

	return nil, nil
}

func (r *locationResolver) Owner(ctx context.Context, obj *db.Location) (*db.BusinessDetail, error) {
	if obj.OwnerID.Valid {
		if businessDetail, err := r.BusinessDetailRepository.GetBusinessDetail(ctx, obj.OwnerID.Int64); err == nil {
			return &businessDetail, nil
		}
	}

	return nil, nil
}

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

func (r *locationResolver) TimeZone(ctx context.Context, obj *db.Location) (*string, error) {
	return util.NullString(obj.TimeZone)
}

func (r *locationResolver) OpeningTime(ctx context.Context, obj *db.Location) (*db.OpeningTime, error) {
	if obj.OpeningTimeID.Valid {
		if openingTime, err := r.OpeningTimeRepository.GetOpeningTime(ctx, obj.OpeningTimeID.Int64); err == nil {
			return &openingTime, nil
		}
	}

	return nil, nil
}

func (r *locationResolver) Images(ctx context.Context, obj *db.Location) ([]db.Image, error) {
	return r.LocationRepository.ListLocationImages(ctx, obj.ID)
}

func (r *locationResolver) EnergyMix(ctx context.Context, obj *db.Location) (*db.EnergyMix, error) {
	if obj.EnergyMixID.Valid {
		if energyMix, err := r.EnergyMixRepository.GetEnergyMix(ctx, obj.EnergyMixID.Int64); err == nil {
			return &energyMix, nil
		}
	}

	return nil, nil
}

func (r *locationResolver) LastUpdated(ctx context.Context, obj *db.Location) (string, error) {
	return obj.LastUpdated.Format(time.RFC3339), nil
}

// Location returns graph.LocationResolver implementation.
func (r *Resolver) Location() graph.LocationResolver { return &locationResolver{r} }

type locationResolver struct{ *Resolver }
