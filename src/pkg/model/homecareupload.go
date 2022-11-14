package model

import "gorm.io/gorm"

type HomecareUpload struct {
	ID             uint64       `gorm:"primaryKey" json:"id"`
	HealthcenterID uint         `json:"-"`
	Healthcenter   Healthcenter `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	HomecareID     uint64       `json:"homecare_id"`
	Homecare       Homecare
	Filename       string `gorm:"size:45" json:"filename"`
	gorm.Model     `json:"-"`
}
