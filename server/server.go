package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"labx-graphql-go-todo/model"
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

	// init database connection
	// example: "host=localhost port=5432 user=postgres password=XXXXXX dbname=labx_todo sslmode=disable"
	log.Println("[INFO] connecting PG database")
	sqlConnString := os.Getenv("LABX_TODO_DB_CONN_STRING")
	ds, err := model.NewDataStore(sqlConnString)
	if err != nil {
		log.Fatal("[ERRO] Connection database error. ", err)
	}
	model.SetDataStore(ds)
	log.Println("[INFO] PG database connected OK")

	// initialize server
	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: newRouter(),
	}
}

// Start server
func (s *server) start() {
	log.Printf("[INFO] API server listening at http://localhost:%v", s.port)

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
