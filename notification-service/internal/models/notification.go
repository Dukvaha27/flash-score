package models

import (
	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	UserID  uint   `json:"user_id"`
	Message string `json:"message"`
	IsRead  bool   `json:"is_read"`
}

type NotificationCreate struct {
	Message string `json:"message" binding:"required"`
}
