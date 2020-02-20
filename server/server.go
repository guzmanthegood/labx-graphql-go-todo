package main

import (
	"context"
	"log"
	"net/http"
	"os"
)

type server struct {
	server *http.Server
	port   string
}

func main() {
	log.Println("====== LABX-GRAPHQL-GO-TODO SERVER ======")
	s := server{}
	s.initialize()
	s.start()
}

const defaultPort = "8080"

func (s *server) initialize() {
	// get service port
	s.port = defaultPort
	if os.Getenv("PORT") != "" {
		s.port = os.Getenv("PORT")
	}

	// initialize server
	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: newRouter(),
	}
}

// Start server
func (s *server) start() {
	log.Printf("[INFO] API server listening at http://localhost:%v", defaultPort)

	if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal("[ERRO] internal server error", err)
	}
}

// Stop server
func (s *server) stop() {
	if err := s.server.Shutdown(context.Background()); err != nil {
		log.Fatal("[ERRO] error during shutdown", err)
	}
}
