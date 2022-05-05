package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"log"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/authentication"
	"github.com/satimoto/go-api/internal/credential"
	"github.com/satimoto/go-datastore/db"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *mutationResolver) CreateCredential(ctx context.Context, input graph.CreateCredentialInput) (*db.Credential, error) {
	if userId := authentication.GetUserId(ctx); userId != nil {
		credentialRequest := credential.NewCreateCredentialRequest(input)
		credentialResponse, err := r.OcpiService.CreateCredential(ctx, credentialRequest)

		if err != nil {
			log.Printf("Error CreateCredential CreateCredential: %v", err)
			log.Printf("%#v", credentialRequest)
			return nil, errors.New("Error creating credential")
		}

		return credential.NewCreateCredential(*credentialResponse), nil
	}

	return nil, gqlerror.Errorf("Not Authenticated")
}

func (r *mutationResolver) RegisterCredential(ctx context.Context, input graph.RegisterCredentialInput) (*graph.Result, error) {
	if userId := authentication.GetUserId(ctx); userId != nil {
		credentialRequest := credential.NewRegisterCredentialRequest(input)
		credentialResponse, err := r.OcpiService.RegisterCredential(ctx, credentialRequest)

		if err != nil {
			log.Printf("Error RegisterCredential RegisterCredential: %v", err)
			log.Printf("%#v", credentialRequest)
			return nil, errors.New("Error registering credential")
		}

		return &graph.Result{ID: credentialResponse.Id}, nil
	}

	return nil, gqlerror.Errorf("Not Authenticated")
}

func (r *mutationResolver) UnregisterCredential(ctx context.Context, input graph.UnregisterCredentialInput) (*graph.Result, error) {
	if userId := authentication.GetUserId(ctx); userId != nil {
		credentialRequest := credential.NewUnregisterCredentialRequest(input)
		credentialResponse, err := r.OcpiService.UnregisterCredential(ctx, credentialRequest)

		if err != nil {
			log.Printf("Error UnregisterCredential UnregisterCredential: %v", err)
			log.Printf("%#v", credentialRequest)
			return nil, errors.New("Error unregistering credential")
		}

		return &graph.Result{ID: credentialResponse.Id}, nil
	}

	return nil, gqlerror.Errorf("Not Authenticated")
}
