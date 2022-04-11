package location

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type LocationRepository interface {
	GetLocationByUid(ctx context.Context, uid string) (db.Location, error)
	ListEvses(ctx context.Context, locationID int64) ([]db.Evse, error)
	ListFacilities(ctx context.Context) ([]db.Facility, error)
	ListLocationsByGeom(ctx context.Context, arg db.ListLocationsByGeomParams) ([]db.Location, error)
	ListLocationDirections(ctx context.Context, locationID int64) ([]db.DisplayText, error)
	ListLocationFacilities(ctx context.Context, locationID int64) ([]db.Facility, error)
	ListLocationImages(ctx context.Context, locationID int64) ([]db.Image, error)
	ListLocations(ctx context.Context) ([]db.Location, error)
	ListRelatedLocations(ctx context.Context, locationID int64) ([]db.GeoLocation, error)
}

type LocationResolver struct {
	Repository LocationRepository
}

func NewResolver(repositoryService *db.RepositoryService) *LocationResolver {
	repo := LocationRepository(repositoryService)
	return &LocationResolver{repo}
}
