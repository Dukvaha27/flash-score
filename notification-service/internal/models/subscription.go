package models

import (
	"errors"

	"gorm.io/gorm"
)

type Subscription struct {
	gorm.Model
	UserID  uint  `json:"user_id" gorm:"not null;uniqueIndex:idx_user_team;uniqueIndex:idx_user_sport"`
	TeamID  *uint `json:"team_id,omitempty" gorm:"index;uniqueIndex:idx_user_team"`
	SportID *uint `json:"sport_id,omitempty" gorm:"index;uniqueIndex:idx_user_sport"`
}

func (s *Subscription) Validate() error {
	if s.TeamID == nil && s.SportID == nil {
		return errors.New("Необходимо указать объект подписки")
	}

	if s.TeamID != nil && s.SportID != nil {
		return errors.New("Подписка может содержать лишь один объект")
	}

	return nil
}
