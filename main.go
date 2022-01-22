package main

import (
	"appslab-ke/kipchoge-go/pkg/routes"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	handler := routes.Router()
	port := ":" + os.Getenv("PORT")
	srv := &http.Server{
		Addr: port,
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