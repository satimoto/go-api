package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/authentication"
	"github.com/satimoto/go-api/internal/command"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *mutationResolver) StartSession(ctx context.Context, input graph.StartSessionInput) (*graph.StartSession, error) {
	if userId := authentication.GetUserId(ctx); userId != nil {
		startSessionRequest := command.NewStartSessionRequest(*userId, input)
		startSessionResponse, err := r.OcpiService.StartSession(ctx, startSessionRequest)

		if err != nil {
			util.LogOnError("API011", "Error starting session", err)
			log.Printf("API011: StartSessionRequest: %#v", startSessionRequest)
			return nil, gqlerror.Errorf("Error starting session")
		}

		return command.NewStartSession(*startSessionResponse), nil
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

func (r *mutationResolver) StopSession(ctx context.Context, input graph.StopSessionInput) (*graph.StopSession, error) {
	if userId := authentication.GetUserId(ctx); userId != nil {
		stopSessionRequest := command.NewStopSessionRequest(*userId, input)
		stopSessionResponse, err := r.OcpiService.StopSession(ctx, stopSessionRequest)

		if err != nil {
			util.LogOnError("API012", "Error stopping session", err)
			log.Printf("API012: StopSessionRequest: %#v", stopSessionRequest)
			return nil, gqlerror.Errorf("Error stopping session")
		}

		return command.NewStopSession(*stopSessionResponse), nil
	}

	return nil, gqlerror.Errorf("Not authenticated")
}
