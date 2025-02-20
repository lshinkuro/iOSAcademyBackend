package config

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"path/filepath"

// 	"course-api/models"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"

// 	"github.com/joho/godotenv"
// )

// var DB *gorm.DB

// // loadEnv ensures the .env file is loaded only from the root directory
// func loadEnv() {
// 	rootPath, err := os.Getwd() // Ambil direktori root proyek
// 	if err != nil {
// 		log.Fatal("Error getting root directory:", err)
// 	}

// 	envPath := filepath.Join(rootPath, ".env") // Pastikan ambil dari root folder

// 	err = godotenv.Load(envPath) // Load hanya dari .env di root
// 	if err != nil {
// 		log.Fatal("Error loading .env file from root folder")
// 	}
// }

// func ConnectDB() {
// 	// Load environment variables from .env file
// 	loadEnv()

// 	log.Println("Attempting to connect to database...")

// 	// Construct DSN using environment variables
// 	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
// 		os.Getenv("PGHOST"),
// 		os.Getenv("PGUSER"),
// 		os.Getenv("PGPASSWORD"),
// 		os.Getenv("PGDATABASE"),
// 		os.Getenv("PGPORT"),
// 		os.Getenv("PGSSLMODE"),
// 	)

// 	// Log DSN parameters (without password)
// 	log.Printf("Database parameters: host=%s user=%s dbname=%s port=%s sslmode=%s\n",
// 		os.Getenv("PGHOST"),
// 		os.Getenv("PGUSER"),
// 		os.Getenv("PGDATABASE"),
// 		os.Getenv("PGPORT"),
// 		os.Getenv("PGSSLMODE"),
// 	)

// 	// Open database connection
// 	var err error
// 	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("Failed to connect to database: ", err)
// 	}

// 	log.Println("Database connection successfully opened")

// 	// Run auto-migrations
// 	log.Println("Attempting to run auto-migrations...")
// 	err = DB.AutoMigrate(
// 		&models.Course{},
// 		&models.User{},
// 		&models.Program{},
// 		&models.Material{},
// 		&models.ContentTopic{},
// 		&models.VideoCourse{})
// 	if err != nil {
// 		log.Fatal("Failed to run auto-migrations: ", err)
// 	}

// 	log.Println("Auto-migrations completed successfully")
// }

import (
	"fmt"
	"log"

	"course-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

// DB global variable
var DB *gorm.DB

func ConnectDB() {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Baca dari .env
	envMap, err := godotenv.Read(".env")
	if err != nil {
		log.Fatal("Failed to read .env file: ", err)
	}

	// DSN format untuk MySQL / TiDB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&tls=true",
		envMap["PGUSER"],
		envMap["PGPASSWORD"],
		envMap["PGHOST"],
		envMap["PGPORT"],
		envMap["PGDATABASE"],
	)

	// Koneksi pakai driver MySQL
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	log.Println("Database connection successfully opened")

	// Auto migrate
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
