package model

import (
	"time"

	"gorm.io/gorm"
)

type PharmacyCourier struct {
	ID         uint64 `gorm:"primaryKey" json:"id"`
	PharmacyID uint64 `json:"pharmacy_id"`
	Pharmacy   Pharmacy
	TakeAt     time.Time `gorm:"default:null" json:"take_at"`
	FinishAt   time.Time `gorm:"default:null" json:"finish_at"`
	IsTake     bool      `gorm:"type:bool;default:false" json:"is_take"`
	IsFinish   bool      `gorm:"type:bool;default:false" json:"is_finish"`
	ReceivedBy string    `gorm:"size:200;default:null" json:"received_by"`
	gorm.Model `json:"-"`
}
