package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/pappz/ha-homework/service"
	"github.com/pappz/ha-homework/web/controllers"
)

type Server struct {
	webServer *http.Server
}

func NewServer(port string, service service.Sector) Server {
	httpServer := http.Server{
		Addr:         port,
		Handler:      controllers.Router(service),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return Server{
		&httpServer,
	}
}

func (s *Server) Listen() {
	go func() {
		err := s.webServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe error: %s", err)
		}
	}()
}

func (s *Server) TearDown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	return s.webServer.Shutdown(ctx)
}
