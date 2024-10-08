package http

import (
	"context"
	"log"
	"midterm_cloud_project_2024/internal/core/helper"
	"net/http"
	"os"
	"time"
)

type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

func NewServer(handler http.Handler) *Server {
	httpServer := &http.Server{
		Handler:      handler,
		Addr:         os.Getenv("HTTP_PORT"),
		ReadTimeout:  helper.ParseDuration(os.Getenv("HTTP_READ_TIMEOUT")),
		WriteTimeout: helper.ParseDuration(os.Getenv("HTTP_WRITE_TIMEOUT")),
	}

	s := &Server{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: 5 * time.Second,
	}

	s.start()

	return s
}

func (s *Server) start() {
	go func() {
		log.Printf("Server is listening on %s", s.server.Addr)
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
