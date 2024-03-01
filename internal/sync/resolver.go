package sync

import (
	"context"
	"sync"

	"github.com/satimoto/go-api/internal/poi"
	"github.com/satimoto/go-datastore/pkg/db"
)

type Sync interface {
	Start(shutdownCtx context.Context, waitGroup *sync.WaitGroup)
	Sync()
}

type SyncService struct {
	PoiResolver *poi.PoiResolver
	shutdownCtx context.Context
	mutex       *sync.Mutex
	waitGroup   *sync.WaitGroup
}

func NewService(repositoryService *db.RepositoryService) Sync {
	return &SyncService{
		PoiResolver: poi.NewResolver(repositoryService),
		mutex:       &sync.Mutex{},
	}
}
