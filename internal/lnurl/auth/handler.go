package auth

import (
	"log"
	"net/http"

	"github.com/fiatjaf/go-lnurl"
	"github.com/go-chi/render"
	"github.com/satimoto/go-api/internal/util"
	"github.com/satimoto/go-datastore/db"
)

func (r *LnUrlAuthResolver) GetHandler(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	k1 := request.URL.Query().Get("k1")
	sig := request.URL.Query().Get("sig")
	key := request.URL.Query().Get("key")

	if ok, _ := lnurl.VerifySignature(k1, sig, key); !ok {
		log.Printf("Error verifying signature: k1: %s sig: %s key: %s", k1, sig, key)
		render.JSON(rw, request, lnurl.ErrorResponse("Error verifying signature"))
		return
	}

	auth, err := r.AuthenticationResolver.Repository.GetAuthenticationByChallenge(ctx, k1)

	if err != nil {
		log.Printf("Authentication not found: k1: %s sig: %s key: %s", k1, sig, key)
		render.JSON(rw, request, lnurl.ErrorResponse("Error verifying signature"))
		return
	}

	r.AuthenticationResolver.Repository.UpdateAuthentication(ctx, db.UpdateAuthenticationParams{
		ID:            auth.ID,
		Signature:     util.SqlNullString(sig),
		LinkingPubkey: util.SqlNullString(key),
	})

	render.JSON(rw, request, lnurl.OkResponse())
}
