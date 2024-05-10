package handlers

import (
	"net/http"

	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/rs/zerolog/log"
)

func ListTenants(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("List Tenants")
	graphReceive := r.Context().Value("graph")
	if graphReceive == nil {
		log.Error().Msg("Graph context is nil")
		return
	}

	graph := graphReceive.(*msgraphsdk.GraphServiceClient)
	//tenants := graph.TenantRelationships().Get()
	log.Info().Any("tenants", tenants).Msg("")
}
