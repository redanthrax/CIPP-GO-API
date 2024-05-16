package repository

import (
	"context"
	"encoding/json"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
	"github.com/redanthrax/cipp-go-api/models"
	"github.com/rs/zerolog/log"
)

type TenantsAzStorage struct {
  db *aztables.ServiceClient
  tb *aztables.Client
}

func NewTenantsAzStorage(db *aztables.ServiceClient) *TenantsAzStorage {
  //initialize table
  _, err := db.DeleteTable(context.TODO(), "Tenants", nil)
  _, err = db.CreateTable(context.TODO(), "Tenants", nil)
  if err != nil {
    log.Error().Err(err).Msg("could not create tenants table")
  }

  return &TenantsAzStorage{db: db, tb: db.NewClient("Tenants")}
}

func (r *TenantsAzStorage) GetAllTenants() ([]models.Tenant, error) {
  //get the tenants in the db
  options := &aztables.ListEntitiesOptions {
    Top: to.Ptr(int32(100)),
  }

  pager := r.tb.NewListEntitiesPager(options)
  var tenants []models.Tenant
  for pager.More() {
    resp, err := pager.NextPage(context.TODO())
    if err != nil {
      return nil, err
    }

    for _, entity := range resp.Entities {
      var tenant models.Tenant
      err = json.Unmarshal(entity, &tenant)
      if err != nil {
        return nil, err
      }

      tenants = append(tenants, tenant)
    }
  }

  return tenants, nil
}

func (r *TenantsAzStorage) AddTenant(tenant models.Tenant) error {
  properties := map[string]interface{}{
    "ExcludeDate":              tenant.ExcludeDate,
    "ExcludeUser":              tenant.ExcludeUser,
    "Excluded":                 tenant.Excluded,
    "GraphErrorCount":          tenant.GraphErrorCount,
    "LastGraphError":           tenant.LastGraphError,
    "LastRefresh":              aztables.EDMDateTime(tenant.LastRefresh),
    "RequiresRefresh":          tenant.RequiresRefresh,
    "CustomerId":               tenant.CustomerId,
    "DefaultDomainName":        tenant.DefaultDomainName,
    "DelegatedPrivilegeStatus": tenant.DelegatedPrivilegeStatus,
    "DisplayName":              tenant.DisplayName,
    "Domains":                  tenant.Domains,
    "HasAutoExtend":            tenant.HasAutoExtend,
    "InitialDomainName":        tenant.InitialDomainName,
    "RelationshipCount":        tenant.RelationshipCount,
    "RelationshipEnd":          aztables.EDMDateTime(tenant.RelationshipEnd),
  }
  entity := aztables.EDMEntity {
    Entity: aztables.Entity {
      PartitionKey: "tenants",
    },
    Properties: properties,
  }

  log.Info().Any("entity", entity).Msg("")

  marshalled, err := json.Marshal(entity)
  if err != nil {
    return err
  }

  log.Info().Any("marsh", marshalled).Msg("")

  _, err = r.tb.AddEntity(context.TODO(), marshalled, nil)
  if err != nil {
    return err
  }

  return nil
}
