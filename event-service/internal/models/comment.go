package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model

	UserID uint `gorm:"not null;index"`

	EventID *uint `gorm:"index;check:comment_target_check,(event_id IS NOT NULL AND commentary_id IS NULL) OR (event_id IS NULL AND commentary_id IS NOT NULL)"`

	CommentaryID *uint `gorm:"index"`

	Text string `gorm:"type:text;not null"`
}
