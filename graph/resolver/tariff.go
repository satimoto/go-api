package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/authentication"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *queryResolver) GetTariff(ctx context.Context, input graph.GetTariffInput) (*db.Tariff, error) {
	if userId := authentication.GetUserId(ctx); userId != nil {
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

func (r *tariffResolver) EnergyMix(ctx context.Context, obj *db.Tariff) (*db.EnergyMix, error) {
	if obj.EnergyMixID.Valid {
		if energyMix, err := r.EnergyMixRepository.GetEnergyMix(ctx, obj.EnergyMixID.Int64); err == nil {
			return &energyMix, nil
		}
	}

	return nil, nil
}

// Tariff returns graph.TariffResolver implementation.
func (r *Resolver) Tariff() graph.TariffResolver {
	return &tariffResolver{r}
}

type tariffResolver struct{ *Resolver }

