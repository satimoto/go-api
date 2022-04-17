package command

import (
	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-datastore/util"
	"github.com/satimoto/go-ocpi-api/ocpirpc/commandrpc"
)

func NewStartSession(response commandrpc.StartSessionResponse) *graph.StartSession {
	return &graph.StartSession{
		ID:              response.Id,
		Status:          response.Status,
		AuthorizationID: response.AuthorizationId,
		LocationUID:     response.LocationUid,
		EvseUID:         util.NilString(response.EvseUid),
	}
}

func NewStopSession(response commandrpc.StopSessionResponse) *graph.StopSession {
	return &graph.StopSession{
		ID:         response.Id,
		Status:     response.Status,
		SessionUID: response.SessionUid,
	}
}
