package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/util"
	"github.com/satimoto/go-datastore/pkg/db"
)

func (r *energyMixResolver) EnergySources(ctx context.Context, obj *db.EnergyMix) ([]db.EnergySource, error) {
	return r.EnergyMixResolver.Repository.ListEnergySources(ctx, obj.ID)
}

func (r *energyMixResolver) EnvironmentalImpact(ctx context.Context, obj *db.EnergyMix) ([]db.EnvironmentalImpact, error) {
	return r.EnergyMixResolver.Repository.ListEnvironmentalImpacts(ctx, obj.ID)
}

func (r *energyMixResolver) SupplierName(ctx context.Context, obj *db.EnergyMix) (*string, error) {
	return util.NullString(obj.SupplierName)
}

func (r *energyMixResolver) EnergyProductName(ctx context.Context, obj *db.EnergyMix) (*string, error) {
	return util.NullString(obj.EnergyProductName)
}

// EnergyMix returns graph.EnergyMixResolver implementation.
func (r *Resolver) EnergyMix() graph.EnergyMixResolver { return &energyMixResolver{r} }

type energyMixResolver struct{ *Resolver }
