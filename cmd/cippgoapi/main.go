package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/redanthrax/cipp-go-api/internal/handlers"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
  err := godotenv.Load()
  if err != nil {
    log.Error().Err(err).Msg("")
  }

	var r *chi.Mux = chi.NewRouter()
	handlers.Handle(r)
  listenAddr := ":8080"
  if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
    listenAddr = ":" + val
  }

	log.Info().Msg("Starting API Server...")
	err = http.ListenAndServe(listenAddr, r)
	if err != nil {
		log.Error().Err(err).Msg("")
	}
}
