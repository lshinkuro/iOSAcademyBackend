package handlers

import (
	"course-api/config"
	"course-api/models"
	"course-api/responses"
	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// GetAllCourses returns all courses
func GetAllCourses(c *fiber.Ctx) error {
	var courses []models.Course
	config.DB.Find(&courses)
	return responses.SendSuccess(c, "Courses found successfully", courses)
}

// GetCourse returns a single course
func GetCourse(c *fiber.Ctx) error {
	id := c.Params("id")
	var course models.Course

	if err := config.DB.First(&course, id).Error; err != nil {
		return responses.SendError(c, fiber.StatusNotFound, "Course not found")
	}

	return responses.SendSuccess(c, "Course found successfully", course)
}

// CreateCourse creates a new course
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

	config.DB.Create(&course)
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