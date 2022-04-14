package businessdetail

import (
	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-datastore/util"
)

func NewCreateBusinessDetailParams(input graph.CreateBusinessDetailInput) db.CreateBusinessDetailParams {
	return db.CreateBusinessDetailParams{
		Name:    input.Name,
		Website: util.SqlNullString(input.Website),
	}
}
