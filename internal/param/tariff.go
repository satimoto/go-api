package param

import (
	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/util"
	"github.com/satimoto/go-datastore/pkg/db"
)

func NewUpdateTariffCapabilitiesParams(input graph.UpdateTariffInput) db.UpdateTariffCapabilitiesParams {
	return db.UpdateTariffCapabilitiesParams{
		Uid:                      util.DefaultString(input.UID, ""),
		CountryCode:              util.DefaultString(input.CountryCode, ""),
		PartyID:                  util.DefaultString(input.PartyID, ""),
		IsIntermediateCdrCapable: input.IsIntermediateCdrCapable,
	}
}
