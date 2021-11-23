package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strings"
	"time"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/util"
	"github.com/satimoto/go-datastore/db"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *mutationResolver) CreateEmailSubscription(ctx context.Context, input graph.CreateEmailSubscriptionInput) (*db.EmailSubscription, error) {
	emailSubscription, err := r.EmailSubscriptionResolver.Repository.CreateEmailSubscription(ctx, db.CreateEmailSubscriptionParams{
		Email:       strings.ToLower(input.Email),
		Code:        util.RandomVerificationCode(),
		IsVerified:  false,
		CreatedDate: time.Now(),
	})

	if err != nil {
		return nil, gqlerror.Errorf("Email subscription already exists")
	}

	// TODO: Send user verification email

	return &emailSubscription, nil
}

func (r *mutationResolver) VerifyEmailSubscription(ctx context.Context, input graph.VerifyEmailSubscriptionInput) (*db.EmailSubscription, error) {
	emailSubscription, err := r.EmailSubscriptionResolver.Repository.GetEmailSubscriptionByEmail(ctx, strings.ToLower(input.Email))

	if err != nil {
		return nil, gqlerror.Errorf("Email subscription not found")
	}

	if emailSubscription.Code == input.Code {
		emailSubscription, err = r.EmailSubscriptionResolver.Repository.UpdateEmailSubscription(ctx, db.UpdateEmailSubscriptionParams{
			ID:         emailSubscription.ID,
			Email:      emailSubscription.Email,
			Code:       emailSubscription.Code,
			IsVerified: true,
		})

		if err != nil {
			return nil, gqlerror.Errorf("Error updating email subscription")
		}

		return &emailSubscription, nil
	}

	return nil, gqlerror.Errorf("Invalid verification code")
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
