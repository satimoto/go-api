package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/base64"
	"log"
	"strconv"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/authentication"
	"github.com/satimoto/go-api/internal/param"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *channelRequestResolver) PaymentHash(ctx context.Context, obj *db.ChannelRequest) (string, error) {
	return base64.StdEncoding.EncodeToString(obj.PaymentHash), nil
}

func (r *channelRequestResolver) PaymentAddr(ctx context.Context, obj *db.ChannelRequest) (string, error) {
	return base64.StdEncoding.EncodeToString(obj.PaymentAddr), nil
}

func (r *channelRequestResolver) AmountMsat(ctx context.Context, obj *db.ChannelRequest) (string, error) {
	return strconv.FormatInt(obj.AmountMsat, 10), nil
}

func (r *channelRequestResolver) Node(ctx context.Context, obj *db.ChannelRequest) (*db.Node, error) {
	if node, err := r.NodeRepository.GetNode(ctx, obj.NodeID); err == nil {
		return &node, nil
	}

	return nil, gqlerror.Errorf("Node not found")
}

func (r *mutationResolver) CreateChannelRequest(ctx context.Context, input graph.CreateChannelRequestInput) (*db.ChannelRequest, error) {
	if userId := authentication.GetUserId(ctx); userId != nil {
		if u, err := r.UserRepository.GetUser(ctx, *userId); err == nil {
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

			// TODO: Improve node selection
			// Could be by number of peers or available liquidity
			var node *db.Node

			if u.NodeID.Valid {
				if n, err := r.NodeRepository.GetNode(ctx, u.NodeID.Int64); err == nil {
					node = &n
				}
			} else {
				if nodes, err := r.NodeRepository.ListNodes(ctx); err == nil && len(nodes) > 0 {
					for _, n := range nodes {
						node = &n
						break
					}
				}
			}

			if node == nil {
				return nil, gqlerror.Errorf("No node available")
			} else if !u.NodeID.Valid || u.NodeID.Int64 != node.ID {
				userUpdateParams := param.NewUpdateUserParams(u)
				userUpdateParams.NodeID = util.SqlNullInt64(node.ID)

				r.UserRepository.UpdateUser(ctx, userUpdateParams)
			}

			channelRequest, err := r.ChannelRequestRepository.CreateChannelRequest(ctx, db.CreateChannelRequestParams{
				UserID:      u.ID,
				NodeID:      node.ID,
				Status:      db.ChannelRequestStatusREQUESTED,
				Pubkey:      u.Pubkey,
				PaymentHash: paymentHashBytes[:],
				PaymentAddr: paymentAddrBytes,
				AmountMsat:  amountMsat,
				SettledMsat: 0,
			})

			if err != nil {
				log.Printf("CreateChannelRequest error: %v", err)
				return nil, gqlerror.Errorf("Channel request already exists")
			}

			return &channelRequest, nil
		}
	}

	return nil, gqlerror.Errorf("Not Authenticated")
}

// ChannelRequest returns graph.ChannelRequestResolver implementation.
func (r *Resolver) ChannelRequest() graph.ChannelRequestResolver { return &channelRequestResolver{r} }

type channelRequestResolver struct{ *Resolver }
