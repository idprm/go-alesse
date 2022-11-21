package model

import "gorm.io/gorm"

type PharmacyUpload struct {
	ID             uint64       `gorm:"primaryKey" json:"id"`
	HealthcenterID uint         `json:"-"`
	Healthcenter   Healthcenter `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PharmacyID     uint64       `json:"pharmacy_id"`
	Pharmacy       Pharmacy
	Filename       string `gorm:"size:45" json:"filename"`
	gorm.Model     `json:"-"`
}
