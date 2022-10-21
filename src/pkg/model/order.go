package model

import "gorm.io/gorm"

type Order struct {
	ID         uint64 `gorm:"primarykey" json:"id"`
	UserID     uint64 `json:"-"`
	User       User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DoctorID   uint   `json:"-"`
	Doctor     Doctor `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Number     string `json:"number"`
	Total      int    `json:"total"`
	Status     string `gorm:"size:25" json:"status"`
	gorm.Model `json:"-"`
}
