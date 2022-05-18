package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-datastore/pkg/db"
)

func (r *energySourceResolver) Source(ctx context.Context, obj *db.EnergySource) (string, error) {
	return string(obj.Source), nil
}

// EnergySource returns graph.EnergySourceResolver implementation.
func (r *Resolver) EnergySource() graph.EnergySourceResolver { return &energySourceResolver{r} }

type energySourceResolver struct{ *Resolver }
