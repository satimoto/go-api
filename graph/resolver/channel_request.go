package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/base64"

	"github.com/lightningnetwork/lnd/lntypes"
	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-datastore/db"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *channelRequestResolver) Pubkey(ctx context.Context, obj *db.ChannelRequest) (string, error) {
	return base64.StdEncoding.EncodeToString(obj.Pubkey), nil
}

func (r *channelRequestResolver) PaymentHash(ctx context.Context, obj *db.ChannelRequest) (string, error) {
	return base64.StdEncoding.EncodeToString(obj.PaymentHash), nil
}

func (r *channelRequestResolver) PaymentAddr(ctx context.Context, obj *db.ChannelRequest) (string, error) {
	return base64.StdEncoding.EncodeToString(obj.PaymentAddr), nil
}

func (r *mutationResolver) CreateChannelRequest(ctx context.Context, input graph.CreateChannelRequestInput) (*db.ChannelRequest, error) {
	pubkeyBytes, err := base64.StdEncoding.DecodeString(input.Pubkey)

	if err != nil {
		return nil, gqlerror.Errorf("Error decoding pubkey")
	}

	preimageBytes, err := base64.StdEncoding.DecodeString(input.Preimage)

	if err != nil {
		return nil, gqlerror.Errorf("Error decoding preimage")
	}

	preimage, err := lntypes.MakePreimage(preimageBytes)

	if err != nil {
		return nil, gqlerror.Errorf("Error making payment hash")
	}

	paymentHashBytes := preimage.Hash()
	paymentAddrBytes, err := base64.StdEncoding.DecodeString(input.PaymentAddr)

	if err != nil {
		return nil, gqlerror.Errorf("Error decoding payment addr")
	}

	channelRequest, err := r.ChannelRequestResolver.Repository.CreateChannelRequest(ctx, db.CreateChannelRequestParams{
		Status:      db.ChannelRequestStatusREQUESTED,
		Pubkey:      pubkeyBytes,
		Preimage:    preimageBytes,
		PaymentHash: paymentHashBytes[:],
		PaymentAddr: paymentAddrBytes,
		AmountMsat:  int64(input.AmountMsat),
		SettledMsat: 0,
	})

	if err != nil {
		return nil, gqlerror.Errorf("Channel request already exists")
	}

	return &channelRequest, nil
}

// ChannelRequest returns graph.ChannelRequestResolver implementation.
func (r *Resolver) ChannelRequest() graph.ChannelRequestResolver { return &channelRequestResolver{r} }

type channelRequestResolver struct{ *Resolver }
