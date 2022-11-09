package model

import (
	"time"

	"gorm.io/gorm"
)

type Shipment struct {
	ID         uint64 `gorm:"primaryKey" json:"id"`
	ChatID     uint64 `json:"chat_id"`
	Chat       Chat
	CourierID  uint `json:"courier_id"`
	Courier    Courier
	TakeAt     time.Time `gorm:"default:null" json:"take_at"`
	DoneAt     time.Time `gorm:"default:null" json:"done_at"`
	IsTake     bool      `gorm:"type:bool;default:false" json:"is_take"`
	IsDone     bool      `gorm:"type:bool;default:false" json:"is_done"`
	gorm.Model `json:"-"`
}
