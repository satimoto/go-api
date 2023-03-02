package poi

import (
	"net/http"

	"github.com/satimoto/go-api/internal/transportation"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/poi"
)

type PoiResolver struct {
	Repository    poi.PoiRepository
	HTTPRequester transportation.HTTPRequester
}

func NewResolver(repositoryService *db.RepositoryService) *PoiResolver {
	return &PoiResolver{
		Repository:    poi.NewRepository(repositoryService),
		HTTPRequester: &http.Client{},
	}
}
