package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"rqpt/hadida/internal/db"
	"rqpt/hadida/internal/handlers"
)

func main() {
	pool, err := db.InitConnectionPool()
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}
	defer pool.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	handler := handlers.NewHandler(pool)

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		log.Fatal("APP_PORT is not set")
	}

	log.Printf("Server running on http://localhost:%s", appPort)
	log.Fatal(http.ListenAndServe(":"+appPort, handler.Routes()))
}
