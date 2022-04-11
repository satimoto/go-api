package businessdetail

import (
	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/util"
	"github.com/satimoto/go-datastore/db"
)

func NewCreateBusinessDetailParams(input graph.CreateBusinessDetailInput) db.CreateBusinessDetailParams {
	return db.CreateBusinessDetailParams{
		Name:    input.Name,
		Website: util.SqlNullString(input.Website),
	}
}
