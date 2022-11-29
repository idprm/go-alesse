package model

import "gorm.io/gorm"

type PharmacyUpload struct {
	ID         uint64 `gorm:"primaryKey" json:"id"`
	PharmacyID uint64 `json:"pharmacy_id"`
	Pharmacy   Pharmacy
	Filename   string `gorm:"size:45" json:"filename"`
	gorm.Model `json:"-"`
}
