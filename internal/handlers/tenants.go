package handlers

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func ListTenants(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("List Tenants")
}
