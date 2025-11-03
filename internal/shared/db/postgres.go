package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func ConnectPostgres(host string, port int, user, password, name string) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, name,
	)

	var db *sql.DB
	var err error

	maxRetries := 5
	retryDelay := 2 * time.Second

	for attempt := 1; attempt <= maxRetries; attempt++ {
		db, err = sql.Open("postgres", dsn)
		if err != nil {
			log.Printf("[db] Attempt %d/%d: failed to open connection: %v", attempt, maxRetries, err)
		} else if pingErr := db.Ping(); pingErr != nil {
			err = pingErr
			log.Printf("[db] Attempt %d/%d: failed to ping DB: %v", attempt, maxRetries, err)
		} else {
			log.Printf("[db] Connected to PostgreSQL on attempt %d", attempt)
			return db, nil
		}

		log.Printf("[db] Retry in %v...", retryDelay)
		time.Sleep(retryDelay)
	}

	return nil, fmt.Errorf("could not connect to PostgreSQL after %d attempts: %v", maxRetries, err)
}
