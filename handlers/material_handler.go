package handlers

import (
	"course-api/config"
	"course-api/models"
	"course-api/responses"
	"course-api/types"
	"course-api/validator"

	"github.com/gofiber/fiber/v2"
)

// GetAllMaterials returns all materials with their related content and video courses
func GetAllMaterials(c *fiber.Ctx) error {
	var materials []models.Material
	if err := config.DB.Preload("Content").Preload("VideoCourses").Find(&materials).Error; err != nil {
		return responses.SendError(c, fiber.StatusInternalServerError, "Error fetching materials")
	}
	return responses.SendSuccess(c, "Materials found successfully", materials)
}

// GetMaterial returns a single material with its related content and video courses
func GetMaterial(c *fiber.Ctx) error {
	id := c.Params("id")
	var material models.Material

	if err := config.DB.Preload("Content").Preload("VideoCourses").First(&material, id).Error; err != nil {
		return responses.SendError(c, fiber.StatusNotFound, "Material not found")
	}

	return responses.SendSuccess(c, "Material found successfully", material)
}

// CreateMaterial creates a new material with its related content and video courses
func CreateMaterial(c *fiber.Ctx) error {
	input := new(models.CreateMaterialInput)

	if err := c.BodyParser(input); err != nil {
		return responses.SendError(c, fiber.StatusBadRequest, "Invalid input")
	}

	if err := validator.Validate.Struct(input); err != nil {
		return responses.SendError(c, fiber.StatusBadRequest, err.Error())
	}

	// Begin transaction
	tx := config.DB.Begin()

	material := models.Material{
		Title:          input.Title,
		Description:    input.Description,
		Icon:           input.Icon,
		Duration:       input.Duration,
		Lessons:        input.Lessons,
		LearningPoints: types.LearningPoint(input.LearningPoints),
	}

	if err := tx.Create(&material).Error; err != nil {
		tx.Rollback()
		return responses.SendError(c, fiber.StatusInternalServerError, "Error creating material")
	}

	// Create associated content topics
	for _, contentInput := range input.Content {
		content := models.ContentTopic{
			Title:      contentInput.Title,
			Topics:     contentInput.Topics,
			MaterialID: material.ID,
		}
		if err := tx.Create(&content).Error; err != nil {
			tx.Rollback()
			return responses.SendError(c, fiber.StatusInternalServerError, "Error creating content")
		}
	}

	// Create associated video courses
	for _, courseInput := range input.VideoCourses {
		course := models.VideoCourse{
			Title:       courseInput.Title,
			Description: courseInput.Description,
			YoutubeID:   courseInput.YoutubeID,
			Duration:    courseInput.Duration,
			Instructor:  courseInput.Instructor,
			Level:       courseInput.Level,
			MaterialID:  material.ID,
		}
		if err := tx.Create(&course).Error; err != nil {
			tx.Rollback()
			return responses.SendError(c, fiber.StatusInternalServerError, "Error creating video course")
		}
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return responses.SendError(c, fiber.StatusInternalServerError, "Error saving material")
	}

	// Fetch the complete material with associations
	var completeMaterial models.Material
	if err := config.DB.Preload("Content").Preload("VideoCourses").First(&completeMaterial, material.ID).Error; err != nil {
		return responses.SendError(c, fiber.StatusInternalServerError, "Error fetching created material")
	}

	return responses.SendSuccess(c, "Material created successfully", completeMaterial)
}

