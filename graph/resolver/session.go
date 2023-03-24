package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"
	"time"

	"github.com/satimoto/go-api/graph"
	metrics "github.com/satimoto/go-api/internal/metric"
	"github.com/satimoto/go-api/internal/middleware"
	"github.com/satimoto/go-api/internal/util"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/param"
	dbUtil "github.com/satimoto/go-datastore/pkg/util"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// UpdateSession is the resolver for the updateSession field.
func (r *mutationResolver) UpdateSession(ctx context.Context, input graph.UpdateSessionInput) (*db.Session, error) {
	backgroundCtx := context.Background()

	if user := middleware.GetCtxUser(ctx, r.UserRepository); user != nil {
		if session, err := r.SessionRepository.GetSessionByUid(backgroundCtx, *input.UID); err == nil {
			if !session.IsConfirmed && session.Status == db.SessionStatusTypePENDING {
				updateSessionByUidParams := param.NewUpdateSessionByUidParams(session)
				updateSessionByUidParams.IsConfirmed = true
				updateSessionByUidParams.Status = db.SessionStatusTypeACTIVE

				updatedSession, err := r.SessionRepository.UpdateSessionByUid(backgroundCtx, updateSessionByUidParams)

				if err != nil {
					metrics.RecordError("API066", "Error updating session", err)
					log.Printf("API066: Params=%#v", updateSessionByUidParams)
					return nil, gqlerror.Errorf("Error updating session")
				}

				sessionUpdateParams := param.NewCreateSessionUpdateParams(updatedSession)
				_, err = r.SessionRepository.CreateSessionUpdate(ctx, sessionUpdateParams)

				if err != nil {
					metrics.RecordError("API068", "Error creating session update", err)
					log.Printf("API068: Params=%#v", sessionUpdateParams)
				}

				session = updatedSession
			}

			return &session, nil
		}
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

// GetSession is the resolver for the getSession field.
func (r *queryResolver) GetSession(ctx context.Context, input graph.GetSessionInput) (*db.Session, error) {
	backgroundCtx := context.Background()

	if userID := middleware.GetUserID(ctx); userID != nil {
		if input.ID != nil {
			if s, err := r.SessionRepository.GetSession(backgroundCtx, *input.ID); err == nil && *userID == s.UserID {
				return &s, nil
			}
		} else if input.UID != nil {
			if s, err := r.SessionRepository.GetSessionByUid(backgroundCtx, *input.UID); err == nil && *userID == s.UserID {
				return &s, nil
			}
		} else if input.AuthorizationID != nil {
			if s, err := r.SessionRepository.GetSessionByAuthorizationID(backgroundCtx, *input.AuthorizationID); err == nil && *userID == s.UserID {
				return &s, nil
			}
		}

		return nil, gqlerror.Errorf("Session not found")
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

// ListSessions is the resolver for the listSessions field.
func (r *queryResolver) ListSessions(ctx context.Context) ([]db.Session, error) {
	backgroundCtx := context.Background()

	if userID := middleware.GetUserID(ctx); userID != nil {
		if s, err := r.SessionRepository.ListInvoicedSessionsByUserID(backgroundCtx, *userID); err == nil {
			return s, nil
		}

		return nil, gqlerror.Errorf("Sessions not found")
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

// AuthorizationID is the resolver for the authorizationId field.
func (r *sessionResolver) AuthorizationID(ctx context.Context, obj *db.Session) (*string, error) {
	return util.NullString(obj.AuthorizationID)
}

// StartDatetime is the resolver for the startDatetime field.
func (r *sessionResolver) StartDatetime(ctx context.Context, obj *db.Session) (string, error) {
	return obj.StartDatetime.Format(time.RFC3339), nil
}

// EndDatetime is the resolver for the endDatetime field.
func (r *sessionResolver) EndDatetime(ctx context.Context, obj *db.Session) (*string, error) {
	return util.NullTime(obj.EndDatetime, time.RFC3339)
}

// AuthMethod is the resolver for the authMethod field.
func (r *sessionResolver) AuthMethod(ctx context.Context, obj *db.Session) (string, error) {
	return string(obj.AuthMethod), nil
}

// Location is the resolver for the location field.
func (r *sessionResolver) Location(ctx context.Context, obj *db.Session) (*db.Location, error) {
	if location, err := r.LocationRepository.GetLocation(ctx, obj.LocationID); err == nil {
		return &location, nil
	}

	return nil, gqlerror.Errorf("Location not found")
}

// Evse is the resolver for the evse field.
func (r *sessionResolver) Evse(ctx context.Context, obj *db.Session) (*db.Evse, error) {
	if evse, err := r.LocationRepository.GetEvse(ctx, obj.EvseID); err == nil {
		return &evse, nil
	}

	return nil, gqlerror.Errorf("Evse not found")
}

// Connector is the resolver for the connector field.
func (r *sessionResolver) Connector(ctx context.Context, obj *db.Session) (*db.Connector, error) {
	if connector, err := r.LocationRepository.GetConnector(ctx, obj.ConnectorID); err == nil {
		return &connector, nil
	}

	return nil, gqlerror.Errorf("Connector not found")
}

// MeterID is the resolver for the meterId field.
func (r *sessionResolver) MeterID(ctx context.Context, obj *db.Session) (*string, error) {
	return util.NullString(obj.MeterID)
}

// InvoiceRequest is the resolver for the invoiceRequest field.
func (r *sessionResolver) InvoiceRequest(ctx context.Context, obj *db.Session) (*db.InvoiceRequest, error) {
	if obj.InvoiceRequestID.Valid {
		if invoiceRequest, err := r.InvoiceRequestRepository.GetInvoiceRequest(ctx, obj.InvoiceRequestID.Int64); err == nil {
			return &invoiceRequest, nil
		}

		return nil, gqlerror.Errorf("Invoice request not found")
	}

	return nil, nil
}

// InvoiceRequests is the resolver for the invoiceRequests field.
func (r *sessionResolver) InvoiceRequests(ctx context.Context, obj *db.Session) ([]db.InvoiceRequest, error) {
	return r.InvoiceRequestRepository.ListInvoiceRequestsBySessionID(ctx, dbUtil.SqlNullInt64(obj.ID))
}

// SessionInvoices is the resolver for the sessionInvoices field.
func (r *sessionResolver) SessionInvoices(ctx context.Context, obj *db.Session) ([]db.SessionInvoice, error) {
	return r.SessionRepository.ListSessionInvoicesBySessionID(ctx, obj.ID)
}

// SessionUpdates is the resolver for the sessionUpdates field.
func (r *sessionResolver) SessionUpdates(ctx context.Context, obj *db.Session) ([]db.SessionUpdate, error) {
	return r.SessionRepository.ListSessionUpdatesBySessionID(ctx, obj.ID)
}

// Status is the resolver for the status field.
func (r *sessionResolver) Status(ctx context.Context, obj *db.Session) (string, error) {
	return string(obj.Status), nil
}

// LastUpdated is the resolver for the lastUpdated field.
func (r *sessionResolver) LastUpdated(ctx context.Context, obj *db.Session) (string, error) {
	return obj.LastUpdated.Format(time.RFC3339), nil
}

// Session returns graph.SessionResolver implementation.
func (r *Resolver) Session() graph.SessionResolver { return &sessionResolver{r} }

type sessionResolver struct{ *Resolver }
