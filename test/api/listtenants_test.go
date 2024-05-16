package api_test

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/redanthrax/cipp-go-api/pkg/msgraph"
)

func TestListGraphTenants(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Error(err)
	}

	graph, err := msgraph.Authenticate()
	if err != nil {
		t.Error(err)
	}

	tenants, err := msgraph.ListTenants(graph)
	if err != nil {
		t.Error(err)
	}

	if tenants == nil {
		t.Error("tenants empty")
	}

	if len(tenants) < 1 {
		t.Error("retrieved 0 tenants")
	}
}

