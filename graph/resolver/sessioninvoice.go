package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/hex"
	"time"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/middleware"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *queryResolver) GetSessionInvoice(ctx context.Context, id int64) (*db.SessionInvoice, error) {
	if userId := middleware.GetUserId(ctx); userId != nil {
		if s, err := r.SessionRepository.GetSessionInvoice(ctx, id); err == nil && userId == &s.UserID {
			return &s, nil
		}

		return nil, gqlerror.Errorf("Session invoice not found")
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

func (r *sessionInvoiceResolver) Signature(ctx context.Context, obj *db.SessionInvoice) (string, error) {
	return hex.EncodeToString(obj.Signature), nil
}

func (r *sessionInvoiceResolver) LastUpdated(ctx context.Context, obj *db.SessionInvoice) (string, error) {
	return obj.LastUpdated.Format(time.RFC3339), nil
}

// Session returns graph.SessioInvoicenResolver implementation.
func (r *Resolver) SessionInvoice() graph.SessionInvoiceResolver { return &sessionInvoiceResolver{r} }

type sessionInvoiceResolver struct{ *Resolver }
