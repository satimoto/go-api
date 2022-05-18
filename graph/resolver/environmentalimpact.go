package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-datastore/pkg/db"
)

func (r *environmentalImpactResolver) Source(ctx context.Context, obj *db.EnvironmentalImpact) (string, error) {
	return string(obj.Source), nil
}

// EnvironmentalImpact returns graph.EnvironmentalImpactResolver implementation.
func (r *Resolver) EnvironmentalImpact() graph.EnvironmentalImpactResolver {
	return &environmentalImpactResolver{r}
}

type environmentalImpactResolver struct{ *Resolver }
