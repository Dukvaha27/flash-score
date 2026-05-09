package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model

	UserID uint `gorm:"not null;index"`

	EventID      *uint `gorm:"index"`
	CommentaryID *uint `gorm:"index"`

	Text string `gorm:"type:text;not null"`
}
