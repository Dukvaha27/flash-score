package models

import "gorm.io/gorm"

type Sport struct {
	gorm.Model
	Name string `json:"name" binding:"required"`
}

type SportAction struct {
	Name string `json:"name" binding:"required"`
}
