package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/credential"
	metrics "github.com/satimoto/go-api/internal/metric"
	"github.com/satimoto/go-api/internal/middleware"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/param"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// CreateCredential is the resolver for the createCredential field.
func (r *mutationResolver) CreateCredential(ctx context.Context, input graph.CreateCredentialInput) (*db.Credential, error) {
	backgroundCtx := context.Background()

	if user := middleware.GetCtxUser(ctx, r.UserRepository); user != nil && user.IsAdmin {
		credentialRequest := credential.NewCreateCredentialRequest(input)
		credentialResponse, err := r.OcpiService.CreateCredential(backgroundCtx, credentialRequest)

		if err != nil {
			metrics.RecordError("API012", "Error creating credential", err)
			log.Printf("API012: CreateCredentialRequest=%#v", credentialRequest)
			return nil, gqlerror.Errorf("Error creating credential")
		}

		return credential.NewCreateCredential(*credentialResponse), nil
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

// RegisterCredential is the resolver for the registerCredential field.
func (r *mutationResolver) RegisterCredential(ctx context.Context, input graph.RegisterCredentialInput) (*graph.ResultID, error) {
	backgroundCtx := context.Background()

	if user := middleware.GetCtxUser(ctx, r.UserRepository); user != nil && user.IsAdmin {
		credentialRequest := credential.NewRegisterCredentialRequest(input)
		credentialResponse, err := r.OcpiService.RegisterCredential(backgroundCtx, credentialRequest)

		if err != nil {
			metrics.RecordError("API013", "Error registering credential", err)
			log.Printf("API013: RegisterCredentialRequest=%#v", credentialRequest)
			return nil, gqlerror.Errorf("Error registering credential")
		}

		return &graph.ResultID{ID: credentialResponse.Id}, nil
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

// SyncCredential is the resolver for the syncCredential field.
func (r *mutationResolver) SyncCredential(ctx context.Context, input graph.SyncCredentialInput) (*graph.ResultID, error) {
	backgroundCtx := context.Background()

	if user := middleware.GetCtxUser(ctx, r.UserRepository); user != nil && user.IsAdmin {
		credentialRequest := credential.NewSyncCredentialRequest(input)
		credentialResponse, err := r.OcpiService.SyncCredential(backgroundCtx, credentialRequest)

		if err != nil {
			metrics.RecordError("API028", "Error syncing credential", err)
			log.Printf("API028: SyncCredentialRequest=%#v", credentialRequest)
			return nil, gqlerror.Errorf("Error syncing credential")
		}

		return &graph.ResultID{ID: credentialResponse.Id}, nil
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

// UnregisterCredential is the resolver for the unregisterCredential field.
func (r *mutationResolver) UnregisterCredential(ctx context.Context, input graph.UnregisterCredentialInput) (*graph.ResultID, error) {
	backgroundCtx := context.Background()

	if user := middleware.GetCtxUser(ctx, r.UserRepository); user != nil && user.IsAdmin {
		credentialRequest := credential.NewUnregisterCredentialRequest(input)
		credentialResponse, err := r.OcpiService.UnregisterCredential(backgroundCtx, credentialRequest)

		if err != nil {
			metrics.RecordError("API014", "Error unregistering credential", err)
			log.Printf("API014: UnregisterCredentialRequest=%#v", credentialRequest)
			return nil, gqlerror.Errorf("Error unregistering credential")
		}

		return &graph.ResultID{ID: credentialResponse.Id}, nil
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

// UpdateCredential is the resolver for the updateCredential field.
func (r *mutationResolver) UpdateCredential(ctx context.Context, input graph.UpdateCredentialInput) (*graph.ResultID, error) {
	backgroundCtx := context.Background()

	if user := middleware.GetCtxUser(ctx, r.UserRepository); user != nil {
		if credential, err := r.CredentialRepository.GetCredential(backgroundCtx, input.ID); err == nil {
			updateCredentialParams := param.NewUpdateCredentialParams(credential)
			updateCredentialParams.IsAvailable = input.IsAvailable

			updatedCredential, err := r.CredentialRepository.UpdateCredential(backgroundCtx, updateCredentialParams)

			if err != nil {
				metrics.RecordError("API078", "Error updating credential", err)
				log.Printf("API078: Params=%#v", updateCredentialParams)
				return nil, gqlerror.Errorf("Error updating credential")
			}

			return &graph.ResultID{ID: updatedCredential.ID}, nil
		}
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

// ListCredentials is the resolver for the listCredentials field.
func (r *queryResolver) ListCredentials(ctx context.Context) ([]db.Credential, error) {
	backgroundCtx := context.Background()

	if user := middleware.GetCtxUser(ctx, r.UserRepository); user != nil && user.IsAdmin {
		return r.CredentialRepository.ListCredentials(backgroundCtx)
	}

	return nil, gqlerror.Errorf("Not authenticated")
}
