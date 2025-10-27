package main

import (
	"context"
	"log"
	"net/http"
	"order-service-demo/internal/config"
	"order-service-demo/internal/db"
	"order-service-demo/internal/handler"
	"order-service-demo/internal/service"
	"order-service-demo/pkg/cache"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	cfg := config.Load()

	log.Println("Connecting to database...")
	dbPool := db.InitDB(cfg.DBConnString)
	defer dbPool.Close()
	log.Println("Database connected successfully")

	cache.Init()
	log.Println("Cache initialized")

	log.Println("Loading cache from database...")
	ordersCount := db.LoadCacheFromDB(dbPool)
	log.Printf("Cache loaded: %d orders", ordersCount)

	log.Println("Connecting to NATS...")
	service.InitNATSAndSubscribe(cfg.NATSURL, dbPool)

	r := mux.NewRouter()
	r.HandleFunc("/order/{id}", handler.GetOrder).Methods("GET")
	r.HandleFunc("/ui", handler.UIHandler).Methods("GET")
	r.HandleFunc("/ui", handler.GetOrderUI).Methods("POST")

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		log.Println("Starting HTTP server on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited properly")
}
