package rest

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func (rs *RestService) mountV1() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Mount("/", rs.mountGraphql())
	router.Mount("/lnurl", rs.mountLnUrl())

	return router
}
