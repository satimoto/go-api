package token

import (
	"context"
	"errors"
	"log"

	metrics "github.com/satimoto/go-api/internal/metric"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi/ocpirpc"
)

func (r *TokenResolver) CreateToken(ctx context.Context, userID int64) (*ocpirpc.CreateTokenResponse, error) {
	createTokenRequest := &ocpirpc.CreateTokenRequest{
		UserId:    userID,
		Type:      string(db.TokenTypeOTHER),
		Whitelist: string(db.TokenWhitelistTypeNEVER),
	}

	createTokenResponse, err := r.OcpiService.CreateToken(ctx, createTokenRequest)

	if err != nil {
		metrics.RecordError("API025", "Error creating user", err)
		log.Printf("API025: CreateTokenRequest=%#v", createTokenRequest)
		return nil, errors.New("Error creating user")
	}

	return createTokenResponse, nil
}
