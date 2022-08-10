package param

import (
	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/util"
	"github.com/satimoto/go-datastore/pkg/db"
)

func NewListLocationsByGeomParams(input graph.ListLocationsInput) db.ListLocationsByGeomParams {
	return db.ListLocationsByGeomParams{
		XMin: input.XMin,
		YMin: input.YMin,
		XMax: input.XMax,
		YMax: input.YMax,
		LastUpdate: util.DefaultString(input.LastUpdate, ""),
	}
}

func NewListLocation(location db.Location) graph.ListLocation {
	return graph.ListLocation{
		UID:             location.Uid,
		Name:            location.Name.String,
		Geom:            location.Geom,
		AvailableEvses:  int(location.AvailableEvses),
		TotalEvses:      int(location.TotalEvses),
		IsRemoteCapable: location.IsRemoteCapable,
		IsRfidCapable:   location.IsRfidCapable,
	}
}
