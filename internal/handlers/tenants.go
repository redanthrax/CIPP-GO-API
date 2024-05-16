package handlers

import (
	//"encoding/json"
	"encoding/json"
	"net/http"
	"github.com/rs/zerolog/log"
)


func (h *Handler) ListTenants(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("List Tenants")
  w.Header().Set("Content-Type", "application/json")
  tenants, err := h.services.Tenants.GetAll()
  if err != nil {
    h.HandleError(err, w)  
    return
  }

	if err := json.NewEncoder(w).Encode(tenants); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
