package node

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type NodeRepository interface {
	CreateNode(ctx context.Context, arg db.CreateNodeParams) (db.Node, error)
	GetNode(ctx context.Context, id int64) (db.Node, error)
}

type NodeResolver struct {
	Repository NodeRepository
}

func NewResolver(repositoryService *db.RepositoryService) *NodeResolver {
	repo := NodeRepository(repositoryService)
	return &NodeResolver{repo}
}