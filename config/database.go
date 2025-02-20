package config

import (
	"fmt"
	"log"
	"os"

	"course-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDB() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Println("Attempting to connect to database...")

	// Construct DSN using environment variables
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("PGHOST"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
		os.Getenv("PGPORT"),
		os.Getenv("PGSSLMODE"),
	)

	// Log DSN parameters (without password)
	log.Printf("Database parameters: host=%s user=%s dbname=%s port=%s sslmode=%s\n",
		os.Getenv("PGHOST"),
		os.Getenv("PGUSER"),
		os.Getenv("PGDATABASE"),
		os.Getenv("PGPORT"),
		os.Getenv("PGSSLMODE"),
	)

	// Open database connection
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	log.Println("Database connection successfully opened")

	// Run auto-migrations
	log.Println("Attempting to run auto-migrations...")
	err = DB.AutoMigrate(
		&models.Course{},
		&models.User{},
		&models.Program{},
		&models.Material{},
		&models.ContentTopic{},
		&models.VideoCourse{})
	if err != nil {
		log.Fatal("Failed to run auto-migrations: ", err)
	}

	log.Println("Auto-migrations completed successfully")
}
