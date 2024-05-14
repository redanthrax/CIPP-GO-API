package handlers

import (
	"encoding/json"
	"net/http"

	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/redanthrax/cipp-go-api/internal/tools"
	"github.com/redanthrax/cipp-go-api/pkg/msgraph"
	"github.com/rs/zerolog/log"
)

func ListTenants(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("List Tenants")
	graph := r.Context().Value("graph").(*msgraphsdk.GraphServiceClient)
	tenants, err := msgraph.ListTenants(graph)
	if err != nil {
		tools.GraphError(err, w)
	}

	if err := json.NewEncoder(w).Encode(tenants); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
