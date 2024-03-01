package rest

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	apiMiddleware "github.com/satimoto/go-api/internal/middleware"
	"github.com/satimoto/go-api/internal/notify"
)

func (rs *RestService) mountV1() *chi.Mux {
	r := notify.NewResolver(rs.RepositoryService)
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(apiMiddleware.IpContext())
	router.Mount("/", rs.mountGraphql())
	router.Mount("/image", rs.mountImage())
	router.Mount("/lnurl", rs.mountLnUrl())
	router.Mount("/pdf", rs.mountPdf())
	router.Post("/notify", r.PostNotify)

	return router
}
