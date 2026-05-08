package models

import "gorm.io/gorm"

type PlayerPosition string

const (
	PlayerPositionForward    PlayerPosition = "forward"
	PlayerPositionMidfielder PlayerPosition = "midfielder"
	PlayerPositionDefender   PlayerPosition = "defender"
	PlayerPositionGoalkeeper PlayerPosition = "goalkeeper"
)

type Player struct {
	gorm.Model
	Name     string         `json:"name" binding:"required"`
	Number   uint           `json:"number" binding:"required"`
	Position PlayerPosition `json:"position" binding:"required,oneof=forward midfielder defender goalkeeper"`
	TeamID   uint           `json:"team_id" binding:"required"`
	Team     Team           `json:"-" gorm:"foreignKey:TeamID"`
}

type PlayerCreate struct {
	Name     string         `json:"name" binding:"required"`
	Number   uint           `json:"number" binding:"required"`
	Position PlayerPosition `json:"position" binding:"required,oneof=forward midfielder defender goalkeeper"`
	TeamID   uint           `json:"team_id" binding:"required"`
}

type PlayerUpdate struct {
	Name     *string         `json:"name"`
	Number   *uint           `json:"number"`
	Position *PlayerPosition `json:"position" binding:"omitempty,oneof=forward midfielder defender goalkeeper"`
	TeamID   *uint           `json:"team_id"`
}
