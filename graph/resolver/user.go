package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/authentication"
	"github.com/satimoto/go-datastore/db"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input graph.CreateUserInput) (*db.User, error) {
	auth, err := r.AuthenticationResolver.Repository.GetAuthenticationByCode(ctx, input.Code)

	if err != nil {
		log.Printf("Authentication not found: code %s: %s", input.Code, err.Error())
		return nil, gqlerror.Errorf("Authentication not found")
	}

	if !auth.Signature.Valid {
		log.Printf("Authentication not yet verified: %s", auth.Challenge)
		return nil, gqlerror.Errorf("Authentication not yet verified")
	}

	user, err := r.UserResolver.Repository.CreateUser(ctx, db.CreateUserParams{
		LinkingPubkey: auth.LinkingPubkey.String,
		Pubkey:        input.Pubkey,
		DeviceToken:   input.DeviceToken,
	})

	if err != nil {
		log.Printf("User already exists: %s", err.Error())
		return nil, gqlerror.Errorf("User already exists")
	}

	return &user, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input graph.UpdateUserInput) (*db.User, error) {
	if userId := authentication.GetUserId(ctx); userId != nil {
		user, err := r.UserResolver.Repository.UpdateUser(ctx, db.UpdateUserParams{
			ID:          *userId,
			DeviceToken: input.DeviceToken,
		})

		if err != nil {
			log.Printf("Error updating user: %s", err.Error())
			return nil, gqlerror.Errorf("Error updating user")
		}

		return &user, nil
	}

	return nil, gqlerror.Errorf("Not Authenticated")
}
