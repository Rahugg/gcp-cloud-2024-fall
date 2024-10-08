package server

import (
	"fmt"
	"log"
	"midterm_cloud_project_2024/internal/routes"
	"midterm_cloud_project_2024/pkg/http"
	"os"
	"os/signal"
	"syscall"
)

type Server struct{}

func New() (*Server, error) {
	return &Server{}, nil
}

func (s *Server) Run() {
	mux := routes.NewRouter()

	server := http.NewServer(mux)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Printf("Shutting down, signal: %s", s)
	case err := <-server.Notify():
		log.Println(fmt.Errorf("Error in server: %w", err))
	}

	if err := server.Shutdown(); err != nil {
		log.Printf("Error during shutdown: %v", err)
	}
}
