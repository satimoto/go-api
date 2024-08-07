package token

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi/ocpirpc"
)

func NewCreateToken(response ocpirpc.CreateTokenResponse) *db.Token {
	return &db.Token{
		ID:           response.Id,
		Uid:          response.Uid,
		Type:         db.TokenType(response.Type),
		AuthID:       response.AuthId,
		VisualNumber: util.SqlNullString(response.VisualNumber),
		Allowed:      db.TokenAllowedType(response.Allowed),
		Whitelist:    db.TokenWhitelistType(response.Whitelist),
	}
}
