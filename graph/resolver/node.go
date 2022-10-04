package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-datastore/pkg/db"
)

// Addr is the resolver for the addr field.
func (r *nodeResolver) Addr(ctx context.Context, obj *db.Node) (string, error) {
	return obj.NodeAddr, nil
}

// Node returns graph.NodeResolver implementation.
func (r *Resolver) Node() graph.NodeResolver { return &nodeResolver{r} }

type nodeResolver struct{ *Resolver }
