package model

import (
	"time"

	"gorm.io/gorm"
)

type PharmacyApothecary struct {
	ID         uint64 `gorm:"primaryKey" json:"id"`
	PharmacyID uint64 `json:"pharmacy_id"`
	Pharmacy   Pharmacy
	ProcessAt  time.Time `gorm:"default:null" json:"process_at"`
	SentAt     time.Time `gorm:"default:null" json:"sent_at"`
	IsProcess  bool      `gorm:"type:bool;default:false" json:"is_process"`
	IsSent     bool      `gorm:"type:bool;default:false" json:"is_sent"`
	gorm.Model `json:"-"`
}
