package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/credential"
	"github.com/satimoto/go-api/internal/middleware"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// CreateCredential is the resolver for the createCredential field.
func (r *mutationResolver) CreateCredential(ctx context.Context, input graph.CreateCredentialInput) (*db.Credential, error) {
	if user := middleware.GetUser(ctx, r.UserRepository); user != nil && user.IsAdmin {
		credentialRequest := credential.NewCreateCredentialRequest(input)
		credentialResponse, err := r.OcpiService.CreateCredential(ctx, credentialRequest)

		if err != nil {
			util.LogOnError("API012", "Error creating credential", err)
			log.Printf("API012: CreateCredentialRequest=%#v", credentialRequest)
			return nil, gqlerror.Errorf("Error creating credential")
		}

		return credential.NewCreateCredential(*credentialResponse), nil
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

// RegisterCredential is the resolver for the registerCredential field.
func (r *mutationResolver) RegisterCredential(ctx context.Context, input graph.RegisterCredentialInput) (*graph.ResultID, error) {
	if user := middleware.GetUser(ctx, r.UserRepository); user != nil && user.IsAdmin {
		credentialRequest := credential.NewRegisterCredentialRequest(input)
		credentialResponse, err := r.OcpiService.RegisterCredential(ctx, credentialRequest)

		if err != nil {
			util.LogOnError("API013", "Error registering credential", err)
			log.Printf("API013: RegisterCredentialRequest=%#v", credentialRequest)
			return nil, gqlerror.Errorf("Error registering credential")
		}

		return &graph.ResultID{ID: credentialResponse.Id}, nil
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

// SyncCredential is the resolver for the syncCredential field.
func (r *mutationResolver) SyncCredential(ctx context.Context, input graph.SyncCredentialInput) (*graph.ResultID, error) {
	if user := middleware.GetUser(ctx, r.UserRepository); user != nil && user.IsAdmin {
		credentialRequest := credential.NewSyncCredentialRequest(input)
		credentialResponse, err := r.OcpiService.SyncCredential(ctx, credentialRequest)

		if err != nil {
			util.LogOnError("API028", "Error syncing credential", err)
			log.Printf("API028: SyncCredentialRequest=%#v", credentialRequest)
			return nil, gqlerror.Errorf("Error syncing credential")
		}

		return &graph.ResultID{ID: credentialResponse.Id}, nil
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

// UnregisterCredential is the resolver for the unregisterCredential field.
func (r *mutationResolver) UnregisterCredential(ctx context.Context, input graph.UnregisterCredentialInput) (*graph.ResultID, error) {
	if user := middleware.GetUser(ctx, r.UserRepository); user != nil && user.IsAdmin {
		credentialRequest := credential.NewUnregisterCredentialRequest(input)
		credentialResponse, err := r.OcpiService.UnregisterCredential(ctx, credentialRequest)

		if err != nil {
			util.LogOnError("API014", "Error unregistering credential", err)
			log.Printf("API014: UnregisterCredentialRequest=%#v", credentialRequest)
			return nil, gqlerror.Errorf("Error unregistering credential")
		}

		return &graph.ResultID{ID: credentialResponse.Id}, nil
	}

	return nil, gqlerror.Errorf("Not authenticated")
}
