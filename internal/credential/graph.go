package credential

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi-api/ocpirpc"
)

func NewCreateCredential(response ocpirpc.CreateCredentialResponse) *db.Credential {
	return &db.Credential{
		ID:          response.Id,
		Url:         response.Url,
		CountryCode: response.CountryCode,
		PartyID:     response.PartyId,
		IsHub:       response.IsHub,
	}
}
