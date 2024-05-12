package handlers

import (
	"context"
	"net/http"

	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/redanthrax/cipp-go-api/internal/tools"
	"github.com/rs/zerolog/log"
	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
)

func ListTenants(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("List Tenants")
	graphReceive := r.Context().Value("graph")
	graph := graphReceive.(*msgraphsdk.GraphServiceClient)
	tenants, err := graph.TenantRelationships().Get(context.Background(), nil)
	if err != nil {
		tools.GraphError(err)
		return
	}

	// Use PageIterator to iterate through all users
	pageIterator, err := msgraphcore.NewPageIterator[models.TenantRelationshipable](
		tenants,
		graph.GetAdapter(),
		models.CreateTenantRelationshipFromDiscriminatorValue)
	err = pageIterator.Iterate(context.Background(), func(tenant models.TenantRelationshipable) bool {
		log.Info().Str("tenant", *tenant.GetMultiTenantOrganization().GetDisplayName())
		return true
	})

	if err != nil {
		tools.GraphError(err)
		return
	}

	log.Info().Any("tenants", *tenants.GetMultiTenantOrganization().GetDisplayName()).Msg("")
}
