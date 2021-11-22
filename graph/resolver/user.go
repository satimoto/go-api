package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-datastore/db"
)

func (r *mutationResolver) RegisterUser(ctx context.Context, input graph.UserInput) (*db.User, error) {
	node, err := r.NodeResolver.Repository.CreateNode(ctx, db.CreateNodeParams{
		Pubkey:  input.Node.Pubkey,
		Address: input.Node.Address,
	})

	if err != nil {
		return nil, err
	}

	user, err := r.UserResolver.Repository.CreateUser(ctx, db.CreateUserParams{
		DeviceToken: input.DeviceToken,
		NodeID:      node.ID,
	})

	return &user, err
}

func (r *queryResolver) Users(ctx context.Context) ([]db.User, error) {
	return r.UserResolver.Repository.ListUsers(ctx)
}

func (r *userResolver) Node(ctx context.Context, obj *db.User) (*db.Node, error) {
	node, err := r.NodeResolver.Repository.GetNode(ctx, obj.NodeID)

	return &node, err
}

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

// User returns graph.UserResolver implementation.
func (r *Resolver) User() graph.UserResolver { return &userResolver{r} }

type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
