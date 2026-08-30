package main

import (
	"log"
	"net/http"
	"os"

	"rqpt/hadida/internal/db"
	"rqpt/hadida/internal/handlers"
)

func main() {
	pool, err := db.InitConnectionPool()
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}
	defer pool.Close()

	handler := handlers.NewHandler(pool)

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		log.Fatal("APP_PORT is not set")
	}

	log.Printf("Server running on http://localhost:%s", appPort)
	log.Fatal(http.ListenAndServe(":"+appPort, handler.Routes()))
}
