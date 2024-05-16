package service

import (
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/redanthrax/cipp-go-api/internal/repository"
	"github.com/redanthrax/cipp-go-api/models"
)

type Tenants interface {
  GetAll() ([]models.Tenant, error)
  AddTenant(tenant models.Tenant) error
}

type Service struct {
  Graph *msgraphsdkgo.GraphServiceClient
  Tenants
}

func NewService(repo *repository.Repository, graph *msgraphsdkgo.GraphServiceClient) *Service {
  return &Service{
    Graph: graph,
    Tenants: NewTenantsService(repo.Tenants),
  }
}
