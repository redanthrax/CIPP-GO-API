package middleware

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog/log"
)

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrapW := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		next.ServeHTTP(wrapW, r)
		log.Info().Str("Method", r.Method).Str("URI", r.RequestURI).Msg("")
	})
}
