package model

import (
	"time"

	"gorm.io/gorm"
)

type Chat struct {
	ID             uint64       `gorm:"primaryKey" json:"id"`
	HealthcenterID uint         `json:"healthcenter_id"`
	Healthcenter   Healthcenter `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OrderID        uint64       `json:"order_id"`
	Order          Order        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DoctorID       uint         `json:"-"`
	Doctor         Doctor       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID         uint64       `json:"user_id"`
	User           User         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ChannelName    string       `gorm:"size:200" json:"channel_name"`
	ChannelUrl     string       `gorm:"size:200;unique" json:"channel_url"`
	LeaveAt        time.Time    `gorm:"default:null" json:"-"`
	IsLeave        bool         `gorm:"type:bool;default:false" json:"is_leave"`
	LatestStatus   string       `gorm:"size:250;default:null" json:"latest_status"`
	LatestLabel    string       `gorm:"size:25;default:null" json:"latest_label"`
	// Category       []Category   `gorm:"many2many:chat_categories;"`
	// Disease        []Disease    `gorm:"many2many:chat_diseases;"`
	gorm.Model `json:"-"`
}