// UpdateMaterial updates an existing material and its related content and video courses
func UpdateMaterial(c *fiber.Ctx) error {
	id := c.Params("id")
	input := new(models.UpdateMaterialInput)

	if err := c.BodyParser(input); err != nil {
		return responses.SendError(c, fiber.StatusBadRequest, "Invalid input")
	}

	var material models.Material
	if err := config.DB.Preload("Content").Preload("VideoCourses").First(&material, id).Error; err != nil {
		return responses.SendError(c, fiber.StatusNotFound, "Material not found")
	}

	// Begin transaction
	tx := config.DB.Begin()

	if input.Title != "" {
		material.Title = input.Title
	}
	if input.Description != "" {
		material.Description = input.Description
	}
	if input.Icon != "" {
		material.Icon = input.Icon
	}
	if input.Duration != 0 {
		material.Duration = input.Duration
	}
	if input.Lessons != 0 {
		material.Lessons = input.Lessons
	}
	if len(input.LearningPoints) > 0 {
		material.LearningPoints = types.LearningPoint(input.LearningPoints)
	}

	if err := tx.Save(&material).Error; err != nil {
		tx.Rollback()
		return responses.SendError(c, fiber.StatusInternalServerError, "Error updating material")
	}

	// Update content topics if provided
	if len(input.Content) > 0 {
		// Delete existing content
		if err := tx.Where("material_id = ?", material.ID).Delete(&models.ContentTopic{}).Error; err != nil {
			tx.Rollback()
			return responses.SendError(c, fiber.StatusInternalServerError, "Error updating content")
		}

		// Create new content
		for _, contentInput := range input.Content {
			content := models.ContentTopic{
				Title:      contentInput.Title,
				Topics:     contentInput.Topics,
				MaterialID: material.ID,
			}
			if err := tx.Create(&content).Error; err != nil {
				tx.Rollback()
				return responses.SendError(c, fiber.StatusInternalServerError, "Error creating content")
			}
		}
	}

	// Update video courses if provided
	if len(input.VideoCourses) > 0 {
		// Delete existing video courses
		if err := tx.Where("material_id = ?", material.ID).Delete(&models.VideoCourse{}).Error; err != nil {
			tx.Rollback()
			return responses.SendError(c, fiber.StatusInternalServerError, "Error updating video courses")
		}

		// Create new video courses
		for _, courseInput := range input.VideoCourses {
			course := models.VideoCourse{
				Title:       courseInput.Title,
				Description: courseInput.Description,
				YoutubeID:   courseInput.YoutubeID,
				Duration:    courseInput.Duration,
				Instructor:  courseInput.Instructor,
				Level:       courseInput.Level,
				MaterialID:  material.ID,
			}
			if err := tx.Create(&course).Error; err != nil {
				tx.Rollback()
				return responses.SendError(c, fiber.StatusInternalServerError, "Error creating video course")
			}
		}
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return responses.SendError(c, fiber.StatusInternalServerError, "Error saving material")
	}

	// Fetch the updated material with associations
	var updatedMaterial models.Material
	if err := config.DB.Preload("Content").Preload("VideoCourses").First(&updatedMaterial, material.ID).Error; err != nil {
		return responses.SendError(c, fiber.StatusInternalServerError, "Error fetching updated material")
	}

	return responses.SendSuccess(c, "Material updated successfully", updatedMaterial)
}

// DeleteMaterial deletes a material and its related content and video courses
func DeleteMaterial(c *fiber.Ctx) error {
	id := c.Params("id")
	var material models.Material

	if err := config.DB.First(&material, id).Error; err != nil {
		return responses.SendError(c, fiber.StatusNotFound, "Material not found")
	}

	// Begin transaction
	tx := config.DB.Begin()

	// Delete related content topics
	if err := tx.Where("material_id = ?", material.ID).Delete(&models.ContentTopic{}).Error; err != nil {
		tx.Rollback()
		return responses.SendError(c, fiber.StatusInternalServerError, "Error deleting content")
	}

	// Delete related video courses
	if err := tx.Where("material_id = ?", material.ID).Delete(&models.VideoCourse{}).Error; err != nil {
		tx.Rollback()
		return responses.SendError(c, fiber.StatusInternalServerError, "Error deleting video courses")
	}

	// Delete the material
	if err := tx.Delete(&material).Error; err != nil {
		tx.Rollback()
		return responses.SendError(c, fiber.StatusInternalServerError, "Error deleting material")
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return responses.SendError(c, fiber.StatusInternalServerError, "Error completing deletion")
	}

	return responses.SendSuccess(c, "Material deleted successfully", nil)
}
