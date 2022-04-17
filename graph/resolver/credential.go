package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/authentication"
	"github.com/satimoto/go-api/internal/credential"
	"github.com/satimoto/go-ocpi-api/ocpirpc/credentialrpc"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"google.golang.org/grpc"
)

func (r *mutationResolver) CreateCredential(ctx context.Context, input graph.CreateCredentialInput) (*graph.CreateCredential, error) {
	if userId := authentication.GetUserId(ctx); userId != nil {
		ocpiRpcAddress := os.Getenv("OCPI_RPC_ADDRESS")
		conn, err := grpc.Dial(ocpiRpcAddress, grpc.WithInsecure())

		if err != nil {
			log.Printf("Error CreateCredential Dial: %v", err)
			log.Printf("OCPI_RPC_ADDRESS=%v", ocpiRpcAddress)
			return nil, errors.New("Error creating credential")
		}

		defer conn.Close()
		credentialClient := credentialrpc.NewCredentialServiceClient(conn)
		credentialRequest := credential.NewCreateCredentialRequest(input)
		credentialResponse, err := credentialClient.CreateCredential(ctx, credentialRequest)

		if err != nil {
			log.Printf("Error CreateCredential CreateCredential: %v", err)
			log.Printf("%#v", credentialRequest)
			return nil, errors.New("Error creating credential")
		}

		return credential.NewCreateCredential(*credentialResponse), nil
	}

	return nil, gqlerror.Errorf("Not Authenticated")
}
