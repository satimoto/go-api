package command

import (
	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/util"
	"github.com/satimoto/go-ocpi/ocpirpc"
)

func NewStartSessionRequest(userID int64, input graph.StartSessionInput) *ocpirpc.StartSessionRequest {
	return &ocpirpc.StartSessionRequest{
		UserId:      userID,
		LocationUid: input.LocationUID,
		EvseUid:     util.DefaultString(input.EvseUID, ""),
	}
}

func NewStopSessionRequest(userID int64, input graph.StopSessionInput) *ocpirpc.StopSessionRequest {
	return &ocpirpc.StopSessionRequest{
		AuthorizationId: input.AuthorizationID,
	}
}
