package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/util"
	"github.com/satimoto/go-datastore/pkg/db"
)

// EnergySources is the resolver for the energySources field.
func (r *energyMixResolver) EnergySources(ctx context.Context, obj *db.EnergyMix) ([]db.EnergySource, error) {
	return r.EnergyMixRepository.ListEnergySources(ctx, obj.ID)
}

// EnvironmentalImpact is the resolver for the environmentalImpact field.
func (r *energyMixResolver) EnvironmentalImpact(ctx context.Context, obj *db.EnergyMix) ([]db.EnvironmentalImpact, error) {
	return r.EnergyMixRepository.ListEnvironmentalImpacts(ctx, obj.ID)
}

// SupplierName is the resolver for the supplierName field.
func (r *energyMixResolver) SupplierName(ctx context.Context, obj *db.EnergyMix) (*string, error) {
	return util.NullString(obj.SupplierName)
}

// EnergyProductName is the resolver for the energyProductName field.
func (r *energyMixResolver) EnergyProductName(ctx context.Context, obj *db.EnergyMix) (*string, error) {
	return util.NullString(obj.EnergyProductName)
}

// EnergyMix returns graph.EnergyMixResolver implementation.
func (r *Resolver) EnergyMix() graph.EnergyMixResolver { return &energyMixResolver{r} }

type energyMixResolver struct{ *Resolver }
