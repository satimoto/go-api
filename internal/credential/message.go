package credential

import (
	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/util"
	"github.com/satimoto/go-ocpi-api/ocpirpc"
)

func NewCreateBusinessDetailRequest(input graph.CreateBusinessDetailInput) *ocpirpc.CreateBusinessDetailRequest {
	request := &ocpirpc.CreateBusinessDetailRequest{
		Name:    input.Name,
		Website: util.DefaultString(input.Website, ""),
	}

	if input.Logo != nil {
		request.Logo = NewCreateImageRequest(*input.Logo)
	}

	return request
}

func NewCreateCredentialRequest(input graph.CreateCredentialInput) *ocpirpc.CreateCredentialRequest {
	request := &ocpirpc.CreateCredentialRequest{
		ClientToken: util.DefaultString(input.ClientToken, ""),
		Url:         input.URL,
		PartyId:     input.PartyID,
		CountryCode: input.CountryCode,
		IsHub:       input.IsHub,
	}

	if input.BusinessDetail != nil {
		request.BusinessDetail = NewCreateBusinessDetailRequest(*input.BusinessDetail)
	}

	return request
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

func NewRegisterCredentialRequest(input graph.RegisterCredentialInput) *ocpirpc.RegisterCredentialRequest {
	return &ocpirpc.RegisterCredentialRequest{
		Id:          input.ID,
		ClientToken: util.DefaultString(input.ClientToken, ""),
	}
}

func NewUnregisterCredentialRequest(input graph.UnregisterCredentialInput) *ocpirpc.UnregisterCredentialRequest {
	return &ocpirpc.UnregisterCredentialRequest{
		Id: input.ID,
	}
}
