package main

import (
	"context"
	"log"
	"net/http"
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

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler.Routes()))
}
