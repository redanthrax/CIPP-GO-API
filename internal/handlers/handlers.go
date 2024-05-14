package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	mid "github.com/redanthrax/cipp-go-api/internal/middleware"
)

func Handle(r *chi.Mux) {
	r.Use(middleware.StripSlashes)
	r.Use(mid.RequestLogger)
	r.Use(mid.GraphAuthenticate)
	r.Route("/api", func(r chi.Router) {
		r.Route("/ListTenants", func(r chi.Router) {
			r.Get("/", ListTenants)
		})
	})
}
