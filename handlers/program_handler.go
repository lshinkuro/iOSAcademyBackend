package handlers

import (
	"course-api/config"
	"course-api/models"
	"course-api/responses"
	"course-api/types"
	"course-api/validator"

	"github.com/gofiber/fiber/v2"
)

// GetAllPrograms returns all programs
func GetAllPrograms(c *fiber.Ctx) error {
	var programs []models.Program
	config.DB.Find(&programs)
	return responses.SendSuccess(c, "Programs found successfully", programs)
}

// GetProgram returns a single program
func GetProgram(c *fiber.Ctx) error {
	id := c.Params("id")
	var program models.Program

	if err := config.DB.First(&program, id).Error; err != nil {
		return responses.SendError(c, fiber.StatusNotFound, "Program not found")
	}

	return responses.SendSuccess(c, "Program found successfully", program)
}

// CreateProgram creates a new program
func CreateProgram(c *fiber.Ctx) error {
	input := new(models.CreateProgramInput)

	if err := c.BodyParser(input); err != nil {
		return responses.SendError(c, fiber.StatusBadRequest, "Invalid input")
	}

	if err := validator.Validate.Struct(input); err != nil {
		return responses.SendError(c, fiber.StatusBadRequest, err.Error())
	}

	program := models.Program{
		Title:    input.Title,
		Type:     input.Type,
		Duration: input.Duration,
		Price:    input.Price,
		Features: types.StringArray(input.Features),
	}

	if err := config.DB.Create(&program).Error; err != nil {
		return responses.SendError(c, fiber.StatusInternalServerError, "Error creating program")
	}

	return responses.SendSuccess(c, "Program created successfully", program)
}

// UpdateProgram updates an existing program
func UpdateProgram(c *fiber.Ctx) error {
	id := c.Params("id")
	input := new(models.UpdateProgramInput)

	if err := c.BodyParser(input); err != nil {
		return responses.SendError(c, fiber.StatusBadRequest, "Invalid input")
	}

	var program models.Program
	if err := config.DB.First(&program, id).Error; err != nil {
		return responses.SendError(c, fiber.StatusNotFound, "Program not found")
	}

	if input.Title != "" {
		program.Title = input.Title
	}
	if input.Type != "" {
		program.Type = input.Type
	}
	if input.Duration != "" {
		program.Duration = input.Duration
	}
	if input.Price != 0 {
		program.Price = input.Price
	}
	if len(input.Features) > 0 {
		program.Features = types.StringArray(input.Features)
	}

	if err := config.DB.Save(&program).Error; err != nil {
		return responses.SendError(c, fiber.StatusInternalServerError, "Error updating program")
	}

	return responses.SendSuccess(c, "Program updated successfully", program)
}

// DeleteProgram deletes a program
func DeleteProgram(c *fiber.Ctx) error {
	id := c.Params("id")
	var program models.Program

	if err := config.DB.First(&program, id).Error; err != nil {
		return responses.SendError(c, fiber.StatusNotFound, "Program not found")
	}

	if err := config.DB.Delete(&program).Error; err != nil {
		return responses.SendError(c, fiber.StatusInternalServerError, "Error deleting program")
	}

	return responses.SendSuccess(c, "Program deleted successfully", nil)
}
