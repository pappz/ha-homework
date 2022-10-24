package server

import (
	"context"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/pappz/ha-homework/service"
	"github.com/pappz/ha-homework/web/api"
)

type Server struct {
	webServer *http.Server
}

// NewServer create a new http server. The service used by controllers.
// The server will listen on the 'addr' address.
func NewServer(addr string, service service.Sector) Server {
	httpServer := http.Server{
		Addr:         addr,
		Handler:      api.Router(service),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return Server{
		&httpServer,
	}
}

// Listen listens on the TCP network address s.Addr.
// In case of error the program drop a fatal error
func (s *Server) Listen() {
	go func() {
		err := s.webServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen error: '%s'", err)
		}
	}()
}

// TearDown gracefully shuts down the server.
// The timeout is 30 second by default.
func (s *Server) TearDown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	return s.webServer.Shutdown(ctx)
}
