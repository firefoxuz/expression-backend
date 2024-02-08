package main

import (
	"expression-backend/api/handler"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

const (
	serverAddr = "0.0.0.0:8082"
)

func init() {
	viper.SetConfigName(".env.json")
	viper.SetConfigType("json")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func main() {
	dbSourceName := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable", "postgres", viper.GetString("database.name"), viper.GetString("database.username"), viper.GetString("database.password"))
	_, err := sqlx.Connect("postgres", dbSourceName)
	if err != nil {
		log.Fatalln(err)
	}
	r := mux.NewRouter()

	r.HandleFunc("/expressions", handler.GetExpressions).Methods(http.MethodGet)
	r.HandleFunc("/expressions", handler.StoreExpression).Methods(http.MethodPost)
	r.HandleFunc("/expressions/{id:[0-9A-Za-z]+}", handler.GetExpression).Methods(http.MethodGet)

	srv := &http.Server{
		Handler:           r,
		Addr:              serverAddr,
		ReadHeaderTimeout: 15 * time.Second,
		WriteTimeout:      15 * time.Second,
	}

	log.Printf("server starting at %s", srv.Addr)
	srv.ListenAndServe()
}
