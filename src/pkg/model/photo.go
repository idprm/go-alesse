package model

import "gorm.io/gorm"

type Photo struct {
	ID         uint64 `gorm:"primaryKey" json:"id"`
	HomecareID uint64 `json:"homecare_id"`
	Homecare   Homecare
	FileName   string `gorm:"size:45" json:"file_name"`
	gorm.Model `json:"-"`
}
