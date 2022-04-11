package evse

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type EvseRepository interface {
	ListConnectors(ctx context.Context, evseID int64) ([]db.Connector, error)
	ListEvseCapabilities(ctx context.Context, evseID int64) ([]db.Capability, error)
	ListEvseDirections(ctx context.Context, evseID int64) ([]db.DisplayText, error)
	ListEvseImages(ctx context.Context, evseID int64) ([]db.Image, error)
	ListEvseParkingRestrictions(ctx context.Context, evseID int64) ([]db.ParkingRestriction, error)
	ListStatusSchedules(ctx context.Context, evseID int64) ([]db.StatusSchedule, error)
}

type EvseResolver struct {
	Repository EvseRepository
}

func NewResolver(repositoryService *db.RepositoryService) *EvseResolver {
	repo := EvseRepository(repositoryService)
	return &EvseResolver{repo}
}
