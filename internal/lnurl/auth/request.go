package auth

import (
	"log"
	"net/http"

	"github.com/fiatjaf/go-lnurl"
	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
)

func (r *LnUrlAuthResolver) GetLnUrlAuth(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	k1 := request.URL.Query().Get("k1")
	sig := request.URL.Query().Get("sig")
	key := request.URL.Query().Get("key")

	if ok, _ := lnurl.VerifySignature(k1, sig, key); !ok {
		log.Printf("API021: Error verifying signature")
		log.Printf("API021: K1=%v Sig=%v Key=%v", k1, sig, key)
		render.JSON(rw, request, lnurl.ErrorResponse("Error verifying signature"))
		return
	}

	auth, err := r.AuthenticationResolver.Repository.GetAuthenticationByChallenge(ctx, k1)

	if err != nil {
		log.Printf("API022: Authentication not found")
		log.Printf("API022: K1=%v Sig=%v Key=%v", k1, sig, key)
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
