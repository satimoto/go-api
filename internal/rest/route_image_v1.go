package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/satimoto/go-api/internal/image"
)

func (rs *RestService) mountImage() *chi.Mux {
	router := chi.NewRouter()
	router.Mount("/circuit", rs.mountImageCircuit())

	return router
}

func (rs *RestService) mountImageCircuit() *chi.Mux {
	r := image.NewResolver(rs.RepositoryService)
	router := chi.NewRouter()

	router.Get("/{referral_code}", r.GetReferralCodeImage)

	return router
}
