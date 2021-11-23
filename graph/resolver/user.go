package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-datastore/db"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *mutationResolver) RegisterUser(ctx context.Context, input graph.RegisterUserInput) (*db.User, error) {
	node, err := r.NodeResolver.Repository.CreateNode(ctx, db.CreateNodeParams{
		Pubkey:  input.Node.Pubkey,
		Address: input.Node.Address,
	})

	if err != nil {
		return nil, gqlerror.Errorf("Node already exists")
	}

	user, err := r.UserResolver.Repository.CreateUser(ctx, db.CreateUserParams{
		DeviceToken: input.DeviceToken,
		NodeID:      node.ID,
	})

	if err != nil {
		return nil, gqlerror.Errorf("User already exists")
	}

	return &user, nil
}

func (r *queryResolver) Nodes(ctx context.Context) ([]db.Node, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Node(ctx context.Context, obj *db.User) (*db.Node, error) {
	node, err := r.NodeResolver.Repository.GetNode(ctx, obj.NodeID)

	if err != nil {
		return nil, gqlerror.Errorf("Node not found")
	}

	return &node, nil
}

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

// User returns graph.UserResolver implementation.
func (r *Resolver) User() graph.UserResolver { return &userResolver{r} }

type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
