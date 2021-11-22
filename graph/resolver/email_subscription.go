package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-datastore/db"
)

func (r *mutationResolver) CreateEmailSubscription(ctx context.Context, input graph.CreateEmailSubscriptionInput) (*db.EmailSubscription, error) {
	if emailSubscription, err := r.EmailSubscriptionResolver.Repository.GetEmailSubscriptionByEmail(ctx, input.Email); err != nil {
		return &emailSubscription, err
	}

	emailSubscription, err := r.EmailSubscriptionResolver.Repository.CreateEmailSubscription(ctx, db.CreateEmailSubscriptionParams{
		Email: input.Email,
		IsVerified: false,
		CreatedDate: time.Now(),
	})

	return &emailSubscription, err
}

func (r *mutationResolver) VerifyEmailSubscription(ctx context.Context, input graph.VerifyEmailSubscriptionInput) (*db.EmailSubscription, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
