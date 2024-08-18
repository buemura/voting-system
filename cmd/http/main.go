package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/buemura/voting-system/internal/config"
	"github.com/buemura/voting-system/internal/database"
	"github.com/buemura/voting-system/internal/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func init() {
	config.LoadEnv()
	database.Connect()
}

func main() {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	h := handler.RegisterRoutes(mux)

	port := ":" + config.PORT
	srv := &http.Server{
		Addr:    port,
		Handler: h,
	}

	go func() {
		log.Println("HTTP Server running at", port, "...")
		if err := srv.ListenAndServe(); err != nil && http.ErrServerClosed != err {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}
	log.Println("Server gracefully stopped")
}
