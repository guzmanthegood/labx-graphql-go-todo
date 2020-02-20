package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/render"
)

// Recover middleware to rescue from panic
func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				switch x := err.(type) {
				case string:
					err = errors.New(x)
				case error:
					err = x
				default:
					err = errors.New("unknown panic")
				}

				render.Status(r, http.StatusInternalServerError)
				render.PlainText(w, r, fmt.Sprintf("internal server error: %v", err))
				log.Printf("[ERRO] internal server error: %v\n", err)
			}
		}()
		next.ServeHTTP(w, r)
	})
}