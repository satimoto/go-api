package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/satimoto/go-api/graph"
	metrics "github.com/satimoto/go-api/internal/metric"
	"github.com/satimoto/go-api/internal/middleware"
	"github.com/satimoto/go-api/internal/token"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi/ocpirpc"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// CreateToken is the resolver for the createToken field.
func (r *mutationResolver) CreateToken(ctx context.Context, input graph.CreateTokenInput) (*db.Token, error) {
	backgroundCtx := context.Background()

	if userID := middleware.GetUserID(ctx); userID != nil {
		if _, err := r.TokenResolver.Repository.GetTokenByUid(backgroundCtx, input.UID); err == nil {
			return nil, gqlerror.Errorf("Error creating token")
		}

		createTokenRequest := &ocpirpc.CreateTokenRequest{
			UserId:    *userID,
			Uid:       input.UID,
			Type:      string(db.TokenTypeRFID),
			Whitelist: string(db.TokenWhitelistTypeNEVER),
		}

		if input.Allowed != nil {
			createTokenRequest.Allowed = *input.Allowed
		}

		if input.Type != nil {
			createTokenRequest.Type = *input.Type
		}

		createTokenResponse, err := r.OcpiService.CreateToken(backgroundCtx, createTokenRequest)

		if err != nil {
			metrics.RecordError("API041", "Error creating token", err)
			log.Printf("API041: CreateTokenRequest=%#v", createTokenRequest)
			return nil, gqlerror.Errorf("Error creating token")
		}

		return token.NewCreateToken(*createTokenResponse), nil
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

// UpdateTokens is the resolver for the updateTokens field.
func (r *mutationResolver) UpdateTokens(ctx context.Context, input graph.UpdateTokensInput) (*graph.ResultOk, error) {
	backgroundCtx := context.Background()

	if user := middleware.GetUser(ctx, r.UserRepository); user != nil && user.IsAdmin {
		updateTokensRequest := &ocpirpc.UpdateTokensRequest{
			UserId:  input.UserID,
			Allowed: input.Allowed,
		}

		if input.UID != nil {
			updateTokensRequest.Uid = *input.UID
		}

		updateTokensResponse, err := r.OcpiService.UpdateTokens(backgroundCtx, updateTokensRequest)

		if err != nil {
			metrics.RecordError("API042", "Error updating token", err)
			log.Printf("API042: UpdateTokensRequest=%#v", updateTokensResponse)
			return nil, gqlerror.Errorf("Error updating token")
		}

		return &graph.ResultOk{Ok: true}, nil
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

// ListTokens is the resolver for the listTokens field.
func (r *queryResolver) ListTokens(ctx context.Context) ([]db.Token, error) {
	backgroundCtx := context.Background()

	if userID := middleware.GetUserID(ctx); userID != nil {
		return r.TokenResolver.Repository.ListRfidTokensByUserID(backgroundCtx, *userID)
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

// Type is the resolver for the type field.
func (r *tokenResolver) Type(ctx context.Context, obj *db.Token) (string, error) {
	return string(obj.Type), nil
}

// VisualNumber is the resolver for the visualNumber field.
func (r *tokenResolver) VisualNumber(ctx context.Context, obj *db.Token) (*string, error) {
	return util.NilString(obj.VisualNumber), nil
}

// Token returns graph.TokenResolver implementation.
func (r *Resolver) Token() graph.TokenResolver { return &tokenResolver{r} }

type tokenResolver struct{ *Resolver }
