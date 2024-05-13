package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	api "github.com/redanthrax/cipp-go-api/api"
	"github.com/redanthrax/cipp-go-api/internal/tools"
	"github.com/rs/zerolog/log"
)

func ListTenants(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("List Tenants")
	graphReceive := r.Context().Value("graph")
	graph := graphReceive.(*msgraphsdk.GraphServiceClient)
	customers, err := graph.TenantRelationships().DelegatedAdminCustomers().Get(context.Background(), nil)
	if err != nil {
		tools.GraphError(err, w)
		return
	}

	iter, err := msgraphcore.NewPageIterator[models.DelegatedAdminCustomerable](
    customers,
    graph.GetAdapter(),
    models.CreateDelegatedAdminCustomerCollectionResponseFromDiscriminatorValue)
	if err != nil {
		tools.GraphError(err, w)
		return
	}

	var tenants []api.Tenant
	err = iter.Iterate(r.Context(), func(rel models.DelegatedAdminCustomerable) bool {
		tenants = append(tenants, api.Tenant{
			ID:   *rel.GetId(),
			Name: *rel.GetDisplayName(),
		})

		return true
	})

	if err != nil {
		tools.GraphError(err, w)
		return
	}

	if err := json.NewEncoder(w).Encode(tenants); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
