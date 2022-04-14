package credential

import (
	"time"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-datastore/util"
)

func NewCreateCredentialParams(input graph.CreateCredentialInput) db.CreateCredentialParams {
	return db.CreateCredentialParams{
		ClientToken: util.SqlNullString(input.ClientToken),
		Url:         input.URL,
		CountryCode: input.CountryCode,
		PartyID:     input.PartyID,
		IsHub:       input.IsHub,
		LastUpdated: time.Now(),
	}
}
