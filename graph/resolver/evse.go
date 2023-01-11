package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/middleware"
	"github.com/satimoto/go-api/internal/util"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/geom"
	dbUtil "github.com/satimoto/go-datastore/pkg/util"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// EvseID is the resolver for the evseId field.
func (r *evseResolver) EvseID(ctx context.Context, obj *db.Evse) (*string, error) {
	return util.NullString(obj.EvseID)
}

// Identifier is the resolver for the identifier field.
func (r *evseResolver) Identifier(ctx context.Context, obj *db.Evse) (*string, error) {
	return util.NullString(obj.Identifier)
}

// Location is the resolver for the location field.
func (r *evseResolver) Location(ctx context.Context, obj *db.Evse) (*db.Location, error) {
	if location, err := r.LocationRepository.GetLocation(ctx, obj.LocationID); err == nil {
		return &location, nil
	}

	return nil, gqlerror.Errorf("Location not found")
}

// Status is the resolver for the status field.
func (r *evseResolver) Status(ctx context.Context, obj *db.Evse) (string, error) {
	return string(obj.Status), nil
}

// StatusSchedule is the resolver for the statusSchedule field.
func (r *evseResolver) StatusSchedule(ctx context.Context, obj *db.Evse) ([]db.StatusSchedule, error) {
	return r.EvseRepository.ListStatusSchedules(ctx, obj.ID)
}

// Capabilities is the resolver for the capabilities field.
func (r *evseResolver) Capabilities(ctx context.Context, obj *db.Evse) ([]graph.TextDescription, error) {
	list := []graph.TextDescription{}

	if capabilities, err := r.EvseRepository.ListEvseCapabilities(ctx, obj.ID); err == nil {
		for _, capability := range capabilities {
			list = append(list, graph.TextDescription{
				Text:        capability.Text,
				Description: capability.Description,
			})
		}
	}

	return list, nil
}

// Connectors is the resolver for the connectors field.
func (r *evseResolver) Connectors(ctx context.Context, obj *db.Evse) ([]db.Connector, error) {
	return r.EvseRepository.ListConnectors(ctx, obj.ID)
}

// FloorLevel is the resolver for the floorLevel field.
func (r *evseResolver) FloorLevel(ctx context.Context, obj *db.Evse) (*string, error) {
	return util.NullString(obj.FloorLevel)
}

// Geom is the resolver for the geom field.
func (r *evseResolver) Geom(ctx context.Context, obj *db.Evse) (*geom.Geometry4326, error) {
	return util.NullGeometry(obj.Geom)
}

// PhysicalReference is the resolver for the physicalReference field.
func (r *evseResolver) PhysicalReference(ctx context.Context, obj *db.Evse) (*string, error) {
	return util.NullString(obj.PhysicalReference)
}

// Directions is the resolver for the directions field.
func (r *evseResolver) Directions(ctx context.Context, obj *db.Evse) ([]db.DisplayText, error) {
	return r.EvseRepository.ListEvseDirections(ctx, obj.ID)
}

// ParkingRestrictions is the resolver for the parkingRestrictions field.
func (r *evseResolver) ParkingRestrictions(ctx context.Context, obj *db.Evse) ([]graph.TextDescription, error) {
	list := []graph.TextDescription{}

	if parkingRestrictions, err := r.EvseRepository.ListEvseParkingRestrictions(ctx, obj.ID); err == nil {
		for _, parkingRestriction := range parkingRestrictions {
			list = append(list, graph.TextDescription{
				Text:        parkingRestriction.Text,
				Description: parkingRestriction.Description,
			})
		}
	}

	return list, nil
}

// Images is the resolver for the images field.
func (r *evseResolver) Images(ctx context.Context, obj *db.Evse) ([]db.Image, error) {
	return r.EvseRepository.ListEvseImages(ctx, obj.ID)
}

// LastUpdated is the resolver for the lastUpdated field.
func (r *evseResolver) LastUpdated(ctx context.Context, obj *db.Evse) (string, error) {
	return obj.LastUpdated.Format(time.RFC3339), nil
}

// GetEvse is the resolver for the getEvse field.
func (r *queryResolver) GetEvse(reqCtx context.Context, input graph.GetEvseInput) (*db.Evse, error) {
	ctx := context.Background()
	
	if userID := middleware.GetUserID(reqCtx); userID != nil {
		if input.ID != nil {
			if evse, err := r.EvseRepository.GetEvse(ctx, *input.ID); err == nil {
				return &evse, nil
			}
		} else if input.UID != nil {
			if evse, err := r.EvseRepository.GetEvseByUid(ctx, *input.UID); err == nil {
				return &evse, nil
			}
		} else if input.EvseID != nil {
			evse, err := r.EvseRepository.GetEvseByEvseID(ctx, dbUtil.SqlNullString(input.EvseID))

			if err != nil {
				// Like search to get the best match
				likeEvseID := fmt.Sprintf("%%%s", *input.EvseID)

				if strings.Contains(*input.EvseID, "*") {
					likeEvseID = fmt.Sprintf("%s%%", *input.EvseID)
				}

				if evses, err := r.EvseRepository.ListEvsesLikeEvseID(ctx, dbUtil.SqlNullString(likeEvseID)); err == nil {
					log.Printf("Like search for %v matched %v result(s)", likeEvseID, len(evses))

					if len(evses) > 0 {
						return &evses[0], nil
					}
				}

				return nil, gqlerror.Errorf("Evse not found")
			}

			return &evse, nil
		} else if input.Identifier != nil {
			dashRegex := regexp.MustCompile(`-`)
			identifier := strings.ToUpper(dashRegex.ReplaceAllString(*input.Identifier, "*"))

			if evse, err := r.EvseRepository.GetEvseByIdentifier(ctx, dbUtil.SqlNullString(identifier)); err == nil {
				return &evse, nil
			}
		}

		return nil, gqlerror.Errorf("Evse not found")
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

// Evse returns graph.EvseResolver implementation.
func (r *Resolver) Evse() graph.EvseResolver { return &evseResolver{r} }

type evseResolver struct{ *Resolver }
