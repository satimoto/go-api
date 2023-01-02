package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/hex"
	"errors"
	"log"
	"time"

	"github.com/satimoto/go-api/graph"
	metrics "github.com/satimoto/go-api/internal/metric"
	"github.com/satimoto/go-api/internal/middleware"
	"github.com/satimoto/go-api/internal/param"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-lsp/lsprpc"
	"github.com/satimoto/go-lsp/pkg/lsp"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// UpdateSessionInvoice is the resolver for the updateSessionInvoice field.
func (r *mutationResolver) UpdateSessionInvoice(ctx context.Context, id int64) (*db.SessionInvoice, error) {
	if user := middleware.GetUser(ctx, r.UserRepository); user != nil {
		if sessionInvoice, err := r.SessionRepository.GetSessionInvoice(ctx, id); err == nil && user.ID == sessionInvoice.UserID {
			if sessionInvoice.IsExpired && !sessionInvoice.IsSettled {
				if !user.NodeID.Valid {
					metrics.RecordError("API048", "Error user has no node", errors.New("no node available"))
					log.Printf("API048: UserID=%v", user.ID)
					return nil, gqlerror.Errorf("No node available")
				}

				node, err := r.NodeRepository.GetNode(ctx, user.NodeID.Int64)

				if err != nil {
					metrics.RecordError("API049", "Error retrieving node", err)
					log.Printf("API049: NodeID=%#v", user.NodeID)
					return nil, gqlerror.Errorf("Error retrieving node")
				}

				lspService := lsp.NewService(node.LspAddr)

				updateSessionInvoiceRequest := &lsprpc.UpdateSessionInvoiceRequest{
					Id:     id,
					UserId: user.ID,
				}

				updateSessionInvoiceResponse, err := lspService.UpdateSessionInvoice(ctx, updateSessionInvoiceRequest)

				if err != nil {
					metrics.RecordError("API050", "Error updating session invoice", err)
					log.Printf("API050: UpdateInvoiceRequest=%#v", updateSessionInvoiceRequest)
					return nil, gqlerror.Errorf("Error updating session invoice")
				}

				updatedSessionInvoice, err := r.SessionRepository.GetSessionInvoice(ctx, id)

				if err != nil {
					metrics.RecordError("API051", "Error updating invoice", err)
					log.Printf("API051: Input=%#v", updateSessionInvoiceRequest)
					log.Printf("API051: Response=%#v", updateSessionInvoiceResponse)
					return nil, gqlerror.Errorf("Error updating invoice")
				}

				return &updatedSessionInvoice, nil

			}

			return &sessionInvoice, nil
		}

		return nil, gqlerror.Errorf("Session invoice not found")
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

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
