package handlers

import (
	"course-api/config"
	"course-api/middleware"
	"course-api/models"
	"course-api/responses"

	"github.com/gofiber/fiber/v2"
)

// SignUp godoc
// @Summary Register a new user
// @Description Create a new user account with the provided details
// @Tags auth
// @Accept json
// @Produce json
// @Param input body models.SignupInput true "User registration details"
// @Success 200 {object} responses.Response{data=map[string]interface{}}
// @Failure 400 {object} responses.Response
// @Failure 500 {object} responses.Response
// @Router /api/v1/auth/signup [post]
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
		Role:     input.Role,
	}

	if err := user.HashPassword(); err != nil {
		return responses.SendError(c, fiber.StatusInternalServerError, "Error creating user")
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return responses.SendError(c, fiber.StatusInternalServerError, "Error creating user")
	}

	token, err := middleware.GenerateToken(user.ID, user.Role)
	if err != nil {
		return responses.SendError(c, fiber.StatusInternalServerError, "Error generating token")
	}

	return responses.SendSuccess(c, "User created successfully", fiber.Map{
		"token": token,
		"user": fiber.Map{
			"id":        user.ID,
			"email":     user.Email,
			"full_name": user.FullName,
			"role":      user.Role,
		},
	})
}

// SignIn godoc
// @Summary Authenticate user
// @Description Authenticate user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param input body models.LoginInput true "User login credentials"
// @Success 200 {object} responses.Response{data=map[string]interface{}}
// @Failure 400 {object} responses.Response
// @Failure 401 {object} responses.Response
// @Failure 500 {object} responses.Response
// @Router /api/v1/auth/signin [post]
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

	token, err := middleware.GenerateToken(user.ID, user.Role)
	if err != nil {
		return responses.SendError(c, fiber.StatusInternalServerError, "Error generating token")
	}

	return responses.SendSuccess(c, "Login successful", fiber.Map{
		"token": token,
		"user": fiber.Map{
			"id":        user.ID,
			"email":     user.Email,
			"full_name": user.FullName,
			"role":      user.Role,
		},
	})
}
