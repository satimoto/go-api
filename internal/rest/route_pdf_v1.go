package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/satimoto/go-api/internal/pdf"
)

func (rs *RestService) mountPdf() *chi.Mux {
	router := chi.NewRouter()
	router.Mount("/invoice", rs.mountPdfInvoice())

	return router
}

func (rs *RestService) mountPdfInvoice() *chi.Mux {
	r := pdf.NewResolver(rs.RepositoryService)
	router := chi.NewRouter()

	router.Get("/{uid}", r.GetInvoicePdf)

	return router
}
