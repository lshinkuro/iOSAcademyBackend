package main

import (
	"course-api/config"
	"course-api/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Println("Starting Course API server...")

	// Initialize database
	log.Println("Initializing database connection...")
	config.ConnectDB()
	log.Println("Database connection initialized successfully")

	// Initialize Redis connection
	log.Println("Initializing Redis connection...")
	config.ConnectRedis()
	defer config.CloseRedis()
	log.Println("Redis connection established")

	// Create Fiber app
	log.Println("Creating Fiber application...")
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			log.Printf("Error occurred: %v\n", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Middleware
	// log.Println("Setting up middleware...")
	// app.Use(cors.New())

	// ✅ Tambahkan middleware CORS di sini
	log.Println("Setting up CORS middleware...")
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Bisa diganti dengan domain frontend
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Setup routes
	log.Println("Setting up routes...")
	routes.SetupRoutes(app)
	log.Println("Routes configured successfully")

	// Start server
	log.Println("Starting server on 0.0.0.0:3000...")
	if err := app.Listen(":3000"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
