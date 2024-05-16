package handlers

import (
	//"encoding/json"
	"encoding/json"
	"net/http"
	"time"

	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/redanthrax/cipp-go-api/internal/tools"
	"github.com/redanthrax/cipp-go-api/models"
	"github.com/redanthrax/cipp-go-api/pkg/msgraph"
	"github.com/rs/zerolog/log"
)


func (h *Handler) ListTenants(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("List Tenants")
  tenants, err := h.services.Tenants.GetAll()
  if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
    log.Error().Err(err).Msg("")
  }

  log.Info().Any("tenants", tenants).Msg("")

  if len(tenants) < 1 {
    //use graph to get tenants
    graph := r.Context().Value("graph").(*msgraphsdkgo.GraphServiceClient)
    tenants, err := msgraph.ListTenants(graph)
    if err != nil {
      tools.GraphError(err, w)
    }

    for _, t := range tenants {
      tenant := models.Tenant {
        LastRefresh: time.Now(),
        CustomerId: t.CustomerId,
        DefaultDomainName: t.DefaultDomainName,
        DisplayName: t.DisplayName,
      }

      err := h.services.Tenants.AddTenant(tenant)
      if err != nil {
        log.Error().Err(err).Msg("")
      }
    }
  }


  w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tenants); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
