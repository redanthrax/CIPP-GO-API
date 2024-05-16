package service

import (
	"time"

	"github.com/redanthrax/cipp-go-api/internal/repository"
	"github.com/redanthrax/cipp-go-api/models"
	"github.com/redanthrax/cipp-go-api/pkg/msgraph"
)

type TenantsService struct {
  repo repository.Tenants
}

func NewTenantsService(repo repository.Tenants) *TenantsService {
  return &TenantsService{repo: repo}
}

func (s *TenantsService) GetAll() ([]models.Tenant, error) {
  tenants, err :=  s.repo.GetAllTenants()
  if err != nil {
    return nil, err
  }

  if len(tenants) < 1 {
    //we have no tenants, get them via graph
    graph, err := msgraph.Authenticate() 
    if err != nil {
      return nil, err
    }

    tenants, err := msgraph.ListTenants(graph)
    if err != nil {
      return nil, err
    }
    for _, t := range tenants {
      tenant := models.Tenant {
        LastRefresh: time.Now(),
        CustomerId: t.CustomerId,
        DefaultDomainName: t.DefaultDomainName,
        DisplayName: t.DisplayName,
      }

      err := s.AddTenant(tenant)
      if err != nil {
        return nil, err
      }
    }
  }

  return s.repo.GetAllTenants()
}

func (s *TenantsService) AddTenant(tenant models.Tenant) error {
  return s.repo.AddTenant(tenant)
}
