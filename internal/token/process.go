package token

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/ocpirpc/tokenrpc"
	"google.golang.org/grpc"
)

func (r *TokenResolver) CreateToken(ctx context.Context, userID int64) (*tokenrpc.TokenResponse, error) {
	ocpiRpcAddress := os.Getenv("OCPI_RPC_ADDRESS")
	conn, err := grpc.Dial(ocpiRpcAddress, grpc.WithInsecure())

	if err != nil {
		log.Printf("Error CreateToken Dial: %v", err)
		log.Printf("OCPI_RPC_ADDRESS=%v", ocpiRpcAddress)
		return nil, errors.New("Error creating user")
	}

	defer conn.Close()
	tokenClient := tokenrpc.NewTokenServiceClient(conn)
	createTokenRequest := &tokenrpc.CreateTokenRequest{
		UserId: userID,
		Type: string(db.TokenTypeOTHER),
	}
	createTokenResponse, err := tokenClient.CreateToken(ctx, createTokenRequest)

	if err != nil {
		log.Printf("Error CreateToken CreateToken: %v", err)
		log.Printf("%#v", createTokenRequest)
		return nil, errors.New("Error creating user")
	}

	return createTokenResponse, nil
}