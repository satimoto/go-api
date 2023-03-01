package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"log"

	"github.com/satimoto/go-api/graph"
	metrics "github.com/satimoto/go-api/internal/metric"
	"github.com/satimoto/go-api/internal/middleware"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-lsp/lsprpc"
	"github.com/satimoto/go-lsp/pkg/lsp"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Addr is the resolver for the addr field.
func (r *nodeResolver) Addr(ctx context.Context, obj *db.Node) (string, error) {
	return obj.NodeAddr, nil
}

// ListChannels is the resolver for the listChannels field.
func (r *queryResolver) ListChannels(ctx context.Context) ([]graph.Channel, error) {
	backgroundCtx := context.Background()

	if user := middleware.GetCtxUser(ctx, r.UserRepository); user != nil {
		if !user.NodeID.Valid {
			metrics.RecordError("API052", "Error user has no node", errors.New("no node available"))
			log.Printf("API052: UserID=%v", user.ID)
			return nil, gqlerror.Errorf("No node available")
		}

		node, err := r.NodeRepository.GetNode(backgroundCtx, user.NodeID.Int64)

		if err != nil {
			metrics.RecordError("API053", "Error retrieving node", err)
			log.Printf("API053: NodeID=%#v", user.NodeID)
			return nil, gqlerror.Errorf("Error retrieving node")
		}

		// TODO: This request should be a non-blocking goroutine
		lspService := lsp.NewService(node.LspAddr)
		listChannels, err := lspService.ListChannels(backgroundCtx, &lsprpc.ListChannelsRequest{})

		if err != nil {
			metrics.RecordError("API054", "Error listing channels", err)
			log.Printf("API054: NodeID=%#v", user.NodeID)
			return nil, gqlerror.Errorf("Error listing channels")
		}

		var channels []graph.Channel

		for _, channelId := range listChannels.ChannelIds {
			channels = append(channels, graph.Channel{
				ChannelID: channelId,
			})
		}

		return channels, nil
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

// Node returns graph.NodeResolver implementation.
func (r *Resolver) Node() graph.NodeResolver { return &nodeResolver{r} }

type nodeResolver struct{ *Resolver }
