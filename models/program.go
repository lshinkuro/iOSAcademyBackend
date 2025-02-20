package models

import (
	"course-api/types"

	"gorm.io/gorm"
)

type Program struct {
	gorm.Model
	Title    string            `json:"title" validate:"required"`
	Type     string            `json:"type" validate:"required,oneof=regular intensive"`
	Duration string            `json:"duration" validate:"required"`
	Price    float64           `json:"price" validate:"required,min=0"`
	Features types.StringArray `json:"features" gorm:"type:json" validate:"required,min=1"`
}

type CreateProgramInput struct {
	Title    string   `json:"title" validate:"required"`
	Type     string   `json:"type" validate:"required,oneof=regular intensive"`
	Duration string   `json:"duration" validate:"required"`
	Price    float64  `json:"price" validate:"required,min=0"`
	Features []string `json:"features" validate:"required,min=1"`
}

type UpdateProgramInput struct {
	Title    string   `json:"title"`
	Type     string   `json:"type" validate:"omitempty,oneof=regular intensive"`
	Duration string   `json:"duration"`
	Price    float64  `json:"price" validate:"omitempty,min=0"`
	Features []string `json:"features" validate:"omitempty,min=1"`
}
