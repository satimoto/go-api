package rest

import (
	apiMiddleware "github.com/satimoto/go-api/internal/middleware"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func (rs *RestService) mountV1() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(apiMiddleware.IpContext())
	router.Mount("/", rs.mountGraphql())
	router.Mount("/image", rs.mountImage())
	router.Mount("/lnurl", rs.mountLnUrl())
	router.Mount("/pdf", rs.mountPdf())

	return router
}
