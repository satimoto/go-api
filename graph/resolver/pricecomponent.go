package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/middleware"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ferp/pkg/rate"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Type is the resolver for the type field.
func (r *priceComponentResolver) Type(ctx context.Context, obj *db.PriceComponent) (string, error) {
	return string(obj.Type), nil
}

// PriceMsat is the resolver for the priceMsat field.
func (r *priceComponentResolver) PriceMsat(ctx context.Context, obj *db.PriceComponent) (int, error) {
	currencyRate, err := r.FerpService.GetRate(obj.Currency)

	if err != nil {
		return 0, gqlerror.Errorf("Error retrieving exchange rate")
	}

	rateMsat := float64(currencyRate.RateMsat)
	priceMsat := int(rateMsat * obj.Price)

	return priceMsat, nil
}

// CommissionMsat is the resolver for the commissionMsat field.
func (r *priceComponentResolver) CommissionMsat(ctx context.Context, obj *db.PriceComponent) (int, error) {
	currencyRate, err := r.FerpService.GetRate(obj.Currency)

	if err != nil {
		return 0, gqlerror.Errorf("Error retrieving exchange rate")
	}

	commissionMsat, err := r.calculateCommissionMsat(ctx, currencyRate, obj)

	if err != nil {
		return 0, gqlerror.Errorf("Error calculating user commission")
	}

	return *commissionMsat, nil
}

// TaxMsat is the resolver for the taxMsat field.
func (r *priceComponentResolver) TaxMsat(ctx context.Context, obj *db.PriceComponent) (*int, error) {
	taxPercent, err := r.calculateTaxPercent(ctx)

	if err != nil {
		// Error retrieving tax percent
		return nil, nil
	}

	currencyRate, err := r.FerpService.GetRate(obj.Currency)

	if err != nil {
		// Error retrieving exchange rate
		return nil, nil
	}

	commissionMsat, err := r.calculateCommissionMsat(ctx, currencyRate, obj)

	if err != nil {
		// Error calculating commission
		return nil, nil
	}

	taxMsat := int((float64(*commissionMsat) / 100) * *taxPercent)

	return &taxMsat, nil
}

// PriceComponent returns graph.PriceComponentResolver implementation.
func (r *Resolver) PriceComponent() graph.PriceComponentResolver { return &priceComponentResolver{r} }

type priceComponentResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
type OperationalContextVariable map[string]interface{}

func (r *priceComponentResolver) calculateCommissionMsat(ctx context.Context, currencyRate *rate.CurrencyRate, obj *db.PriceComponent) (*int, error) {
	user := middleware.GetUser(ctx, r.Resolver.UserRepository)

	if user == nil {
		// Error retrieving user
		return nil, gqlerror.Errorf("Error retrieving user commission")
	}

	rateMsat := float64(currencyRate.RateMsat)
	priceMsat := rateMsat * obj.Price
	commissionMsat := int((priceMsat / 100) * user.CommissionPercent)

	return &commissionMsat, nil
}
