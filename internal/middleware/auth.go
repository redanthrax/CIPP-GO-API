package middleware

import (
	"context"
	"net/http"

	"github.com/redanthrax/cipp-go-api/internal/tools"
	"github.com/redanthrax/cipp-go-api/pkg/msgraph"
)

func GraphAuthenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		graph, err := msgraph.Authenticate()
		if err != nil {
			tools.GraphError(err, w)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		ctx := context.WithValue(r.Context(), "graph", graph)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
