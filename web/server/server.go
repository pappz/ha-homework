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
	service   service.Sector
	webServer http.Server
}

func NewServer() Server {
	return Server{}
}

func (s *Server) Listen(port string) {
	s.webServer = http.Server{
		Addr:         port,
		Handler:      controllers.Router(s.service),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

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
