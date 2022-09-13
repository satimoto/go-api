package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *queryResolver) GetRate(ctx context.Context, currency string) (*graph.Rate, error) {
	currencyRate, err := r.FerpService.GetRate(currency)

	if err != nil {
		util.LogOnError("API015", "Error retrieving rate", err)
		log.Printf("API015: Currency=%v", currency)
		return nil, gqlerror.Errorf("Error retrieving rate")
	}

	return &graph.Rate{
		Rate:  strconv.FormatInt(currencyRate.Rate, 10),
		RateMsat: strconv.FormatInt(currencyRate.RateMsat, 10),
		LastUpdated: currencyRate.LastUpdated.Format(time.RFC3339),
	}, nil
}
