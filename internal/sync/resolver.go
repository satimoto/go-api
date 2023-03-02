package sync

import (
	"context"
	"sync"

	"github.com/satimoto/go-api/internal/poi"
	"github.com/satimoto/go-datastore/pkg/db"
)

type SyncRepository interface{}

type SyncService struct {
	Repository  SyncRepository
	PoiResolver *poi.PoiResolver
	shutdownCtx context.Context
	waitGroup   *sync.WaitGroup
}

func NewService(repositoryService *db.RepositoryService) *SyncService {
	repo := SyncRepository(repositoryService)

	return &SyncService{
		Repository:  repo,
		PoiResolver: poi.NewResolver(repositoryService),
	}
}
