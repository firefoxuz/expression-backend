package main

import (
	"expression-backend/internal/app/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

const (
	serverAddr = "0.0.0.0:8082"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/expressions", handlers.GetExpressions).Methods(http.MethodGet)
	r.HandleFunc("/expressions", handlers.StoreExpression).Methods(http.MethodPost)
	r.HandleFunc("/expressions/{id:[0-9A-Za-z]+}", handlers.GetExpression).Methods(http.MethodGet)

	srv := &http.Server{
		Handler:           r,
		Addr:              serverAddr,
		ReadHeaderTimeout: 15 * time.Second,
		WriteTimeout:      15 * time.Second,
	}

	log.Printf("server starting at %s", srv.Addr)
	srv.ListenAndServe()
}
