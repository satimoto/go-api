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

func (r *evseResolver) EvseID(ctx context.Context, obj *db.Evse) (*string, error) {
	return util.NullString(obj.EvseID)
}

func (r *evseResolver) Status(ctx context.Context, obj *db.Evse) (string, error) {
	return string(obj.Status), nil
}

func (r *evseResolver) StatusSchedule(ctx context.Context, obj *db.Evse) ([]db.StatusSchedule, error) {
	return r.EvseResolver.Repository.ListStatusSchedules(ctx, obj.ID)
}

func (r *evseResolver) Capabilities(ctx context.Context, obj *db.Evse) ([]graph.TextDescription, error) {
	list := []graph.TextDescription{}

	if capabilities, err := r.EvseResolver.Repository.ListEvseCapabilities(ctx, obj.ID); err == nil {
		for _, capability := range capabilities {
			list = append(list, graph.TextDescription{
				Text:        capability.Text,
				Description: capability.Description,
			})
		}
	}

	return list, nil
}

func (r *evseResolver) Connectors(ctx context.Context, obj *db.Evse) ([]db.Connector, error) {
	return r.EvseResolver.Repository.ListConnectors(ctx, obj.ID)
}

func (r *evseResolver) FloorLevel(ctx context.Context, obj *db.Evse) (*string, error) {
	return util.NullString(obj.FloorLevel)
}

func (r *evseResolver) Geom(ctx context.Context, obj *db.Evse) (*geom.Geometry4326, error) {
	return util.NullGeometry(obj.Geom)
}

func (r *evseResolver) PhysicalReference(ctx context.Context, obj *db.Evse) (*string, error) {
	return util.NullString(obj.PhysicalReference)
}

func (r *evseResolver) Directions(ctx context.Context, obj *db.Evse) ([]db.DisplayText, error) {
	return r.EvseResolver.Repository.ListEvseDirections(ctx, obj.ID)
}

func (r *evseResolver) ParkingRestrictions(ctx context.Context, obj *db.Evse) ([]graph.TextDescription, error) {
	list := []graph.TextDescription{}

	if parkingRestrictions, err := r.EvseResolver.Repository.ListEvseParkingRestrictions(ctx, obj.ID); err == nil {
		for _, parkingRestriction := range parkingRestrictions {
			list = append(list, graph.TextDescription{
				Text:        parkingRestriction.Text,
				Description: parkingRestriction.Description,
			})
		}
	}

	return list, nil

}
func (r *evseResolver) Images(ctx context.Context, obj *db.Evse) ([]db.Image, error) {
	return r.EvseResolver.Repository.ListEvseImages(ctx, obj.ID)
}

func (r *evseResolver) LastUpdated(ctx context.Context, obj *db.Evse) (string, error) {
	return obj.LastUpdated.Format(time.RFC3339Nano), nil
}

// Evse returns graph.EvseResolver implementation.
func (r *Resolver) Evse() graph.EvseResolver { return &evseResolver{r} }

type evseResolver struct{ *Resolver }
