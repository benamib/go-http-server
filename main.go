package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	log "github.com/sirupsen/logrus"
)

const PORT = 8080

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprintf(w, "Hello World\n"); err != nil {
			log.WithError(err).Fatal("unexpected error")
		}
	})

	log.WithField("port", PORT).Info("Serving http server...")
	if err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), r); err != nil {
		log.WithError(err).Fatal("unexpected error")
		return
	}
}
