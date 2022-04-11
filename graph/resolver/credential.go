package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/authentication"
	"github.com/satimoto/go-api/internal/credential"
	"github.com/satimoto/go-datastore/db"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *mutationResolver) CreateCredential(ctx context.Context, input graph.CreateCredentialInput) (*db.Credential, error) {
	if userId := authentication.GetUserId(ctx); userId != nil {
		params := credential.NewCreateCredentialParams(input)
		businessDetail, err := r.CreateBusinessDetail(ctx, *input.BusinessDetail)

		if err != nil {
			return nil, err
		}

		params.BusinessDetailID = businessDetail.ID

		c, err := r.CredentialResolver.Repository.CreateCredential(ctx, params)

		if err != nil {
			return nil, gqlerror.Errorf("Error creating credential")
		}

		return &c, nil
	}

	return nil, gqlerror.Errorf("Not Authenticated")
}

func (r *credentialResolver) BusinessDetail(ctx context.Context, obj *db.Credential) (*db.BusinessDetail, error) {
	if businessDetail, err := r.BusinessDetailResolver.Repository.GetBusinessDetail(ctx, obj.BusinessDetailID); err == nil {
		return &businessDetail, nil
	}

	return nil, nil
}

// Credential returns graph.CredentialResolver implementation.
func (r *Resolver) Credential() graph.CredentialResolver { return &credentialResolver{r} }

type credentialResolver struct{ *Resolver }
