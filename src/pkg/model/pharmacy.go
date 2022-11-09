package model

import (
	"time"

	"gorm.io/gorm"
)

type Pharmacy struct {
	ID           uint   `gorm:"primaryKey"`
	ChatID       uint64 `json:"chat_id"`
	Chat         Chat
	ApothecaryID uint       `json:"-"`
	Apothecary   Apothecary `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ProcessAt    time.Time  `gorm:"default:null" json:"process_at"`
	SentAt       time.Time  `gorm:"default:null" json:"sent_at"`
	IsProcess    bool       `gorm:"type:bool;default:false" json:"is_process"`
	IsSent       bool       `gorm:"type:bool;default:false" json:"is_sent"`
	gorm.Model   `json:"-"`
}
