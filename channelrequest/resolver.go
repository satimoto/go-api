package channelrequest

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type ChannelRequestRepository interface {
	CreateChannelRequest(ctx context.Context, arg db.CreateChannelRequestParams) (db.ChannelRequest, error)
	GetChannelRequestByPaymentHash(ctx context.Context, paymentHash []byte) (db.ChannelRequest, error)
	UpdateChannelRequest(ctx context.Context, arg db.UpdateChannelRequestParams) (db.ChannelRequest, error)
}

type ChannelRequestResolver struct {
	Repository ChannelRequestRepository
}

func NewResolver(repositoryService *db.RepositoryService) *ChannelRequestResolver {
	repo := ChannelRequestRepository(repositoryService)

	return &ChannelRequestResolver{
		Repository: repo,
	}
}
