package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName      string `json:"full_name" gorm:"not null"`
	Email         string `json:"email" gorm:"not null;uniqueIndex"`
	FavoriteSport string `json:"favorite_sport"`
	PasswordHash  string `json:"-"`
	Role          string `json:"role" gorm:"not null"`
}

type UserCreate struct {
	FullName      string `json:"full_name" binding:"required"`
	Email         string `json:"email" binding:"required,email"`
	FavoriteSport string `json:"favorite_sport"`
	Role          string `json:"role" binding:"required"`
}

type UserUpdate struct {
	FullName      *string `json:"full_name"`
	Email         *string `json:"email" binding:"omitempty,email"`
	FavoriteSport *string `json:"favorite_sport"`
	Role          *string `json:"role" binding:"omitempty"`
}

type RegisterRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"required,min=3"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
