package models

import (
	"course-api/types"

	"gorm.io/gorm"
)

type ContentTopic struct {
	gorm.Model
	Title      string            `json:"title" validate:"required"`
	Topics     types.StringArray `json:"topics" gorm:"type:json" validate:"required"`
	MaterialID uint              `json:"material_id"`
	Material   *Material         `json:"-" gorm:"foreignKey:MaterialID"`
}

type VideoCourse struct {
	gorm.Model
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	YoutubeID   string    `json:"youtube_id" validate:"required"`
	Duration    string    `json:"duration" validate:"required"`
	Instructor  string    `json:"instructor" validate:"required"`
	Level       string    `json:"level" validate:"required,oneof=beginner intermediate advanced"`
	MaterialID  uint      `json:"material_id"`
	Material    *Material `json:"-" gorm:"foreignKey:MaterialID"`
}

type Material struct {
	gorm.Model
	Title          string              `json:"title" validate:"required"`
	Description    string              `json:"description" validate:"required"`
	Icon           string              `json:"icon" validate:"required"`
	Duration       int                 `json:"duration" validate:"required,min=1"`
	Lessons        int                 `json:"lessons" validate:"required,min=1"`
	LearningPoints types.LearningPoint `json:"learningPoints" gorm:"type:json" validate:"required"`
	Content        []ContentTopic      `json:"content" gorm:"foreignKey:MaterialID"`
	VideoCourses   []VideoCourse       `json:"videoCourses" gorm:"foreignKey:MaterialID"`
}

type CreateMaterialInput struct {
	Title          string         `json:"title" validate:"required"`
	Description    string         `json:"description" validate:"required"`
	Icon           string         `json:"icon" validate:"required"`
	Duration       int            `json:"duration" validate:"required,min=1"`
	Lessons        int            `json:"lessons" validate:"required,min=1"`
	LearningPoints []string       `json:"learningPoints" validate:"required"`
	Content        []ContentTopic `json:"content" validate:"required"`
	VideoCourses   []VideoCourse  `json:"videoCourses" validate:"required"`
}

type UpdateMaterialInput struct {
	Title          string         `json:"title"`
	Description    string         `json:"description"`
	Icon           string         `json:"icon"`
	Duration       int            `json:"duration" validate:"omitempty,min=1"`
	Lessons        int            `json:"lessons" validate:"omitempty,min=1"`
	LearningPoints []string       `json:"learningPoints"`
	Content        []ContentTopic `json:"content"`
	VideoCourses   []VideoCourse  `json:"videoCourses"`
}
