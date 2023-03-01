package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"log"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/command"
	metrics "github.com/satimoto/go-api/internal/metric"
	"github.com/satimoto/go-api/internal/middleware"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// StartSession is the resolver for the startSession field.
func (r *mutationResolver) StartSession(ctx context.Context, input graph.StartSessionInput) (*graph.StartSession, error) {
	backgroundCtx := context.Background()

	if user := middleware.GetCtxUser(ctx, r.UserRepository); user != nil {
		if !user.DeviceToken.Valid {
			metrics.RecordError("API046", "Error starting session", errors.New("notifications not enabled"))
			log.Printf("API046: UserID: %#v", user.ID)
			return nil, gqlerror.Errorf("Notifications not enabled")
		}

		startSessionRequest := command.NewStartSessionRequest(user.ID, input)
		startSessionResponse, err := r.OcpiService.StartSession(backgroundCtx, startSessionRequest)

		if err != nil {
			metrics.RecordError("API011", "Error starting session", err)
			log.Printf("API011: StartSessionRequest: %#v", startSessionRequest)
			return nil, gqlerror.Errorf("Error starting session")
		}

		return command.NewStartSession(*startSessionResponse), nil
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

// StopSession is the resolver for the stopSession field.
func (r *mutationResolver) StopSession(ctx context.Context, input graph.StopSessionInput) (*graph.StopSession, error) {
	backgroundCtx := context.Background()

	if userID := middleware.GetUserID(ctx); userID != nil {
		stopSessionRequest := command.NewStopSessionRequest(*userID, input)
		stopSessionResponse, err := r.OcpiService.StopSession(backgroundCtx, stopSessionRequest)

		if err != nil {
			metrics.RecordError("API012", "Error stopping session", err)
			log.Printf("API012: StopSessionRequest: %#v", stopSessionRequest)
			return nil, gqlerror.Errorf("Error stopping session")
		}

		return command.NewStopSession(*stopSessionResponse), nil
	}

	return nil, gqlerror.Errorf("Not authenticated")
}
