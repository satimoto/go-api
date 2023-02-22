package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-datastore/pkg/db"
)

// ListSessionUpdates is the resolver for the listSessionUpdates field.
func (r *queryResolver) ListSessionUpdates(ctx context.Context, id int64) ([]db.SessionUpdate, error) {	
	return r.SessionRepository.ListSessionUpdatesBySessionID(ctx, id)
}

// Status is the resolver for the status field.
func (r *sessionUpdateResolver) Status(ctx context.Context, obj *db.SessionUpdate) (string, error) {
	return string(obj.Status), nil
}

// LastUpdated is the resolver for the lastUpdated field.
func (r *sessionUpdateResolver) LastUpdated(ctx context.Context, obj *db.SessionUpdate) (string, error) {
	return obj.LastUpdated.Format(time.RFC3339), nil
}

// SessionUpdate returns graph.SessionUpdateResolver implementation.
func (r *Resolver) SessionUpdate() graph.SessionUpdateResolver { return &sessionUpdateResolver{r} }

type sessionUpdateResolver struct{ *Resolver }
