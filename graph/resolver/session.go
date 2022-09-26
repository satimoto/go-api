package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/middleware"
	"github.com/satimoto/go-api/internal/util"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *queryResolver) GetSession(ctx context.Context, input graph.GetSessionInput) (*db.Session, error) {
	if userId := middleware.GetUserId(ctx); userId != nil {
		if input.ID != nil {
			if s, err := r.SessionRepository.GetSession(ctx, *input.ID); err == nil && userId == &s.UserID {
				return &s, nil
			}
		} else if input.UID != nil {
			if s, err := r.SessionRepository.GetSessionByUid(ctx, *input.UID); err == nil && userId == &s.UserID {
				return &s, nil
			}
		} else if input.AuthorizationID != nil {
			if s, err := r.SessionRepository.GetSessionByAuthorizationID(ctx, *input.AuthorizationID); err == nil && userId == &s.UserID {
				return &s, nil
			}
		}

		return nil, gqlerror.Errorf("Session not found")
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

func (r *sessionResolver) AuthorizationID(ctx context.Context, obj *db.Session) (*string, error) {
	return util.NullString(obj.AuthorizationID)
}

func (r *sessionResolver) AuthMethod(ctx context.Context, obj *db.Session) (string, error) {
	return string(obj.AuthMethod), nil
}

func (r *sessionResolver) StartDatetime(ctx context.Context, obj *db.Session) (string, error) {
	return obj.StartDatetime.Format(time.RFC3339), nil
}

func (r *sessionResolver) EndDatetime(ctx context.Context, obj *db.Session) (*string, error) {
	return util.NullTime(obj.EndDatetime, time.RFC3339)
}

func (r *sessionResolver) Location(ctx context.Context, obj *db.Session) (*db.Location, error) {
	if location, err := r.LocationRepository.GetLocation(ctx, obj.LocationID); err == nil {
		return &location, nil
	}

	return nil, gqlerror.Errorf("Location not found")
}

func (r *sessionResolver) Evse(ctx context.Context, obj *db.Session) (*db.Evse, error) {
	if evse, err := r.LocationRepository.GetEvse(ctx, obj.EvseID); err == nil {
		return &evse, nil
	}

	return nil, gqlerror.Errorf("Evse not found")
}

func (r *sessionResolver) Connector(ctx context.Context, obj *db.Session) (*db.Connector, error) {
	if connector, err := r.LocationRepository.GetConnector(ctx, obj.ConnectorID); err == nil {
		return &connector, nil
	}

	return nil, gqlerror.Errorf("Connector not found")
}

func (r *sessionResolver) MeterID(ctx context.Context, obj *db.Session) (*string, error) {
	return util.NullString(obj.MeterID)
}

func (r *sessionResolver) SessionInvoices(ctx context.Context, obj *db.Session) ([]db.SessionInvoice, error) {
	return r.SessionRepository.ListSessionInvoices(ctx, obj.ID)
}

func (r *sessionResolver) Status(ctx context.Context, obj *db.Session) (string, error) {
	return string(obj.Status), nil
}

func (r *sessionResolver) LastUpdated(ctx context.Context, obj *db.Session) (string, error) {
	return obj.LastUpdated.Format(time.RFC3339), nil
}

// Session returns graph.SessionResolver implementation.
func (r *Resolver) Session() graph.SessionResolver { return &sessionResolver{r} }

type sessionResolver struct{ *Resolver }
