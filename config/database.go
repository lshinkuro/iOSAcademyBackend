package config

import (
	"fmt"
	"log"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"course-api/models"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	// Log the connection attempt
	log.Println("Attempting to connect to database...")

	// Using more robust connection string with additional parameters
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require pool_max_conns=10 pool_min_conns=1",
		os.Getenv("PGHOST"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
		os.Getenv("PGPORT"),
	)

	// Log DSN parameters (without password)
	log.Printf("Database parameters: host=%s user=%s dbname=%s port=%s\n",
		os.Getenv("PGHOST"),
		os.Getenv("PGUSER"),
		os.Getenv("PGDATABASE"),
		os.Getenv("PGPORT"),
	)

	// Open database connection with retries
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
		PreferSimpleProtocol: true, // Disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	log.Println("Database connection successfully opened")

	// Log auto-migration attempt
	log.Println("Attempting to run auto-migrations...")

	// Add User model to auto-migrations
	err = DB.AutoMigrate(&models.Course{}, &models.User{})
	if err != nil {
		log.Fatal("Failed to run auto-migrations: ", err)
	}

	log.Println("Auto-migrations completed successfully")
}
