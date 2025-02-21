package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Role string

const (
	RoleAdmin   Role = "admin"
	RoleStudent Role = "student"
	RoleMentor  Role = "mentor"
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique" validate:"required,email"`
	Password string `json:"-" validate:"required,min=6"`
	FullName string `json:"full_name" validate:"required"`
	Role     Role   `json:"role" gorm:"type:varchar(10);default:'student'" validate:"required,oneof=admin student mentor"`
}

type SignupInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	FullName string `json:"full_name" validate:"required"`
	Role     Role   `json:"role" validate:"required,oneof=admin student mentor"`
}

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// HashPassword hashes the user's password
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword verifies the provided password against the hashed password
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
