package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/satimoto/go-api/graph"
	metrics "github.com/satimoto/go-api/internal/metric"
	"github.com/satimoto/go-api/internal/middleware"
	"github.com/satimoto/go-api/internal/param"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// UpdateTariff is the resolver for the updateTariff field.
func (r *mutationResolver) UpdateTariff(ctx context.Context, input graph.UpdateTariffInput) (*graph.ResultOk, error) {
	if user := middleware.GetUser(ctx, r.UserRepository); user != nil && user.IsAdmin {
		updateTariffCapabilitiesParams := param.NewUpdateTariffCapabilitiesParams(input)

		err := r.TariffRepository.UpdateTariffCapabilities(ctx, updateTariffCapabilitiesParams)

		if err != nil {
			metrics.RecordError("API029", "Error updating tariff", err)
			log.Printf("API029: Params=%#v", updateTariffCapabilitiesParams)
			return nil, gqlerror.Errorf("Error updating tariff")
		}

		return &graph.ResultOk{Ok: true}, nil
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

// GetTariff is the resolver for the getTariff field.
func (r *queryResolver) GetTariff(ctx context.Context, input graph.GetTariffInput) (*db.Tariff, error) {
	if userID := middleware.GetUserID(ctx); userID != nil {
		if input.ID != nil {
			if t, err := r.TariffRepository.GetTariff(ctx, *input.ID); err == nil {
				return &t, nil
			}
		} else if input.UID != nil {
			if t, err := r.TariffRepository.GetTariffByUid(ctx, *input.UID); err == nil {
				return &t, nil
			}
		}

		return nil, gqlerror.Errorf("Tariff not found")
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

// CurrencyRate is the resolver for the currencyRate field.
func (r *tariffResolver) CurrencyRate(ctx context.Context, obj *db.Tariff) (int, error) {
	currencyRate, err := r.FerpService.GetRate(obj.Currency)

	if err != nil {
		return 0, gqlerror.Errorf("Error retrieving exchange rate")
	}

	return int(currencyRate.Rate), nil
}

// CurrencyRateMsat is the resolver for the currencyRateMsat field.
func (r *tariffResolver) CurrencyRateMsat(ctx context.Context, obj *db.Tariff) (int, error) {
	currencyRate, err := r.FerpService.GetRate(obj.Currency)

	if err != nil {
		return 0, gqlerror.Errorf("Error retrieving exchange rate")
	}

	return int(currencyRate.RateMsat), nil
}

// CommissionPercent is the resolver for the commissionPercent field.
func (r *tariffResolver) CommissionPercent(ctx context.Context, obj *db.Tariff) (float64, error) {
	user := middleware.GetUser(ctx, r.UserRepository)

	if user == nil {
		return 0, gqlerror.Errorf("Error retrieving user commission")
	}

	return user.CommissionPercent, nil
}

// TaxPercent is the resolver for the taxPercent field.
func (r *tariffResolver) TaxPercent(ctx context.Context, obj *db.Tariff) (*float64, error) {
	taxPercent, err := r.calculateTaxPercent(ctx)

	if err != nil {
		return nil, nil
	}

	return taxPercent, nil
}

// Elements is the resolver for the elements field.
func (r *tariffResolver) Elements(ctx context.Context, obj *db.Tariff) ([]graph.TariffElement, error) {
	list := []graph.TariffElement{}

	if elements, err := r.TariffRepository.ListElements(ctx, obj.ID); err == nil {
		for _, element := range elements {
			tariffElement := graph.TariffElement{}

			if priceComponents, err := r.TariffRepository.ListPriceComponents(ctx, element.ID); err == nil {
				tariffElement.PriceComponents = priceComponents
			}

			if element.ElementRestrictionID.Valid {
				if restriction, err := r.TariffRepository.GetElementRestriction(ctx, element.ElementRestrictionID.Int64); err == nil {
					tariffElement.Restrictions = &restriction
				}
			}

			list = append(list, tariffElement)
		}
	}

	return list, nil
}

// EnergyMix is the resolver for the energyMix field.
func (r *tariffResolver) EnergyMix(ctx context.Context, obj *db.Tariff) (*db.EnergyMix, error) {
	if obj.EnergyMixID.Valid {
		if energyMix, err := r.EnergyMixRepository.GetEnergyMix(ctx, obj.EnergyMixID.Int64); err == nil {
			return &energyMix, nil
		}
	}

	return nil, nil
}

// Tariff returns graph.TariffResolver implementation.
func (r *Resolver) Tariff() graph.TariffResolver { return &tariffResolver{r} }

type tariffResolver struct{ *Resolver }
