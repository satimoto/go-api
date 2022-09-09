package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"
	"net/url"

	"github.com/google/uuid"
	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/authentication"
	"github.com/satimoto/go-api/internal/lnurl"
	"github.com/satimoto/go-api/internal/lnurl/auth"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *mutationResolver) CreateAuthentication(ctx context.Context, action graph.AuthenticationAction) (*graph.CreateAuthentication, error) {
	authentication, err := r.AuthenticationResolver.Repository.CreateAuthentication(ctx, db.CreateAuthenticationParams{
		Action:    db.AuthenticationActions(action),
		Challenge: lnurl.RandomK1(),
		Code:      uuid.NewString(),
	})

	if err != nil {
		util.LogOnError("API003", "Error creating authentication", err)
		return nil, gqlerror.Errorf("Error creating authentication")
	}

	params := url.Values{}
	params.Add("k1", authentication.Challenge)
	params.Add("tag", action.String())

	callbackUrl, err := auth.GenerateLnUrl("v1", authentication.Challenge)

	if err != nil {
		util.LogOnError("API004", "Error generating LNURL", err)
		return nil, gqlerror.Errorf("Error creating authentication")
	}

	return &graph.CreateAuthentication{
		Code:  authentication.Code,
		LnURL: callbackUrl,
	}, nil
}

func (r *mutationResolver) ExchangeAuthentication(ctx context.Context, code string) (*graph.ExchangeAuthentication, error) {
	auth, err := r.AuthenticationResolver.Repository.GetAuthenticationByCode(ctx, code)

	if err != nil {
		util.LogOnError("API005", "Authentication not found", err)
		log.Printf("API005: Code=%v", code)
		return nil, gqlerror.Errorf("Authentication not found")
	}

	if !auth.Signature.Valid {
		log.Printf("API006: Authentication not yet verified")
		log.Printf("API006: AuthChallenge=%v", auth.Challenge)
		return nil, gqlerror.Errorf("Authentication not yet verified")
	}

	user, err := r.UserRepository.GetUserByLinkingPubkey(ctx, auth.LinkingPubkey.String)

	if err != nil {
		util.LogOnError("API007", "No linked user", err)
		return nil, gqlerror.Errorf("No linked user")
	}

	token, err := authentication.SignToken(user)

	if err != nil {
		util.LogOnError("API008", "Error signing token", err)
		return nil, gqlerror.Errorf("Error signing token")
	}

	r.AuthenticationResolver.Repository.DeleteAuthentication(ctx, auth.ID)

	return &graph.ExchangeAuthentication{
		Token: token,
	}, nil
}

func (r *queryResolver) VerifyAuthentication(ctx context.Context, code string) (*graph.VerifyAuthentication, error) {
	authentication, err := r.AuthenticationResolver.Repository.GetAuthenticationByCode(ctx, code)

	if err != nil || !authentication.Signature.Valid {
		return &graph.VerifyAuthentication{
			Verified: false,
		}, nil
	}

	return &graph.VerifyAuthentication{
		Verified: true,
	}, nil
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
