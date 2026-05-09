package models

import "gorm.io/gorm"

type ReactionsType string

const (
	Like  ReactionsType = "like"
	Fire  ReactionsType = "fire"
	Shock ReactionsType = "shock"
	Sad   ReactionsType = "sad"
	Laugh ReactionsType = "laugh"
)

type Reaction struct {
	gorm.Model

	UserID uint `gorm:"not null;index"`

	EventID      *uint `gorm:"index"`
	CommentaryID *uint `gorm:"index"`

	ReactionsType ReactionsType `gorm:"type:varchar(16);not null"`
}
