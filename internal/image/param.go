package image

import (
	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/util"
	"github.com/satimoto/go-datastore/db"
)

func NewCreateImageParams(input graph.CreateImageInput) db.CreateImageParams {
	return db.CreateImageParams{
		Url:       input.URL,
		Thumbnail: util.SqlNullString(input.Thumbnail),
		Category:  db.ImageCategory(input.Category),
		Width:     util.SqlNullInt32(input.Width),
		Height:    util.SqlNullInt32(input.Height),
	}
}
