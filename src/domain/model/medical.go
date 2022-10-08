package model

import "gorm.io/gorm"

type Medical struct {
	ID         uint64 `gorm:"primaryKey" json:"id"`
	gorm.Model `json:"-"`
}
