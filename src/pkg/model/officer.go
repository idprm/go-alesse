package model

import "time"

type Officer struct {
	ID             uint         `gorm:"primaryKey" json:"id"`
	HealthcenterID uint         `json:"-"`
	Healthcenter   Healthcenter `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name           string       `gorm:"size:100;not null" json:"name"`
	Photo          string       `gorm:"size:150;not null" json:"photo"`
	Phone          string       `gorm:"size:15;unique" json:"phone"`
	LoginAt        time.Time    `gorm:"default:null" json:"-"`
	IpAddress      string       `gorm:"size:25" json:"ip_address"`
	IsActive       bool         `gorm:"type:bool" json:"is_active"`
}
