package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (rs *RestService) mountHealth() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	return router
}
