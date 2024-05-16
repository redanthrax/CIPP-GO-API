package msgraph

import (
	"context"

	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	azmodels "github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/redanthrax/cipp-go-api/models"
)

func ListTenants(graph *msgraphsdk.GraphServiceClient) ([]models.Tenant, error) {
	customers, err := graph.TenantRelationships().DelegatedAdminCustomers().Get(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	iter, err := msgraphcore.NewPageIterator[azmodels.DelegatedAdminCustomerable](
		customers,
		graph.GetAdapter(),
		azmodels.CreateDelegatedAdminCustomerCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, err
	}

	var tenants []models.Tenant
	err = iter.Iterate(context.Background(), func(rel azmodels.DelegatedAdminCustomerable) bool {
		var tenant models.Tenant
		tenant.CustomerId = *rel.GetId()
		tenant.DisplayName = *rel.GetDisplayName()
		d, _ := graph.TenantRelationships().FindTenantInformationByTenantIdWithTenantId(&tenant.CustomerId).Get(context.Background(), nil)
		if d != nil {
			tenant.DefaultDomainName = *d.GetDefaultDomainName()
		}

		tenants = append(tenants, tenant)
		return true
	})

	if err != nil {
		return nil, err
	}

	return tenants, nil
}
