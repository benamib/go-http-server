package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

const PORT = 8080

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.WithField("path", "/").Infof("Incomming Request")
		if _, err := fmt.Fprintf(w, "Hello Workd\n"); err != nil {
			log.WithError(err).Fatal("unexpected error")
		}
	})

	log.WithField("port", PORT).Info("Serving http server...")
	if err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil); err != nil {
		log.WithError(err).Fatal("unexpected error")
		return
	}
}
