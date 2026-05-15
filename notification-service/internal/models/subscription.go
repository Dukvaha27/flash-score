package models

import (
	"time"

	"github.com/Dukvaha27/flash-score/notification-service/internal/errors"
)

type Subscription struct {
	ID        uint
	CreatedAt time.Time
	UserID    uint  `json:"user_id" gorm:"not null;uniqueIndex:idx_user_team;uniqueIndex:idx_user_sport"`
	TeamID    *uint `json:"team_id,omitempty" gorm:"uniqueIndex:idx_user_team"`
	SportID   *uint `json:"sport_id,omitempty" gorm:"uniqueIndex:idx_user_sport"`
}

type SubscriptionCreate struct {
	TeamID  *uint `json:"team_id,omitempty"`
	SportID *uint `json:"sport_id,omitempty"`
}

func (s *SubscriptionCreate) Validate() error {
	if s.TeamID == nil && s.SportID == nil {
		return errors.ErrSubscriptionTargetRequired
	}

	if s.TeamID != nil && s.SportID != nil {
		return errors.ErrTeamOrSportRequired
	}

	return nil
}
