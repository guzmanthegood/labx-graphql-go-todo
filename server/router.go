package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func newRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeJSON)) 	// set response type as json by default
	r.Use(Recover)											// recover from panics
	//r.Use(middleware.Logger)								// log transactions

	http.Handle("/", http.FileServer(http.Dir("./server/")))

	// api routes
	r.Route("/", func(r chi.Router) {
		r.Get("/", playground)
		r.Get("/status", status)
		r.Get("/panic", panicTest)
		r.Post("/query", graphQLService)
		// r.NotFound(notFound)
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

func status(w http.ResponseWriter, r *http.Request) {
	render.PlainText(w, r, "OK")
}

func panicTest(w http.ResponseWriter, r *http.Request) {
	panic("panic test")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	log.Printf("[ERRO] resource not found: %s\n", r.URL.Path)
	render.Status(r, http.StatusNotFound)
	render.PlainText(w, r, fmt.Sprintf("resource not found: %v", r.URL.Path))
}

func playground(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("./server/playground.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	html, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	render.HTML(w, r, string(html))
}