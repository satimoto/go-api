package image

import (
	"github.com/satimoto/go-datastore/pkg/db"
)

type ImageRepository interface {
}

type ImageResolver struct {
	Repository ImageRepository
}

func NewResolver(repositoryService *db.RepositoryService) *ImageResolver {
	repo := ImageRepository(repositoryService)

	return &ImageResolver{
		Repository: repo,
	}
}
