package param

import (
	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-datastore/pkg/db"
)

func NewCreatePartyParams(input graph.CreatePartyInput) db.CreatePartyParams {
	return db.CreatePartyParams{
		CredentialID: input.CredentialID,
		CountryCode: input.CountryCode,
		PartyID: input.PartyID,
		IsIntermediateCdrCapable: input.IsIntermediateCdrCapable,
		PublishLocation: input.PublishLocation,
		PublishNullTariff: input.PublishNullTariff,
	}
}

func NewUpdatePartyByCredentialParams(input graph.UpdatePartyInput) db.UpdatePartyByCredentialParams {
	return db.UpdatePartyByCredentialParams{
		CredentialID: input.CredentialID,
		CountryCode: input.CountryCode,
		PartyID: input.PartyID,
		IsIntermediateCdrCapable: input.IsIntermediateCdrCapable,
		PublishLocation: input.PublishLocation,
		PublishNullTariff: input.PublishNullTariff,
	}
}
