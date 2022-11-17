package param

import (
	"time"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/util"
	"github.com/satimoto/go-datastore/pkg/db"
	dbUtil "github.com/satimoto/go-datastore/pkg/util"
)

func NewListLocationsByGeomParams(input graph.ListLocationsInput) db.ListLocationsByGeomParams {
	return db.ListLocationsByGeomParams{
		XMin:     util.DefaultFloat(input.XMin, 0),
		YMin:     util.DefaultFloat(input.YMin, 0),
		XMax:     util.DefaultFloat(input.XMax, 0),
		YMax:     util.DefaultFloat(input.YMax, 0),
		Interval: int32(util.DefaultInt(input.Interval, 0)),
	}
}

func NewListLocation(location db.Location) graph.ListLocation {
	return graph.ListLocation{
		UID:             location.Uid,
		CountryCode:     dbUtil.NilString(location.CountryCode),
		PartyID:         dbUtil.NilString(location.PartyID),
		Name:            location.Name.String,
		Geom:            location.Geom,
		AvailableEvses:  int(location.AvailableEvses),
		TotalEvses:      int(location.TotalEvses),
		IsRemoteCapable: location.IsRemoteCapable,
		IsRfidCapable:   location.IsRfidCapable,
		AddedDate:       location.AddedDate.Format(time.RFC3339),
	}
}
