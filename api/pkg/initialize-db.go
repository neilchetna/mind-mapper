package pkg

import (
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDatabase() (*gorm.DB, error) {
	dbHost := os.Getenv("PG_HOST")
	dbPort := os.Getenv("PG_PORT")
	dbUser := os.Getenv("PG_USER")
	dbPass := os.Getenv("PG_PASSWORD")
	dbName := os.Getenv("PG_DB_NAME")

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)

	var config gorm.Config
	db, err := gorm.Open(postgres.Open(connectionString), &config)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	pool, err := db.DB()
	if err != nil {
		return nil, errors.New("failed to connect to PostgreSQL database")
	}

	if err := pool.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}

	return db, nil
}
