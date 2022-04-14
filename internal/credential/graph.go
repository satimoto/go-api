package credential

import (
	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-ocpi-api/ocpirpc/credentialrpc"
)

func NewCreateCredential(response credentialrpc.CredentialResponse) *graph.CreateCredential {
	return &graph.CreateCredential{
		ID:          response.Id,
		URL:         response.Url,
		CountryCode: response.CountryCode,
		PartyID:     response.PartyId,
		IsHub:       response.IsHub,
	}
}
