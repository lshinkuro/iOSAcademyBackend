package routes

import (
	"course-api/handlers"
	"course-api/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Middleware
	app.Use(middleware.LoggerMiddleware())

	// API routes
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Auth routes (public)
	auth := v1.Group("/auth")
	auth.Post("/signup", handlers.SignUp)
	auth.Post("/signin", handlers.SignIn)

	// Courses routes (protected)
	courses := v1.Group("/courses")
	courses.Use(middleware.Protected()) // Add authentication middleware
	courses.Get("/", handlers.GetAllCourses)
	courses.Get("/:id", handlers.GetCourse)
	courses.Post("/", handlers.CreateCourse)
	courses.Put("/:id", handlers.UpdateCourse)
	courses.Delete("/:id", handlers.DeleteCourse)
}