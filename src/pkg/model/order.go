package model

import "gorm.io/gorm"

type Order struct {
	ID             uint64       `gorm:"primaryKey" json:"id"`
	HealthcenterID uint         `json:"healthcenter_id"`
	Healthcenter   Healthcenter `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID         uint64       `json:"-"`
	User           User         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DoctorID       uint         `json:"-"`
	Doctor         Doctor       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Number         string       `gorm:"size:25" json:"number"`
	Total          float32      `gorm:"size:8" json:"total"`
	Status         string       `gorm:"size:25" json:"status"`
	gorm.Model     `json:"-"`
}
