package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/param"
	"github.com/satimoto/go-api/internal/util"
	"github.com/satimoto/go-datastore/pkg/db"
	dbUtil "github.com/satimoto/go-datastore/pkg/util"
	"github.com/vektah/gqlparser/v2/gqlerror"
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) CreateBusinessDetail(ctx context.Context, input graph.CreateBusinessDetailInput) (*db.BusinessDetail, error) {
	params := param.NewCreateBusinessDetailParams(input)

	if input.Logo != nil {
		if logo, err := r.CreateImage(ctx, *input.Logo); err == nil {
			params.LogoID = dbUtil.SqlNullInt64(logo.ID)
		}
	}

	businessDetail, err := r.BusinessDetailRepository.CreateBusinessDetail(ctx, params)

	if err != nil {
		return nil, gqlerror.Errorf("Error creating business detail")
	}

	return &businessDetail, nil
}
