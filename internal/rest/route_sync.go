package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (rs *RestService) mountSync() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/sync", func(w http.ResponseWriter, r *http.Request) {
		go rs.SyncService.Sync()
		w.Write([]byte("ok"))
	})

	return router
}
