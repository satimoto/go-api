package token

import (
	"context"
	"errors"
	"log"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/ocpirpc"
)

func (r *TokenResolver) CreateToken(ctx context.Context, userID int64) (*ocpirpc.CreateTokenResponse, error) {
	createTokenRequest := &ocpirpc.CreateTokenRequest{
		UserId: userID,
		Type:   string(db.TokenTypeOTHER),
	}

	createTokenResponse, err := r.OcpiService.CreateToken(ctx, createTokenRequest)

	if err != nil {
		log.Printf("Error CreateToken CreateToken: %v", err)
		log.Printf("%#v", createTokenRequest)
		return nil, errors.New("Error creating user")
	}

	return createTokenResponse, nil
}
