package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/hex"
	"errors"
	"log"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/middleware"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/param"
	"github.com/satimoto/go-datastore/pkg/util"
	ocpiTokenAuthorization "github.com/satimoto/go-ocpi/pkg/ocpi/tokenauthorization"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// UpdateTokenAuthorization is the resolver for the updateTokenAuthorization field.
func (r *mutationResolver) UpdateTokenAuthorization(ctx context.Context, input graph.UpdateTokenAuthorizationInput) (*db.TokenAuthorization, error) {
	if userID := middleware.GetUserID(ctx); userID != nil {
		if tokenAuthorization, err := r.TokenAuthorizationRepository.GetTokenAuthorizationByAuthorizationID(ctx, input.AuthorizationID); err == nil {
			if user, err := r.UserRepository.GetUserByTokenID(ctx, tokenAuthorization.TokenID); err == nil && user.ID == *userID {
				updateTokenAuthorizationParams := param.NewUpdateTokenAuthorizationParams(tokenAuthorization)
				updateTokenAuthorizationParams.Authorized = input.Authorized

				tokenAuthorization, err = r.TokenAuthorizationRepository.UpdateTokenAuthorizationByAuthorizationID(ctx, updateTokenAuthorizationParams)

				if err != nil {
					util.LogOnError("API043", "Error updating token authorization", err)
					log.Printf("API043: Params=%#v", updateTokenAuthorizationParams)
					return nil, errors.New("Error updating token authorization")
				}

				return &tokenAuthorization, nil
			}
		}
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

// CountryCode is the resolver for the countryCode field.
func (r *tokenAuthorizationResolver) CountryCode(ctx context.Context, obj *db.TokenAuthorization) (*string, error) {
	return util.NilString(obj.CountryCode), nil
}

// PartyID is the resolver for the partyId field.
func (r *tokenAuthorizationResolver) PartyID(ctx context.Context, obj *db.TokenAuthorization) (*string, error) {
	return util.NilString(obj.PartyID), nil
}

// LocationUID is the resolver for the locationUid field.
func (r *tokenAuthorizationResolver) LocationUID(ctx context.Context, obj *db.TokenAuthorization) (*string, error) {
	return util.NilString(obj.LocationID), nil
}

// Token is the resolver for the token field.
func (r *tokenAuthorizationResolver) Token(ctx context.Context, obj *db.TokenAuthorization) (*db.Token, error) {
	if token, err := r.TokenResolver.Repository.GetToken(ctx, obj.TokenID); err == nil {
		return &token, nil
	}

	return nil, gqlerror.Errorf("Token not found")
}

// VerificationKey is the resolver for the verificationKey field.
func (r *tokenAuthorizationResolver) VerificationKey(ctx context.Context, obj *db.TokenAuthorization) (*string, error) {
	if verificationKey, err := ocpiTokenAuthorization.CreateVerificationKey(*obj); err == nil {
		encodedVerificationKey := hex.EncodeToString(verificationKey)
		return &encodedVerificationKey, nil
	}

	return nil, gqlerror.Errorf("Error generating verification key")
}

// TokenAuthorization returns graph.TokenAuthorizationResolver implementation.
func (r *Resolver) TokenAuthorization() graph.TokenAuthorizationResolver {
	return &tokenAuthorizationResolver{r}
}

type tokenAuthorizationResolver struct{ *Resolver }
