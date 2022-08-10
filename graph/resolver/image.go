package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/param"
	"github.com/satimoto/go-api/internal/util"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *mutationResolver) CreateImage(ctx context.Context, input graph.CreateImageInput) (*db.Image, error) {
	params := param.NewCreateImageParams(input)

	i, err := r.ImageRepository.CreateImage(ctx, params)

	if err != nil {
		return nil, gqlerror.Errorf("Error creating image")
	}

	return &i, nil
}

func (r *imageResolver) Thumbnail(ctx context.Context, obj *db.Image) (*string, error) {
	return util.NullString(obj.Thumbnail)
}

func (r *imageResolver) Category(ctx context.Context, obj *db.Image) (string, error) {
	return string(obj.Category), nil
}

func (r *imageResolver) Width(ctx context.Context, obj *db.Image) (*int, error) {
	return util.NullInt(obj.Width)
}

func (r *imageResolver) Height(ctx context.Context, obj *db.Image) (*int, error) {
	return util.NullInt(obj.Width)
}

// Logo returns graph.LogoResolver implementation.
func (r *Resolver) Image() graph.ImageResolver { return &imageResolver{r} }

type imageResolver struct{ *Resolver }
