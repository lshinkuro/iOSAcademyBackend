package handlers

import (
	"course-api/config"
	"course-api/models"
	"course-api/responses"
	"course-api/types"
	"course-api/validator"

	"github.com/gofiber/fiber/v2"
)

// GetContentTopics godoc
// @Summary Get all content topics for a material
// @Description Retrieve all content topics for a specific material
// @Tags content
// @Accept json
// @Produce json
// @Param material_id path int true "Material ID"
// @Security ApiKeyAuth
// @Success 200 {object} responses.Response{data=[]models.ContentTopic}
// @Failure 401 {object} responses.Response
// @Failure 404 {object} responses.Response
// @Router /materials/{material_id}/content [get]
func GetContentTopics(c *fiber.Ctx) error {
	materialID := c.Params("material_id")
	var contentTopics []models.ContentTopic

	if err := config.DB.Where("material_id = ?", materialID).Order("order asc").Find(&contentTopics).Error; err != nil {
		return responses.SendError(c, fiber.StatusNotFound, "Content topics not found")
	}

	return responses.SendSuccess(c, "Content topics found successfully", contentTopics)
}

// GetContentTopic godoc
// @Summary Get a specific content topic
// @Description Retrieve a specific content topic by its ID
// @Tags content
// @Accept json
// @Produce json
// @Param id path int true "Content Topic ID"
// @Security ApiKeyAuth
// @Success 200 {object} responses.Response{data=models.ContentTopic}
// @Failure 401 {object} responses.Response
// @Failure 404 {object} responses.Response
// @Router /content/{id} [get]
func GetContentTopic(c *fiber.Ctx) error {
	id := c.Params("id")
	var contentTopic models.ContentTopic

	if err := config.DB.First(&contentTopic, id).Error; err != nil {
		return responses.SendError(c, fiber.StatusNotFound, "Content topic not found")
	}

	return responses.SendSuccess(c, "Content topic found successfully", contentTopic)
}

// CreateContentTopic godoc
// @Summary Create a new content topic
// @Description Create a new content topic with HTML content
// @Tags content
// @Accept json
// @Produce json
// @Param input body models.CreateContentTopicInput true "Content Topic creation details"
// @Security ApiKeyAuth
// @Success 200 {object} responses.Response{data=models.ContentTopic}
// @Failure 400 {object} responses.Response
// @Failure 401 {object} responses.Response
// @Router /content [post]
func CreateContentTopic(c *fiber.Ctx) error {
	input := new(models.CreateContentTopicInput)

	if err := c.BodyParser(input); err != nil {
		return responses.SendError(c, fiber.StatusBadRequest, "Invalid input")
	}

	if err := validator.Validate.Struct(input); err != nil {
		return responses.SendError(c, fiber.StatusBadRequest, err.Error())
	}

	contentTopic := models.ContentTopic{
		Title:      input.Title,
		Content:    input.Content,
		Topics:     types.StringArray(input.Topics),
		Order:      input.Order,
		MaterialID: input.MaterialID,
	}

	if err := config.DB.Create(&contentTopic).Error; err != nil {
		return responses.SendError(c, fiber.StatusInternalServerError, "Error creating content topic")
	}

	return responses.SendSuccess(c, "Content topic created successfully", contentTopic)
}

// UpdateContentTopic godoc
// @Summary Update a content topic
// @Description Update an existing content topic's details including HTML content
// @Tags content
// @Accept json
// @Produce json
// @Param id path int true "Content Topic ID"
// @Param input body models.UpdateContentTopicInput true "Content Topic update details"
// @Security ApiKeyAuth
// @Success 200 {object} responses.Response{data=models.ContentTopic}
// @Failure 400 {object} responses.Response
// @Failure 401 {object} responses.Response
// @Failure 404 {object} responses.Response
// @Router /content/{id} [put]
func UpdateContentTopic(c *fiber.Ctx) error {
	id := c.Params("id")
	input := new(models.UpdateContentTopicInput)

	if err := c.BodyParser(input); err != nil {
		return responses.SendError(c, fiber.StatusBadRequest, "Invalid input")
	}

	var contentTopic models.ContentTopic
	if err := config.DB.First(&contentTopic, id).Error; err != nil {
		return responses.SendError(c, fiber.StatusNotFound, "Content topic not found")
	}

	if input.Title != "" {
		contentTopic.Title = input.Title
	}
	if input.Content != "" {
		contentTopic.Content = input.Content
	}
	if len(input.Topics) > 0 {
		contentTopic.Topics = types.StringArray(input.Topics)
	}
	if input.Order != 0 {
		contentTopic.Order = input.Order
	}

	if err := config.DB.Save(&contentTopic).Error; err != nil {
		return responses.SendError(c, fiber.StatusInternalServerError, "Error updating content topic")
	}

	return responses.SendSuccess(c, "Content topic updated successfully", contentTopic)
}

// DeleteContentTopic godoc
// @Summary Delete a content topic
// @Description Delete a content topic by its ID
// @Tags content
// @Accept json
// @Produce json
// @Param id path int true "Content Topic ID"
// @Security ApiKeyAuth
// @Success 200 {object} responses.Response
// @Failure 401 {object} responses.Response
// @Failure 404 {object} responses.Response
// @Router /content/{id} [delete]
func DeleteContentTopic(c *fiber.Ctx) error {
	id := c.Params("id")
	var contentTopic models.ContentTopic

	if err := config.DB.First(&contentTopic, id).Error; err != nil {
		return responses.SendError(c, fiber.StatusNotFound, "Content topic not found")
	}

	if err := config.DB.Delete(&contentTopic).Error; err != nil {
		return responses.SendError(c, fiber.StatusInternalServerError, "Error deleting content topic")
	}

	return responses.SendSuccess(c, "Content topic deleted successfully", nil)
}
