package rest

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	apiMiddleware "github.com/satimoto/go-api/internal/middleware"
)

type NostrNamesDto map[string]interface{}

type NostrDto struct {
	Names NostrNamesDto `json:"names"`
}

func (rs *RestService) mountWellKnown() *chi.Mux {
	router := chi.NewRouter()

	namesDto := make(NostrNamesDto)
	namesDto["satimoto"] = "23249c4d0e3dec5e29240dfc248ef9b5944558441e3363139c68fb3f587f1b3c"

	nostrDto := NostrDto{
		Names: namesDto,
	}

	router.Use(middleware.Logger)
	router.Use(apiMiddleware.IpContext())
	router.Get("/nostr.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(nostrDto)
	})

	return router
}
