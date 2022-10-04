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

// Thumbnail is the resolver for the thumbnail field.
func (r *imageResolver) Thumbnail(ctx context.Context, obj *db.Image) (*string, error) {
	return util.NullString(obj.Thumbnail)
}

// Category is the resolver for the category field.
func (r *imageResolver) Category(ctx context.Context, obj *db.Image) (string, error) {
	return string(obj.Category), nil
}

// Width is the resolver for the width field.
func (r *imageResolver) Width(ctx context.Context, obj *db.Image) (*int, error) {
	return util.NullInt(obj.Width)
}

// Height is the resolver for the height field.
func (r *imageResolver) Height(ctx context.Context, obj *db.Image) (*int, error) {
	return util.NullInt(obj.Width)
}

// Image returns graph.ImageResolver implementation.
func (r *Resolver) Image() graph.ImageResolver { return &imageResolver{r} }

type imageResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) CreateImage(ctx context.Context, input graph.CreateImageInput) (*db.Image, error) {
	params := param.NewCreateImageParams(input)

	i, err := r.ImageRepository.CreateImage(ctx, params)

	if err != nil {
		return nil, gqlerror.Errorf("Error creating image")
	}

	return &i, nil
}
