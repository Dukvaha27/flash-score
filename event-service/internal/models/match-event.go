package models

import (
	"gorm.io/gorm"
)

type EventType string

const (
	Goal         EventType = "goal"
	YellowCard   EventType = "yellow_card"
	RedCard      EventType = "red_card"
	Substitution EventType = "substitution"
	Penalty      EventType = "penalty"
	VarDecision  EventType = "var_decision"
	HalfStart    EventType = "half_start"
	HalfEnd      EventType = "half_end"
	FullTime     EventType = "full_time"
	Injury       EventType = "injury"
	Timeout      EventType = "timeout"
)

type MatchEvent struct {
	gorm.Model

	MatchID   uint      `gorm:"not null;index"`
	EventType EventType `gorm:"type:varchar(32);not null;index"`
	Minute    int       `gorm:"not null"`

	TeamID   *uint `gorm:"index"`
	PlayerID *uint `gorm:"index"`

	Text      string `gorm:"type:text"`
	CreatedBy uint   `gorm:"not null;index"`
}
