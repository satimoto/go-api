package account

import "context"

func (r *AccountResolver) GetTaxPercentByCountry(ctx context.Context, country string, fallback float64) float64 {
	countryAccount, err := r.Repository.GetCountryAccountByCountry(ctx, country)

	if err != nil {
		return fallback
	}

	return countryAccount.TaxPercent
}