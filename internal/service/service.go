package service

import (
	"github.com/redanthrax/cipp-go-api/internal/repository"
	"github.com/redanthrax/cipp-go-api/models"
)

type Tenants interface {
  GetAll() ([]models.Tenant, error)
  AddTenant(tenant models.Tenant) error
}

type Service struct {
  Tenants
}

func NewService(repo *repository.Repository) *Service {
  return &Service{
    Tenants: NewTenantsService(repo.Tenants),
  }
}
