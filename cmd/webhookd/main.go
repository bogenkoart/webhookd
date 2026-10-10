package main

import (
	"log"
	"net/http"
	"time"

	"github.com/bogenkoart/webhookd/internal/api"
)

func main() {
	srv := &http.Server{
		Addr:              ":8080",
		Handler:           api.NewRouter(),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Printf("webhookd: listening on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
