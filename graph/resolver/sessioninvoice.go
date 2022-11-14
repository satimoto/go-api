package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/hex"
	"time"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/middleware"
	"github.com/satimoto/go-api/internal/param"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// GetSessionInvoice is the resolver for the getSessionInvoice field.
func (r *queryResolver) GetSessionInvoice(ctx context.Context, id int64) (*db.SessionInvoice, error) {
	if userID := middleware.GetUserID(ctx); userID != nil {
		if s, err := r.SessionRepository.GetSessionInvoice(ctx, id); err == nil && *userID == s.UserID {
			return &s, nil
		}

		return nil, gqlerror.Errorf("Session invoice not found")
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

// ListSessionInvoices is the resolver for the listSessionInvoices field.
func (r *queryResolver) ListSessionInvoices(ctx context.Context, input graph.ListSessionInvoicesInput) ([]db.SessionInvoice, error) {
	if userID := middleware.GetUserID(ctx); userID != nil {
		listSessionInvoicesByUserIDParams := param.NewListSessionInvoicesByUserIDParams(*userID, input)

		if s, err := r.SessionRepository.ListSessionInvoicesByUserID(ctx, listSessionInvoicesByUserIDParams); err == nil {
			return s, nil
		}

		return nil, gqlerror.Errorf("Session invoices not found")
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

// Signature is the resolver for the signature field.
func (r *sessionInvoiceResolver) Signature(ctx context.Context, obj *db.SessionInvoice) (string, error) {
	return hex.EncodeToString(obj.Signature), nil
}

// LastUpdated is the resolver for the lastUpdated field.
func (r *sessionInvoiceResolver) LastUpdated(ctx context.Context, obj *db.SessionInvoice) (string, error) {
	return obj.LastUpdated.Format(time.RFC3339), nil
}

// SessionInvoice returns graph.SessionInvoiceResolver implementation.
func (r *Resolver) SessionInvoice() graph.SessionInvoiceResolver { return &sessionInvoiceResolver{r} }

type sessionInvoiceResolver struct{ *Resolver }
