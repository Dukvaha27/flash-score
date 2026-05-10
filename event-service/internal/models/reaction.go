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

	UserID uint `gorm:"not null;index;uniqueIndex:idx_reaction_user_event,where:event_id IS NOT NULL AND deleted_at IS NULL;uniqueIndex:idx_reaction_user_commentary,where:commentary_id IS NOT NULL AND deleted_at IS NULL"`

	EventID *uint `gorm:"index;uniqueIndex:idx_reaction_user_event,where:event_id IS NOT NULL AND deleted_at IS NULL;check:reaction_target_check,(event_id IS NOT NULL AND commentary_id IS NULL) OR (event_id IS NULL AND commentary_id IS NOT NULL)"`

	CommentaryID *uint `gorm:"index;uniqueIndex:idx_reaction_user_commentary,where:commentary_id IS NOT NULL AND deleted_at IS NULL"`

	ReactionsType ReactionsType `gorm:"type:varchar(16);not null"`
}
