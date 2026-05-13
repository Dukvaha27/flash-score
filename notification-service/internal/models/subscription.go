package models

import "gorm.io/gorm"

type Subscription struct {
	gorm.Model
	UserID  uint `json:"user_id"`
	TeamID  uint `json:"team_id"`
	SportID uint `json:"sport_id"`
}
