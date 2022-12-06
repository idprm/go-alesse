package model

import "gorm.io/gorm"

type HomecareResume struct {
	ID         uint64 `gorm:"primaryKey" json:"id"`
	HomecareID uint64 `json:"homecare_id"`
	Homecare   Homecare

	gorm.Model
}
