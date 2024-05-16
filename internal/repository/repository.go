package repository

import (
	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
	"github.com/redanthrax/cipp-go-api/models"
)

type Tenants interface {
  GetAllTenants() ([]models.Tenant, error)
  AddTenant(tenant models.Tenant) error
}

type Repository struct {
  Tenants
}

func NewRepository(db *aztables.ServiceClient) *Repository {
  return &Repository{
    Tenants: NewTenantsAzStorage(db),
  }
}

