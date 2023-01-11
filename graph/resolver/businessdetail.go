package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/util"
	"github.com/satimoto/go-datastore/pkg/db"
)

// Website is the resolver for the website field.
func (r *businessDetailResolver) Website(ctx context.Context, obj *db.BusinessDetail) (*string, error) {
	return util.NullString(obj.Website)
}

// Logo is the resolver for the logo field.
func (r *businessDetailResolver) Logo(ctx context.Context, obj *db.BusinessDetail) (*db.Image, error) {
	if obj.LogoID.Valid {
		if image, err := r.ImageRepository.GetImage(ctx, obj.LogoID.Int64); err == nil {
			return &image, nil
		}
	}

	return nil, nil
}

// BusinessDetail returns graph.BusinessDetailResolver implementation.
func (r *Resolver) BusinessDetail() graph.BusinessDetailResolver { return &businessDetailResolver{r} }

type businessDetailResolver struct{ *Resolver }
