package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/redanthrax/cipp-go-api/internal/handlers"
	"github.com/redanthrax/cipp-go-api/internal/repository"
	"github.com/redanthrax/cipp-go-api/internal/service"
	"github.com/redanthrax/cipp-go-api/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	if err := godotenv.Load(); err != nil {
		log.Fatal().Err(err).Msg("")
  }

  dbConfig := repository.Config{
    StorageAccount: os.Getenv("AzureWebJobsStorage"),
  }

  db, err := repository.NewDB(dbConfig)
  if err != nil {
    log.Fatal().Err(err).Msg("")
  }

  repo := repository.NewRepository(db)
  services := service.NewService(repo)
  hand := handlers.NewHandler(services)
	listenAddr := "8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = val
	}

  srv := new(server.Server)
  go func() {
    if err = srv.Run(listenAddr, hand.InitRoutes()); err != nil && err != http.ErrServerClosed {
      log.Fatal().Err(err).Msg("error running server")
    }
  }()
  
  log.Info().Str("port", listenAddr).Msg("server listening")
  ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
  defer stop()
  <-ctx.Done()
  log.Info().Msg("server shutting down")
  if err := srv.Shutdown(context.Background()); err != nil {
    log.Error().Err(err).Msg("error shutting down the server")
  }
}
