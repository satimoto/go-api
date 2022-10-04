package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/base64"
	"encoding/binary"
	"log"
	"math/big"
	"strconv"

	"github.com/lightningnetwork/lnd/lnwire"
	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/middleware"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/param"
	dbUtil "github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-lsp/lsprpc"
	"github.com/satimoto/go-lsp/pkg/lsp"
	"github.com/satimoto/go-lsp/pkg/util"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// PaymentHash is the resolver for the paymentHash field.
func (r *channelRequestResolver) PaymentHash(ctx context.Context, obj *db.ChannelRequest) (string, error) {
	return base64.StdEncoding.EncodeToString(obj.PaymentHash), nil
}

// PaymentAddr is the resolver for the paymentAddr field.
func (r *channelRequestResolver) PaymentAddr(ctx context.Context, obj *db.ChannelRequest) (string, error) {
	return base64.StdEncoding.EncodeToString(obj.PaymentAddr), nil
}

// AmountMsat is the resolver for the amountMsat field.
func (r *channelRequestResolver) AmountMsat(ctx context.Context, obj *db.ChannelRequest) (string, error) {
	return strconv.FormatInt(obj.AmountMsat, 10), nil
}

// Node is the resolver for the node field.
func (r *channelRequestResolver) Node(ctx context.Context, obj *db.ChannelRequest) (*db.Node, error) {
	if node, err := r.NodeRepository.GetNode(ctx, obj.NodeID); err == nil {
		return &node, nil
	}

	return nil, gqlerror.Errorf("Node not found")
}

// PendingChanID is the resolver for the pendingChanId field.
func (r *channelRequestResolver) PendingChanID(ctx context.Context, obj *db.ChannelRequest) (string, error) {
	bigInt := new(big.Int)
	bigInt.SetBytes(obj.PendingChanID)

	return bigInt.Text(10), nil
}

// Scid is the resolver for the scid field.
func (r *channelRequestResolver) Scid(ctx context.Context, obj *db.ChannelRequest) (string, error) {
	scid := binary.LittleEndian.Uint64(obj.Scid)

	return strconv.FormatUint(scid, 10), nil
}

// CreateChannelRequest is the resolver for the createChannelRequest field.
func (r *mutationResolver) CreateChannelRequest(ctx context.Context, input graph.CreateChannelRequestInput) (*db.ChannelRequest, error) {
	if user := middleware.GetUser(ctx, r.UserRepository); user != nil {
		paymentHashBytes, err := base64.StdEncoding.DecodeString(input.PaymentHash)

		if err != nil {
			return nil, gqlerror.Errorf("Error decoding paymentHash")
		}

		paymentAddrBytes, err := base64.StdEncoding.DecodeString(input.PaymentAddr)

		if err != nil {
			return nil, gqlerror.Errorf("Error decoding paymentAddr")
		}

		amountMsat, err := strconv.ParseInt(input.AmountMsat, 10, 64)

		if err != nil {
			return nil, gqlerror.Errorf("Error decoding amountMsat")
		}

		amount := int64(lnwire.MilliSatoshi(amountMsat).ToSatoshis())
		channelRequestMaxAmount := int64(dbUtil.GetEnvInt32("CHANNEL_REQUEST_MAX_AMOUNT", 200000))

		if amount > channelRequestMaxAmount {
			return nil, gqlerror.Errorf("Amount exceeds %v limit", channelRequestMaxAmount)
		}

		// TODO: Improve node selection
		// Could be by number of peers or available liquidity
		var node *db.Node

		if user.NodeID.Valid {
			if n, err := r.NodeRepository.GetNode(ctx, user.NodeID.Int64); err == nil {
				node = &n
			}
		} else {
			if nodes, err := r.NodeRepository.ListActiveNodes(ctx); err == nil && len(nodes) > 0 {
				for _, n := range nodes {
					node = &n
					break
				}
			}
		}

		if node == nil {
			return nil, gqlerror.Errorf("No node available")
		} else if !user.NodeID.Valid || user.NodeID.Int64 != node.ID {
			userUpdateParams := param.NewUpdateUserParams(*user)
			userUpdateParams.NodeID = dbUtil.SqlNullInt64(node.ID)

			r.UserRepository.UpdateUser(ctx, userUpdateParams)
		}

		// TODO: This request should be a non-blocking goroutine
		lspService := lsp.NewService(node.LspAddr)

		openChannelRequest := &lsprpc.OpenChannelRequest{
			Pubkey:     user.Pubkey,
			Amount:     amount,
			AmountMsat: amountMsat,
		}

		openChannelResponse, err := lspService.OpenChannel(ctx, openChannelRequest)

		if err != nil {
			dbUtil.LogOnError("API009", "Error allocating scid", err)
			log.Printf("API009: OpenChannelRequest=%#v", openChannelRequest)
			return nil, gqlerror.Errorf("Error requesting payment channel")
		}

		scidBytes := util.Uint64ToBytes(openChannelResponse.Scid)

		createChannelRequestParams := db.CreateChannelRequestParams{
			UserID:                    user.ID,
			NodeID:                    node.ID,
			Status:                    db.ChannelRequestStatusREQUESTED,
			Pubkey:                    user.Pubkey,
			PaymentHash:               paymentHashBytes[:],
			PaymentAddr:               paymentAddrBytes,
			Amount:                    amount,
			AmountMsat:                amountMsat,
			SettledMsat:               0,
			PendingChanID:             openChannelResponse.PendingChanId,
			Scid:                      scidBytes,
			FeeBaseMsat:               openChannelResponse.FeeBaseMsat,
			FeeProportionalMillionths: int64(openChannelResponse.FeeProportionalMillionths),
			CltvExpiryDelta:           int64(openChannelResponse.CltvExpiryDelta),
		}

		channelRequest, err := r.ChannelRequestRepository.CreateChannelRequest(ctx, createChannelRequestParams)

		if err != nil {
			dbUtil.LogOnError("API010", "Error creating channel request", err)
			log.Printf("API010: CreateChannelRequestParams=%#v", createChannelRequestParams)
			return nil, gqlerror.Errorf("Channel request already exists")
		}

		return &channelRequest, nil
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

// ChannelRequest returns graph.ChannelRequestResolver implementation.
func (r *Resolver) ChannelRequest() graph.ChannelRequestResolver { return &channelRequestResolver{r} }

type channelRequestResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *channelRequestResolver) FeeBaseMsat(ctx context.Context, obj *db.ChannelRequest) (int, error) {
	return int(obj.FeeBaseMsat), nil
}
func (r *channelRequestResolver) FeeProportionalMillionths(ctx context.Context, obj *db.ChannelRequest) (int, error) {
	return int(obj.FeeProportionalMillionths), nil
}
func (r *channelRequestResolver) CltvExpiryDelta(ctx context.Context, obj *db.ChannelRequest) (int, error) {
	return int(obj.CltvExpiryDelta), nil
}
