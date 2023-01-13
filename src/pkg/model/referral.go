package model

import (
	"time"

	"gorm.io/gorm"
)

type Referral struct {
	ID           uint64     `gorm:"primaryKey" json:"id"`
	ChatID       uint64     `json:"-"`
	Chat         Chat       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DoctorID     uint       `json:"-"`
	Doctor       Doctor     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SpecialistID uint       `json:"-"`
	Specialist   Specialist `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ChannelName  string     `gorm:"size:200" json:"channel_name"`
	ChannelUrl   string     `gorm:"size:200" json:"channel_url"`
	LeaveAt      time.Time  `gorm:"default:null" json:"-"`
	IsLeave      bool       `gorm:"type:bool;default:false" json:"is_leave"`
	gorm.Model   `json:"-"`
}
