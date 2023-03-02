package param

import (
	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/util"
	"github.com/satimoto/go-datastore/pkg/db"
)

func NewListPoisByGeomParams(input graph.ListPoisInput) db.ListPoisByGeomParams {
	return db.ListPoisByGeomParams{
		XMin:            util.DefaultFloat(input.XMin, 0),
		YMin:            util.DefaultFloat(input.YMin, 0),
		XMax:            util.DefaultFloat(input.XMax, 0),
		YMax:            util.DefaultFloat(input.YMax, 0),
	}
}
