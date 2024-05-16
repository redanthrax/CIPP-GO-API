package service

import (
	"github.com/redanthrax/cipp-go-api/internal/repository"
	"github.com/redanthrax/cipp-go-api/models"
)

type TenantsService struct {
  repo repository.Tenants
}

func NewTenantsService(repo repository.Tenants) *TenantsService {
  return &TenantsService{repo: repo}
}

func (s *TenantsService) GetAll() ([]models.Tenant, error) {
  return s.repo.GetAllTenants()
}

func (s *TenantsService) AddTenant(tenant models.Tenant) error {
  return s.repo.AddTenant(tenant)
}
