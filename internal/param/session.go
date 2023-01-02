package param

import (
	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
)

func NewListSessionInvoicesByUserIDParams(userID int64, input graph.ListSessionInvoicesInput) db.ListSessionInvoicesByUserIDParams {
	return db.ListSessionInvoicesByUserIDParams{
		ID: userID,
		IsSettled: util.DefaultBool(input.IsSettled, false),
		IsExpired: util.DefaultBool(input.IsExpired, false),
	}
}
