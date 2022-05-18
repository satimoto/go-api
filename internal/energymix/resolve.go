package energymix

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

type EnergyMixRepository interface {
	GetEnergyMix(ctx context.Context, id int64) (db.EnergyMix, error)
	ListEnergySources(ctx context.Context, energyMixID int64) ([]db.EnergySource, error)
	ListEnvironmentalImpacts(ctx context.Context, energyMixID int64) ([]db.EnvironmentalImpact, error)
}

type EnergyMixResolver struct {
	Repository EnergyMixRepository
}

func NewResolver(repositoryService *db.RepositoryService) *EnergyMixResolver {
	repo := EnergyMixRepository(repositoryService)
	return &EnergyMixResolver{repo}
}
