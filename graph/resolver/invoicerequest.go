package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/middleware"
	"github.com/satimoto/go-api/internal/util"
	"github.com/satimoto/go-datastore/pkg/db"
	dbUtil "github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-lsp/lsprpc"
	"github.com/satimoto/go-lsp/pkg/lsp"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *queryResolver) ListInvoiceRequests(ctx context.Context) ([]db.InvoiceRequest, error) {
	if userId := middleware.GetUserId(ctx); userId != nil {
		if invoiceRequests, err := r.InvoiceRequestRepository.ListUnsettledInvoiceRequests(ctx, *userId); err == nil {
			return invoiceRequests, nil
		}
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

func (r *mutationResolver) UpdateInvoiceRequest(ctx context.Context, input graph.UpdateInvoiceRequestInput) (*db.InvoiceRequest, error) {
	if user := middleware.GetUser(ctx, r.UserRepository); user != nil {
		if !user.NodeID.Valid {
			log.Printf("API026: Error user has no node")
			log.Printf("API026: Input=%#v", input)
			return nil, gqlerror.Errorf("No node available")
		}

		node, err := r.NodeRepository.GetNode(ctx, user.NodeID.Int64)

		if err != nil {
			dbUtil.LogOnError("API030", "Error retrieving node", err)
			log.Printf("API030: Input=%#v", input)
			return nil, gqlerror.Errorf("Error retrieving node")
		}

		// TODO: This request should be a non-blocking goroutine
		lspService := lsp.NewService(node.LspAddr)

		updateInvoiceRequest := &lsprpc.UpdateInvoiceRequest{
			Id:             input.ID,
			UserId:         user.ID,
			PaymentRequest: input.PaymentRequest,
		}

		updateInvoiceResponse, err := lspService.UpdateInvoice(ctx, updateInvoiceRequest)

		if err != nil {
			dbUtil.LogOnError("API031", "Error updating invoice", err)
			log.Printf("API031: UpdateInvoiceRequest=%#v", updateInvoiceRequest)
			return nil, gqlerror.Errorf("Error updating invoice")
		}

		invoiceRequest, err := r.InvoiceRequestRepository.GetInvoiceRequest(ctx, input.ID)

		if err != nil {
			dbUtil.LogOnError("API032", "Error updating invoice", err)
			log.Printf("API032: Input=%#v", updateInvoiceRequest)
			log.Printf("API032: Response=%#v", updateInvoiceResponse)
			return nil, gqlerror.Errorf("Error updating invoice")
		}

		return &invoiceRequest, nil
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

func (r *invoiceRequestResolver) CommissionFiat(ctx context.Context, obj *db.InvoiceRequest) (*float64, error) {
	return util.NullFloat(obj.CommissionFiat)
}

func (r *invoiceRequestResolver) CommissionMsat(ctx context.Context, obj *db.InvoiceRequest) (*int, error) {
	return util.NullInt(obj.CommissionMsat)
}

func (r *invoiceRequestResolver) PriceFiat(ctx context.Context, obj *db.InvoiceRequest) (*float64, error) {
	return util.NullFloat(obj.PriceFiat)
}

func (r *invoiceRequestResolver) PriceMsat(ctx context.Context, obj *db.InvoiceRequest) (*int, error) {
	return util.NullInt(obj.PriceMsat)
}

func (r *invoiceRequestResolver) Promotion(ctx context.Context, obj *db.InvoiceRequest) (*db.Promotion, error) {
	if promotion, err := r.PromotionRepository.GetPromotion(ctx, obj.PromotionID); err == nil {
		return &promotion, nil
	}

	return nil, gqlerror.Errorf("Promotion not found")
}

func (r *invoiceRequestResolver) TaxFiat(ctx context.Context, obj *db.InvoiceRequest) (*float64, error) {
	return util.NullFloat(obj.TaxFiat)
}

func (r *invoiceRequestResolver) TaxMsat(ctx context.Context, obj *db.InvoiceRequest) (*int, error) {
	return util.NullInt(obj.TaxMsat)
}

// Session returns graph.SessioInvoicenResolver implementation.
func (r *Resolver) InvoiceRequest() graph.InvoiceRequestResolver { return &invoiceRequestResolver{r} }

type invoiceRequestResolver struct{ *Resolver }
