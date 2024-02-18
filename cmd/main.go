package main

import (
	"context"
	"expression-backend/api/handler"
	"expression-backend/internal/models"
	"expression-backend/internal/redis"
	"expression-backend/internal/services"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"math/rand"
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
	r := mux.NewRouter()
	r.HandleFunc("/expressions/{id:[0-9]+}", handler.GetExpression).Methods(http.MethodGet)
	r.HandleFunc("/expressions", handler.GetExpressions).Methods(http.MethodGet)
	r.HandleFunc("/expressions", handler.StoreExpression).Methods(http.MethodPost)
	r.HandleFunc("/agents", handler.GetAgents).Methods(http.MethodGet)
	dir := http.Dir("./assets")

	fs := http.FileServer(dir)
	r.PathPrefix("/").Handler(http.StripPrefix("/", fs))

	srv := &http.Server{
		Handler:           r,
		Addr:              serverAddr,
		ReadHeaderTimeout: 15 * time.Second,
		WriteTimeout:      15 * time.Second,
	}

	log.Printf("server starting at %s", srv.Addr)

	go func() {
		services.NewAgentCounters()
		rdb, err := redis.GetConnection()
		if err != nil {
			panic(err)
		}
		pubSub := rdb.Subscribe(context.Background(), "ping_channel")
		for {

			msg, err := pubSub.ReceiveMessage(context.Background())
			if err != nil {
				panic(err)
			}
			services.AddAgent(msg.Payload, time.Now())
			log.Println(msg.Channel, msg.Payload)
		}
	}()

	go func() {
		for {
			data := models.ExpressionData{}
			ms, _ := data.GetNotFinished()

			for _, model := range *ms {
				task := services.NewTask(rand.Int(), model.Expression, time.Duration(model.TimeLimit)*time.Millisecond)

				taskResponse := task.Execute()
				model.IsValid = taskResponse.IsValid
				model.Result = taskResponse.Result
				model.IsTimeLimit = taskResponse.IsTimeLimit
				model.Update()
			}

			time.Sleep(5 * time.Second)
		}
	}()

	srv.ListenAndServe()
}
