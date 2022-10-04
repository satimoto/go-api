package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/util"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/geom"
)

// Identifier is the resolver for the identifier field.
func (r *evseResolver) Identifier(ctx context.Context, obj *db.Evse) (*string, error) {
	return util.NullString(obj.Identifier)
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

// Evse returns graph.EvseResolver implementation.
func (r *Resolver) Evse() graph.EvseResolver { return &evseResolver{r} }

type evseResolver struct{ *Resolver }
