package model

import (
	"time"
)

type Doctor struct {
	ID                   uint         `gorm:"primaryKey" json:"id"`
	HealthcenterID       uint         `json:"-"`
	Healthcenter         Healthcenter `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Username             string       `gorm:"size:25;unique;not null" json:"username"`
	Name                 string       `gorm:"size:100;not null" json:"name"`
	Photo                string       `gorm:"size:150;not null" json:"photo"`
	Type                 string       `gorm:"size:200" json:"type"`
	Number               string       `gorm:"size:100" json:"number"`
	Experience           int          `gorm:"size:2" json:"experience"`
	GraduatedFrom        string       `gorm:"size:150" json:"graduated_from"`
	ConsultationSchedule string       `gorm:"size:250" json:"consultation_schedule"`
	PlacePractice        string       `gorm:"size:250" json:"place_practice"`
	Phone                string       `gorm:"size:15" json:"phone"`
	Start                time.Time    `gorm:"type:time" json:"start"`
	End                  time.Time    `gorm:"type:time" json:"end"`
	IsActive             bool         `gorm:"type:bool" json:"is_active"`
}
