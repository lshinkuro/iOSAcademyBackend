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

	// Courses routes
	courses := v1.Group("/courses")
	courses.Get("/", handlers.GetAllCourses)
	courses.Get("/:id", handlers.GetCourse)
	courses.Post("/", handlers.CreateCourse)
	courses.Put("/:id", handlers.UpdateCourse)
	courses.Delete("/:id", handlers.DeleteCourse)
}