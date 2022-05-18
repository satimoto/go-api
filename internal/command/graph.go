package command

import (
	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi-api/ocpirpc"
)

func NewStartSession(response ocpirpc.StartSessionResponse) *graph.StartSession {
	return &graph.StartSession{
		ID:              response.Id,
		Status:          response.Status,
		AuthorizationID: response.AuthorizationId,
		LocationUID:     response.LocationUid,
		EvseUID:         util.NilString(response.EvseUid),
	}
}

func NewStopSession(response ocpirpc.StopSessionResponse) *graph.StopSession {
	return &graph.StopSession{
		ID:         response.Id,
		Status:     response.Status,
		SessionUID: response.SessionUid,
	}
}
