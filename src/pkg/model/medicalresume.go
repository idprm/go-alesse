package model

import "gorm.io/gorm"

type MedicalResume struct {
	ID         uint64 `gorm:"primaryKey" json:"id"`
	ChatID     uint64 `json:"chat_id"`
	Chat       Chat
	Number     string  `gorm:"size:25" json:"number"`
	DiseaseID  uint    `json:"disease_id"`
	Disease    Disease `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Diagnosis  string  `gorm:"type:text" json:"diagnosis"`
	gorm.Model `json:"-"`
}
