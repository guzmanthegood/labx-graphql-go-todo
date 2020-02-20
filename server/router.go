package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func newRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeJSON)) 	// set response type as json by default
	r.Use(Recover)											// recover from panics
	r.Use(middleware.Logger)								// log transactions

	// api routes
	r.Route("/", func(r chi.Router) {
		r.Get("/status", Status)
		r.Get("/panic", Panic)
		r.NotFound(NotFound)
	})

	// log all routes
	walkFunc := func(m string, r string, h http.Handler, mi ...func(http.Handler) http.Handler) error {
		log.Printf("[INFO] %s %s\n", m, r)
		return nil
	}
	if err := chi.Walk(r, walkFunc); err != nil {
		log.Panicf("[ERRO] logging error: %s\n", err.Error())
	}

	return r
}

// Status handler function
func Status(w http.ResponseWriter, r *http.Request) {
	render.PlainText(w, r, "OK")
}

// Panic handler function
func Panic(w http.ResponseWriter, r *http.Request) {
	panic("panic test")
}

// NotFound render status function
func NotFound(w http.ResponseWriter, r *http.Request) {
	log.Printf("[ERRO] resource not found: %s\n", r.URL.Path)
	render.Status(r, http.StatusNotFound)
	render.PlainText(w, r, fmt.Sprintf("resource not found: %v", r.URL.Path))
}
