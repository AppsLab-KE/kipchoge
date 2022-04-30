package main

import (
	"appslab-ke/kipchoge-go/internal/storage"
	"appslab-ke/kipchoge-go/pkg/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	dbClient, err := storage.NewStorage()
	if err != nil {
		log.Fatalf("failed to connect to dbClient %v", err)
	}

	err = dbClient.Ping()
	if err != nil {
		log.Printf("fail to ping database %v", err)
	}

	defer func(db storage.Storage) {
		err := db.Defer()
		if err != nil {
			log.Fatalln("failed to close connection")
		}
	}(dbClient)

	handler := handlers.NewApp()
	port := ":" + os.Getenv("PORT")
	srv := &http.Server{
		Addr:    port,
		Handler: handler,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln("Failed to start server.")
		}
	}()

	// Shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	log.Println("Shutting down server")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalln("Forced to shutdown")
	}
}
