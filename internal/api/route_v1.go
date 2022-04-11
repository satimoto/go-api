package api

import "github.com/go-chi/chi/v5"


func (rs *RouterService) mountV1() *chi.Mux {
	router := chi.NewRouter()

	router.Mount("/", rs.mountGraphql())
	router.Mount("/lnurl", rs.mountLnUrl())

	return router
}