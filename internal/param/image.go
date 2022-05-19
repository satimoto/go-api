package param

import (
	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
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
