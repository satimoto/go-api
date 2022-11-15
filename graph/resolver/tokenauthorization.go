package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/hex"
	"errors"
	"log"

	"github.com/satimoto/go-api/graph"
	metrics "github.com/satimoto/go-api/internal/metric"
	"github.com/satimoto/go-api/internal/middleware"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi/ocpirpc"
	ocpiTokenAuthorization "github.com/satimoto/go-ocpi/pkg/ocpi/tokenauthorization"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// UpdateTokenAuthorization is the resolver for the updateTokenAuthorization field.
func (r *mutationResolver) UpdateTokenAuthorization(ctx context.Context, input graph.UpdateTokenAuthorizationInput) (*db.TokenAuthorization, error) {
	if userID := middleware.GetUserID(ctx); userID != nil {
		updateTokenAuthorizationRequest := &ocpirpc.UpdateTokenAuthorizationRequest{
			AuthorizationId: input.AuthorizationID,
			Authorize:       input.Authorized,
		}

		_, err := r.OcpiService.UpdateTokenAuthorization(ctx, updateTokenAuthorizationRequest)

		if err != nil {
			metrics.RecordError("API042", "Error updating token authorization", err)
			log.Printf("API042: Params=%#v", updateTokenAuthorizationRequest)
		}

		tokenAuthorization, err := r.TokenAuthorizationRepository.GetTokenAuthorizationByAuthorizationID(ctx, input.AuthorizationID)

		if err != nil {
			metrics.RecordError("API043", "Error retrieving token authorization", err)
			log.Printf("API043: AuthorizationID=%v", input.AuthorizationID)
			return nil, errors.New("Error updating token authorization")
		}

		return &tokenAuthorization, nil
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

// Location is the resolver for the location field.
func (r *tokenAuthorizationResolver) Location(ctx context.Context, obj *db.TokenAuthorization) (*db.Location, error) {
	if obj.LocationID.Valid {
		if location, err := r.LocationRepository.GetLocationByUid(ctx, obj.LocationID.String); err == nil {
			return &location, nil
		}
	}

	return nil, nil
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
