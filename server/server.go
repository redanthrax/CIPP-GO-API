package server

import (
	"net/http"
	"time"
  "context"
)

type Server struct {
  httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
  s.httpServer = &http.Server{
    Addr: ":" + port,
    ReadTimeout: 10 * time.Second,
    WriteTimeout: 10 * time.Second,
    Handler: handler,
  }

  return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
  return s.httpServer.Shutdown(ctx)
}
