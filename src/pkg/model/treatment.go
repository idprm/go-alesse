package model

import "gorm.io/gorm"

type Treatment struct {
	ID         uint64 `gorm:"primaryKey" json:"id"`
	ChatID     uint64 `json:"chat_id"`
	Chat       Chat
	OfficerID  uint `json:"officer_id"`
	Officer    Officer
	Photo      string `gorm:"size:25" json:"photo"`
	gorm.Model `json:"-"`
}
