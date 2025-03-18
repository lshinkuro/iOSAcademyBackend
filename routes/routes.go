package routes

import (
	"course-api/handlers"
	"course-api/middleware"
	"course-api/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SetupRoutes(app *fiber.App) {
	// Swagger route
	app.Get("/api/v1/swagger/*", swagger.HandlerDefault)

	// Middleware global
	app.Use(middleware.LoggerMiddleware())

	// Base API Group
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Auth routes (public)
	auth := v1.Group("/auth")
	auth.Post("/signup", handlers.SignUp)
	auth.Post("/signin", handlers.SignIn)

	// Courses routes (protected)
	courses := v1.Group("/courses")
	courses.Use(middleware.Protected()) // Auth middleware for all courses routes
	courses.Get("/", handlers.GetAllCourses)
	courses.Get("/:id", handlers.GetCourse)

	// Only Admin & Mentor can modify courses
	courses.Use(middleware.RequireRole(models.RoleAdmin, models.RoleMentor))
	courses.Post("/", handlers.CreateCourse)
	courses.Put("/:id", handlers.UpdateCourse)
	courses.Delete("/:id", handlers.DeleteCourse)

	// Programs routes (protected)
	programs := v1.Group("/programs")
	programs.Use(middleware.Protected()) // Auth middleware for all programs routes
	programs.Get("/", handlers.GetAllPrograms)
	programs.Get("/:id", handlers.GetProgram)

	// Only Admin & Mentor can modify programs
	programs.Use(middleware.RequireRole(models.RoleAdmin, models.RoleMentor))
	programs.Post("/", handlers.CreateProgram)
	programs.Put("/:id", handlers.UpdateProgram)
	programs.Delete("/:id", handlers.DeleteProgram)

	// Materials routes (protected)
	materials := v1.Group("/materials")
	materials.Use(middleware.Protected()) // Auth middleware for all materials routes
	materials.Get("/", handlers.GetAllMaterials)
	materials.Get("/:id", handlers.GetMaterial)

	// Only Admin & Mentor can modify materials
	materials.Use(middleware.RequireRole(models.RoleAdmin, models.RoleMentor))
	materials.Post("/", handlers.CreateMaterial)
	materials.Put("/:id", handlers.UpdateMaterial)
	materials.Delete("/:id", handlers.DeleteMaterial)

	// Content Topic routes (protected)
	content := v1.Group("/content")
	content.Use(middleware.Protected())
	content.Get("/material/:material_id", handlers.GetContentTopics)
	content.Get("/:id", handlers.GetContentTopic)
	// Restrict content management to admin and mentor roles
	content.Use(middleware.RequireRole(models.RoleAdmin, models.RoleMentor))
	content.Post("/", handlers.CreateContentTopic)
	content.Put("/:id", handlers.UpdateContentTopic)
	content.Delete("/:id", handlers.DeleteContentTopic)
}
