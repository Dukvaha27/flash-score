package models

import "gorm.io/gorm"

type Commentary struct {
	gorm.Model

	MatchID uint `gorm:"not null;index"`
	Minute  int  `gorm:"not null"`

	Text     string `gorm:"type:text;not null"`
	IsPinned bool   `gorm:"not null;default:false"`

	CreatedBy uint `gorm:"not null;index"`
}
