package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	mid "github.com/redanthrax/cipp-go-api/internal/middleware"
	"github.com/redanthrax/cipp-go-api/internal/service"
)

type Handler struct {
  services *service.Service
}

func NewHandler(services *service.Service) *Handler {
  return &Handler{services: services}
}

func (h *Handler) InitRoutes() http.Handler {
  r := chi.NewRouter()
	r.Use(middleware.StripSlashes)
	r.Use(mid.RequestLogger)
	r.Use(mid.GraphAuthenticate)
	r.Route("/api", func(r chi.Router) {
		r.Route("/ListTenants", func(r chi.Router) {
			r.Get("/", h.ListTenants)
		})
	})

  return r
}
