package command

import (
	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/util"
	"github.com/satimoto/go-ocpi-api/ocpirpc/commandrpc"
)

func NewStartSessionRequest(userID int64, input graph.StartSessionInput) *commandrpc.StartSessionRequest {
	return &commandrpc.StartSessionRequest{
		UserId:      userID,
		LocationUid: input.LocationUID,
		EvseUid:     util.DefaultString(input.EvseUID, ""),
	}
}

func NewStopSessionRequest(userID int64, input graph.StopSessionInput) *commandrpc.StopSessionRequest {
	return &commandrpc.StopSessionRequest{
		SessionUid: input.SessionUID,
	}
}
