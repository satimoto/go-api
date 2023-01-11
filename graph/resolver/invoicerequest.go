package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"log"

	"github.com/satimoto/go-api/graph"
	metrics "github.com/satimoto/go-api/internal/metric"
	"github.com/satimoto/go-api/internal/middleware"
	"github.com/satimoto/go-api/internal/util"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-lsp/lsprpc"
	"github.com/satimoto/go-lsp/pkg/lsp"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Promotion is the resolver for the promotion field.
func (r *invoiceRequestResolver) Promotion(ctx context.Context, obj *db.InvoiceRequest) (*db.Promotion, error) {
	if promotion, err := r.PromotionRepository.GetPromotion(ctx, obj.PromotionID); err == nil {
		return &promotion, nil
	}

	return nil, gqlerror.Errorf("Promotion not found")
}

// PriceFiat is the resolver for the priceFiat field.
func (r *invoiceRequestResolver) PriceFiat(ctx context.Context, obj *db.InvoiceRequest) (*float64, error) {
	return util.NullFloat(obj.PriceFiat)
}

// PriceMsat is the resolver for the priceMsat field.
func (r *invoiceRequestResolver) PriceMsat(ctx context.Context, obj *db.InvoiceRequest) (*int, error) {
	return util.NullInt(obj.PriceMsat)
}

// CommissionFiat is the resolver for the commissionFiat field.
func (r *invoiceRequestResolver) CommissionFiat(ctx context.Context, obj *db.InvoiceRequest) (*float64, error) {
	return util.NullFloat(obj.CommissionFiat)
}

// CommissionMsat is the resolver for the commissionMsat field.
func (r *invoiceRequestResolver) CommissionMsat(ctx context.Context, obj *db.InvoiceRequest) (*int, error) {
	return util.NullInt(obj.CommissionMsat)
}

// TaxFiat is the resolver for the taxFiat field.
func (r *invoiceRequestResolver) TaxFiat(ctx context.Context, obj *db.InvoiceRequest) (*float64, error) {
	return util.NullFloat(obj.TaxFiat)
}

// TaxMsat is the resolver for the taxMsat field.
func (r *invoiceRequestResolver) TaxMsat(ctx context.Context, obj *db.InvoiceRequest) (*int, error) {
	return util.NullInt(obj.TaxMsat)
}

// PaymentRequest is the resolver for the paymentRequest field.
func (r *invoiceRequestResolver) PaymentRequest(ctx context.Context, obj *db.InvoiceRequest) (*string, error) {
	return util.NullString(obj.PaymentRequest)
}

// UpdateInvoiceRequest is the resolver for the updateInvoiceRequest field.
func (r *mutationResolver) UpdateInvoiceRequest(reqCtx context.Context, input graph.UpdateInvoiceRequestInput) (*db.InvoiceRequest, error) {
	ctx := context.Background()
	
	if user := middleware.GetUser(reqCtx, r.UserRepository); user != nil {
		if !user.NodeID.Valid {
			metrics.RecordError("API026", "Error user has no node", errors.New("no node available"))
			log.Printf("API026: Input=%#v", input)
			return nil, gqlerror.Errorf("No node available")
		}

		node, err := r.NodeRepository.GetNode(ctx, user.NodeID.Int64)

		if err != nil {
			metrics.RecordError("API030", "Error retrieving node", err)
			log.Printf("API030: Input=%#v", input)
			return nil, gqlerror.Errorf("Error retrieving node")
		}

		// TODO: This request should be a non-blocking goroutine
		lspService := lsp.NewService(node.LspAddr)

		updateInvoiceRequest := &lsprpc.UpdateInvoiceRequestRequest{
			Id:             input.ID,
			UserId:         user.ID,
			PaymentRequest: input.PaymentRequest,
		}

		updateInvoiceResponse, err := lspService.UpdateInvoiceRequest(ctx, updateInvoiceRequest)

		if err != nil {
			metrics.RecordError("API031", "Error updating invoice", err)
			log.Printf("API031: UpdateInvoiceRequest=%#v", updateInvoiceRequest)
			return nil, gqlerror.Errorf("Error updating invoice")
		}

		invoiceRequest, err := r.InvoiceRequestRepository.GetInvoiceRequest(ctx, input.ID)

		if err != nil {
			metrics.RecordError("API032", "Error updating invoice", err)
			log.Printf("API032: Input=%#v", updateInvoiceRequest)
			log.Printf("API032: Response=%#v", updateInvoiceResponse)
			return nil, gqlerror.Errorf("Error updating invoice")
		}

		return &invoiceRequest, nil
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

// ListInvoiceRequests is the resolver for the listInvoiceRequests field.
func (r *queryResolver) ListInvoiceRequests(reqCtx context.Context) ([]db.InvoiceRequest, error) {
	ctx := context.Background()
	
	if userID := middleware.GetUserID(reqCtx); userID != nil {
		if invoiceRequests, err := r.InvoiceRequestRepository.ListUnsettledInvoiceRequests(ctx, *userID); err == nil {
			return invoiceRequests, nil
		}
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

// InvoiceRequest returns graph.InvoiceRequestResolver implementation.
func (r *Resolver) InvoiceRequest() graph.InvoiceRequestResolver { return &invoiceRequestResolver{r} }

type invoiceRequestResolver struct{ *Resolver }
