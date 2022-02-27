package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/base64"

	"github.com/satimoto/go-api/authentication"
	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/user"
	"github.com/satimoto/go-api/util"
	"github.com/satimoto/go-datastore/db"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *channelRequestResolver) PaymentHash(ctx context.Context, obj *db.ChannelRequest) (string, error) {
	return base64.StdEncoding.EncodeToString(obj.PaymentHash), nil
}

func (r *channelRequestResolver) PaymentAddr(ctx context.Context, obj *db.ChannelRequest) (string, error) {
	return base64.StdEncoding.EncodeToString(obj.PaymentAddr), nil
}

func (r *channelRequestResolver) Node(ctx context.Context, obj *db.ChannelRequest) (*db.Node, error) {
	if node, err := r.NodeResolver.Repository.GetNode(ctx, obj.NodeID); err == nil {
		return &node, nil
	}

	return nil, gqlerror.Errorf("Node not found")
}

func (r *mutationResolver) CreateChannelRequest(ctx context.Context, input graph.CreateChannelRequestInput) (*db.ChannelRequest, error) {
	if userId := authentication.GetUserId(ctx); userId != nil {
		if u, err := r.UserResolver.Repository.GetUser(ctx, *userId); err == nil {
			paymentHashBytes, err := base64.StdEncoding.DecodeString(input.PaymentHash)

			if err != nil {
				return nil, gqlerror.Errorf("Error decoding paymentHash")
			}

			paymentAddrBytes, err := base64.StdEncoding.DecodeString(input.PaymentAddr)

			if err != nil {
				return nil, gqlerror.Errorf("Error decoding paymentAddr")
			}

			// TODO: Improve node selection
			// Could be by number of peers or available liquidity
			nodeId := u.NodeID.Int64

			if !u.NodeID.Valid {
				if nodes, err := r.NodeResolver.Repository.ListNodes(ctx); err == nil && len(nodes) > 0 {
					for _, node := range nodes {
						nodeId = node.ID
						break
					}

					userUpdateParams := user.NewUpdateUserParams(u)
					userUpdateParams.NodeID = util.SqlNullInt64(nodeId)

					r.UserResolver.Repository.UpdateUser(ctx, userUpdateParams)
				}
			}

			channelRequest, err := r.ChannelRequestResolver.Repository.CreateChannelRequest(ctx, db.CreateChannelRequestParams{
				UserID:      u.ID,
				NodeID:      nodeId,
				Status:      db.ChannelRequestStatusREQUESTED,
				Pubkey:      u.Pubkey,
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
	}

	return nil, gqlerror.Errorf("Not Authenticated")
}

// ChannelRequest returns graph.ChannelRequestResolver implementation.
func (r *Resolver) ChannelRequest() graph.ChannelRequestResolver { return &channelRequestResolver{r} }

type channelRequestResolver struct{ *Resolver }
