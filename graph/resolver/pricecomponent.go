package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-datastore/pkg/db"
)

func (r *priceComponentResolver) Type(ctx context.Context, obj *db.PriceComponent) (string, error) {
	return string(obj.Type), nil
}

// PriceComponent returns graph.PriceComponentResolver implementation.
func (r *Resolver) PriceComponent() graph.PriceComponentResolver {
	return &priceComponentResolver{r}
}

type priceComponentResolver struct{ *Resolver }
