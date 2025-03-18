package handlers

import (
	"course-api/config"
	"course-api/models"
	"course-api/responses"
	"course-api/types"
	"course-api/validator"

	"github.com/gofiber/fiber/v2"
)

// GetAllPrograms godoc
// @Summary Get all programs
// @Description Retrieve all programs from the system
// @Tags programs
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} responses.Response{data=[]models.Program}
// @Failure 401 {object} responses.Response
// @Failure 500 {object} responses.Response
// @Router /programs [get]
// GetAllPrograms returns all programs
func GetAllPrograms(c *fiber.Ctx) error {
	var programs []models.Program
	config.DB.Find(&programs)
	return responses.SendSuccess(c, "Programs found successfully", programs)
}

// GetProgram godoc
// @Summary Get a program by ID
// @Description Retrieve a specific program by its ID
// @Tags programs
// @Accept json
// @Produce json
// @Param id path int true "Program ID"
// @Security ApiKeyAuth
// @Success 200 {object} responses.Response{data=models.Program}
// @Failure 401 {object} responses.Response
// @Failure 404 {object} responses.Response
// @Router /programs/{id} [get]
// GetProgram returns a single program
func GetProgram(c *fiber.Ctx) error {
	id := c.Params("id")
	var program models.Program

	if err := config.DB.First(&program, id).Error; err != nil {
		return responses.SendError(c, fiber.StatusNotFound, "Program not found")
	}

	return responses.SendSuccess(c, "Program found successfully", program)
}

// CreateProgram godoc
// @Summary Create a new program
// @Description Create a new program with the provided details
// @Tags programs
// @Accept json
// @Produce json
// @Param input body models.CreateProgramInput true "Program creation details"
// @Security ApiKeyAuth
// @Success 200 {object} responses.Response{data=models.Program}
// @Failure 400 {object} responses.Response
// @Failure 401 {object} responses.Response
// @Failure 500 {object} responses.Response
// @Router /programs [post]
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

// UpdateProgram godoc
// @Summary Update a program
// @Description Update an existing program's details
// @Tags programs
// @Accept json
// @Produce json
// @Param id path int true "Program ID"
// @Param input body models.UpdateProgramInput true "Program update details"
// @Security ApiKeyAuth
// @Success 200 {object} responses.Response{data=models.Program}
// @Failure 400 {object} responses.Response
// @Failure 401 {object} responses.Response
// @Failure 404 {object} responses.Response
// @Router /programs/{id} [put]
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

// DeleteProgram godoc
// @Summary Delete a program
// @Description Delete a program by its ID
// @Tags programs
// @Accept json
// @Produce json
// @Param id path int true "Program ID"
// @Security ApiKeyAuth
// @Success 200 {object} responses.Response
// @Failure 401 {object} responses.Response
// @Failure 404 {object} responses.Response
// @Router /programs/{id} [delete]
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
