package main

import (
	"log"

	"rqpt/hadida/internal/db"
)

func main() {
	pool, err := db.InitConnectionPool()
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}
	defer pool.Close()
}
