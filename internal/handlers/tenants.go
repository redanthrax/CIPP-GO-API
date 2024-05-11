package handlers

import (
	"context"
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
	tenants, err := graph.TenantRelationships().Get(context.Background(), nil)
	if err != nil {
		log.Error().Err(err).Msg("")
		return
	}

	log.Info().Any("tenants", *tenants.GetMultiTenantOrganization().GetDisplayName()).Msg("")
}
