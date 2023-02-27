package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/satimoto/go-api/graph"
	metrics "github.com/satimoto/go-api/internal/metric"
	"github.com/satimoto/go-api/internal/middleware"
	"github.com/satimoto/go-api/internal/param"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// CreateParty is the resolver for the createParty field.
func (r *mutationResolver) CreateParty(ctx context.Context, input graph.CreatePartyInput) (*db.Party, error) {
	backgroundCtx := context.Background()

	if user := middleware.GetCtxUser(ctx, r.UserRepository); user != nil && user.IsAdmin {
		createPartyParams := param.NewCreatePartyParams(input)

		party, err := r.PartyRepository.CreateParty(backgroundCtx, createPartyParams)

		if err != nil {
			metrics.RecordError("API055", "Error creating party", err)
			log.Printf("API055: Params=%#v", createPartyParams)
			return nil, gqlerror.Errorf("Error creating party")
		}

		return &party, nil
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

// UpdateParty is the resolver for the updateParty field.
func (r *mutationResolver) UpdateParty(ctx context.Context, input graph.UpdatePartyInput) (*db.Party, error) {
	backgroundCtx := context.Background()

	if user := middleware.GetCtxUser(ctx, r.UserRepository); user != nil && user.IsAdmin {
		updatePartyByCredentialParams := param.NewUpdatePartyByCredentialParams(input)

		party, err := r.PartyRepository.UpdatePartyByCredential(backgroundCtx, updatePartyByCredentialParams)

		if err != nil {
			metrics.RecordError("API056", "Error updating party", err)
			log.Printf("API056: Params=%#v", updatePartyByCredentialParams)
			return nil, gqlerror.Errorf("Error updating party")
		}

		return &party, nil
	}

	return nil, gqlerror.Errorf("Not authenticated")
}
