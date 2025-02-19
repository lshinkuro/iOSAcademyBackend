package handlers

import (
	"course-api/config"
	"course-api/middleware"
	"course-api/models"
	"course-api/responses"
	"github.com/gofiber/fiber/v2"
)

func SignUp(c *fiber.Ctx) error {
	input := new(models.SignupInput)

	if err := c.BodyParser(input); err != nil {
		return responses.SendError(c, fiber.StatusBadRequest, "Invalid input")
	}

	if err := validate.Struct(input); err != nil {
		return responses.SendError(c, fiber.StatusBadRequest, err.Error())
	}

	// Check if user already exists
	var existingUser models.User
	if err := config.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		return responses.SendError(c, fiber.StatusBadRequest, "Email already registered")
	}

	user := models.User{
		Email:    input.Email,
		Password: input.Password,
		FullName: input.FullName,
	}

	if err := user.HashPassword(); err != nil {
		return responses.SendError(c, fiber.StatusInternalServerError, "Error creating user")
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return responses.SendError(c, fiber.StatusInternalServerError, "Error creating user")
	}

	token, err := middleware.GenerateToken(user.ID)
	if err != nil {
		return responses.SendError(c, fiber.StatusInternalServerError, "Error generating token")
	}

	return responses.SendSuccess(c, "User created successfully", fiber.Map{
		"token": token,
		"user": fiber.Map{
			"id":        user.ID,
			"email":     user.Email,
			"full_name": user.FullName,
		},
	})
}

func SignIn(c *fiber.Ctx) error {
	input := new(models.LoginInput)

	if err := c.BodyParser(input); err != nil {
		return responses.SendError(c, fiber.StatusBadRequest, "Invalid input")
	}

	if err := validate.Struct(input); err != nil {
		return responses.SendError(c, fiber.StatusBadRequest, err.Error())
	}

	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return responses.SendError(c, fiber.StatusUnauthorized, "Invalid credentials")
	}

	if err := user.CheckPassword(input.Password); err != nil {
		return responses.SendError(c, fiber.StatusUnauthorized, "Invalid credentials")
	}

	token, err := middleware.GenerateToken(user.ID)
	if err != nil {
		return responses.SendError(c, fiber.StatusInternalServerError, "Error generating token")
	}

	return responses.SendSuccess(c, "Login successful", fiber.Map{
		"token": token,
		"user": fiber.Map{
			"id":        user.ID,
			"email":     user.Email,
			"full_name": user.FullName,
		},
	})
}
