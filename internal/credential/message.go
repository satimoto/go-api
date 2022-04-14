package credential

import (
	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/util"
	"github.com/satimoto/go-ocpi-api/ocpirpc"
	"github.com/satimoto/go-ocpi-api/ocpirpc/credentialrpc"
)

func NewCreateBusinessDetailRequest(input graph.CreateBusinessDetailInput) *ocpirpc.CreateBusinessDetailRequest {
	return &ocpirpc.CreateBusinessDetailRequest{
		Name:    input.Name,
		Website: util.DefaultString(input.Website, ""),
	}
}

func NewCreateCredentialRequest(input graph.CreateCredentialInput) *credentialrpc.CreateCredentialRequest {
	return &credentialrpc.CreateCredentialRequest{
		ClientToken: util.DefaultString(input.ClientToken, ""),
		Url:         input.URL,
		PartyId:     input.PartyID,
		CountryCode: input.CountryCode,
		IsHub:       input.IsHub,
	}
}

func NewCreateImageRequest(input graph.CreateImageInput) *ocpirpc.CreateImageRequest {
	return &ocpirpc.CreateImageRequest{
		Url:       input.URL,
		Thumbnail: util.DefaultString(input.Thumbnail, ""),
		Category:  string(input.Category),
		Type:      input.Type,
		Width:     int32(util.DefaultInt(input.Width, 0)),
		Height:    int32(util.DefaultInt(input.Height, 0)),
	}
}

func (r *CredentialResolver) CreateCredentialRequest(input graph.CreateCredentialInput) *credentialrpc.CreateCredentialRequest {
	response := NewCreateCredentialRequest(input)

	if input.BusinessDetail != nil {
		response.BusinessDetail = r.createBusinessDetailRequest(*input.BusinessDetail)
	}

	return response
}

func (r *CredentialResolver) createBusinessDetailRequest(input graph.CreateBusinessDetailInput) *ocpirpc.CreateBusinessDetailRequest {
	response := NewCreateBusinessDetailRequest(input)

	if input.Logo != nil {
		response.Logo = NewCreateImageRequest(*input.Logo)
	}

	return response
}
