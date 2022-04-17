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
	"github.com/satimoto/go-api/internal/command"
	"github.com/satimoto/go-ocpi-api/ocpirpc/commandrpc"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"google.golang.org/grpc"
)

func (r *mutationResolver) StartSession(ctx context.Context, input graph.StartSessionInput) (*graph.StartSession, error) {
	if userId := authentication.GetUserId(ctx); userId != nil {
		ocpiRpcAddress := os.Getenv("OCPI_RPC_ADDRESS")
		conn, err := grpc.Dial(ocpiRpcAddress, grpc.WithInsecure())

		if err != nil {
			log.Printf("Error StartSession Dial: %v", err)
			log.Printf("OCPI_RPC_ADDRESS=%v", ocpiRpcAddress)
			return nil, errors.New("Error starting session")
		}

		defer conn.Close()
		commandClient := commandrpc.NewCommandServiceClient(conn)
		startSessionRequest := command.NewStartSessionRequest(*userId, input)
		startSessionResponse, err := commandClient.StartSession(ctx, startSessionRequest)

		if err != nil {
			log.Printf("Error StartSession StartSession: %v", err)
			log.Printf("%#v", startSessionRequest)
			return nil, errors.New("Error starting session")
		}

		return command.NewStartSession(*startSessionResponse), nil
	}

	return nil, gqlerror.Errorf("Not Authenticated")
}

func (r *mutationResolver) StopSession(ctx context.Context, input graph.StopSessionInput) (*graph.StopSession, error) {
	if userId := authentication.GetUserId(ctx); userId != nil {
		ocpiRpcAddress := os.Getenv("OCPI_RPC_ADDRESS")
		conn, err := grpc.Dial(ocpiRpcAddress, grpc.WithInsecure())

		if err != nil {
			log.Printf("Error StopSession Dial: %v", err)
			log.Printf("OCPI_RPC_ADDRESS=%v", ocpiRpcAddress)
			return nil, errors.New("Error stopping session")
		}

		defer conn.Close()
		commandClient := commandrpc.NewCommandServiceClient(conn)
		stopSessionRequest := command.NewStopSessionRequest(*userId, input)
		stopSessionResponse, err := commandClient.StopSession(ctx, stopSessionRequest)

		if err != nil {
			log.Printf("Error StopSession StopSession: %v", err)
			log.Printf("%#v", stopSessionRequest)
			return nil, errors.New("Error stopping session")
		}

		return command.NewStopSession(*stopSessionResponse), nil
	}

	return nil, gqlerror.Errorf("Not Authenticated")
}