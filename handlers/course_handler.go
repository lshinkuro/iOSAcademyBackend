package handlers

import (
	"context"
	"course-api/config"
	"course-api/models"
	"course-api/responses"
	"course-api/utils/cache"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

// GetAllCourses returns all courses with Redis caching
func GetAllCourses(c *fiber.Ctx) error {
	ctx := context.Background()
	cacheKey := "courses:all"
	var courses []models.Course

	// Try to get courses from cache
	err := cache.Get(ctx, cacheKey, &courses)
	if err == nil {
		return responses.SendSuccess(c, "Courses found in cache", courses)
	}

	// If not in cache, get from database
	if err := config.DB.Find(&courses).Error; err != nil {
		return responses.SendError(c, fiber.StatusInternalServerError, "Error fetching courses")
	}

	// Store in cache
	if err := cache.Set(ctx, cacheKey, courses, cache.DefaultExpiration); err != nil {
		// Log the error but don't fail the request
		fmt.Printf("Error caching courses: %v\n", err)
	}

	return responses.SendSuccess(c, "Courses found successfully", courses)
}

// GetCourse returns a single course with Redis caching
func GetCourse(c *fiber.Ctx) error {
	id := c.Params("id")
	ctx := context.Background()
	cacheKey := fmt.Sprintf("courses:%s", id)
	var course models.Course

	// Try to get course from cache
	err := cache.Get(ctx, cacheKey, &course)
	if err == nil {
		return responses.SendSuccess(c, "Course found in cache", course)
	}

	// If not in cache, get from database
	if err := config.DB.First(&course, id).Error; err != nil {
		return responses.SendError(c, fiber.StatusNotFound, "Course not found")
	}

	// Store in cache
	if err := cache.Set(ctx, cacheKey, course, cache.DefaultExpiration); err != nil {
		fmt.Printf("Error caching course: %v\n", err)
	}

	return responses.SendSuccess(c, "Course found successfully", course)
}

// CreateCourse creates a new course and invalidates cache
func CreateCourse(c *fiber.Ctx) error {
	input := new(models.CreateCourseInput)

	if err := c.BodyParser(input); err != nil {
		return responses.SendError(c, fiber.StatusBadRequest, "Invalid input")
	}

	if err := validate.Struct(input); err != nil {
		return responses.SendError(c, fiber.StatusBadRequest, err.Error())
	}

	course := models.Course{
		Title:       input.Title,
		Description: input.Description,
		Instructor:  input.Instructor,
		Duration:    input.Duration,
		Price:       input.Price,
	}

	if err := config.DB.Create(&course).Error; err != nil {
		return responses.SendError(c, fiber.StatusInternalServerError, "Error creating course")
	}

	// Invalidate the all courses cache
	ctx := context.Background()
	_ = cache.Delete(ctx, "courses:all")

	return responses.SendSuccess(c, "Course created successfully", course)
}

// UpdateCourse updates an existing course
func UpdateCourse(c *fiber.Ctx) error {
	id := c.Params("id")
	input := new(models.UpdateCourseInput)

	if err := c.BodyParser(input); err != nil {
		return responses.SendError(c, fiber.StatusBadRequest, "Invalid input")
	}

	var course models.Course
	if err := config.DB.First(&course, id).Error; err != nil {
		return responses.SendError(c, fiber.StatusNotFound, "Course not found")
	}

	if input.Title != "" {
		course.Title = input.Title
	}
	if input.Description != "" {
		course.Description = input.Description
	}
	if input.Instructor != "" {
		course.Instructor = input.Instructor
	}
	if input.Duration != 0 {
		course.Duration = input.Duration
	}
	if input.Price != 0 {
		course.Price = input.Price
	}

	config.DB.Save(&course)
	return responses.SendSuccess(c, "Course updated successfully", course)
}

// DeleteCourse deletes a course
func DeleteCourse(c *fiber.Ctx) error {
	id := c.Params("id")
	var course models.Course

	if err := config.DB.First(&course, id).Error; err != nil {
		return responses.SendError(c, fiber.StatusNotFound, "Course not found")
	}

	config.DB.Delete(&course)
	return responses.SendSuccess(c, "Course deleted successfully", nil)
}
