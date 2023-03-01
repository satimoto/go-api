package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"
	"time"

	"github.com/satimoto/go-api/graph"
	metrics "github.com/satimoto/go-api/internal/metric"
	"github.com/satimoto/go-api/internal/middleware"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// CreateReferral is the resolver for the createReferral field.
func (r *mutationResolver) CreateReferral(ctx context.Context, input graph.CreateReferralInput) (*graph.ResultID, error) {
	backgroundCtx := context.Background()
	user, err := r.UserRepository.GetUserByReferralCode(backgroundCtx, util.SqlNullString(input.Referrer))

	if err != nil {
		metrics.RecordError("API033", "Error retrieving user", err)
		log.Printf("API033: Referrer: %v", input.Referrer)
		return nil, gqlerror.Errorf("Error creating referral")
	}

	promotion, err := r.PromotionRepository.GetPromotionByCode(backgroundCtx, input.Code)

	if err != nil {
		metrics.RecordError("API034", "Error retrieving promotion", err)
		log.Printf("API034: Code: %v", input.Code)
		return nil, gqlerror.Errorf("Error creating referral")
	}

	ipAddress := middleware.GetIPAddress(ctx)

	if ipAddress == nil {
		metrics.RecordError("API035", "Error ip address not found", err)
		log.Printf("API035: Input: %#v", input)
		return nil, gqlerror.Errorf("Error ip address not found")
	}

	createReferralParams := db.CreateReferralParams{
		UserID:      user.ID,
		PromotionID: promotion.ID,
		IpAddress:   *ipAddress,
		LastUpdated: time.Now(),
	}

	referrer, err := r.ReferralRepository.CreateReferral(backgroundCtx, createReferralParams)

	if err != nil {
		metrics.RecordError("API036", "Error creating referral", err)
		log.Printf("API036: Params: %#v", createReferralParams)
		return nil, gqlerror.Errorf("Error creating referral")
	}

	return &graph.ResultID{ID: referrer.ID}, nil
}
