package models

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	Name      string `json:"name" binding:"required"`
	ShortName string `json:"short_name" binding:"required"`
	City      string `json:"city" binding:"required"`
	SportID   uint   `json:"sport_id" binding:"required"`
	Sport     Sport  `json:"-" gorm:"foreignKey:SportID"`
}

type TeamCreate struct {
	Name      string `json:"name" binding:"required"`
	ShortName string `json:"short_name" binding:"required"`
	City      string `json:"city" binding:"required"`
	SportID   uint   `json:"sport_id" binding:"required"`
}

type TeamUpdate struct {
	Name      *string `json:"name"`
	ShortName *string `json:"short_name"`
	City      *string `json:"city"`
	SportID   *uint   `json:"sport_id"`
}
