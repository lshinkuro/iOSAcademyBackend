package models

import (
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model        // This embeds gorm.Model which includes ID, CreatedAt, UpdatedAt, and DeletedAt
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Instructor  string    `json:"instructor" validate:"required"`
	Duration    int       `json:"duration" validate:"required,min=1"`
	Price       float64   `json:"price" validate:"required,min=0"`
}

type CreateCourseInput struct {
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Instructor  string  `json:"instructor" validate:"required"`
	Duration    int     `json:"duration" validate:"required,min=1"`
	Price       float64 `json:"price" validate:"required,min=0"`
}

type UpdateCourseInput struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Instructor  string  `json:"instructor"`
	Duration    int     `json:"duration" validate:"min=1"`
	Price       float64 `json:"price" validate:"min=0"`
}