package model

import "gorm.io/gorm"

type MedicalResume struct {
	ID             uint64 `gorm:"primaryKey" json:"id"`
	ChatID         uint64 `json:"chat_id"`
	Chat           Chat
	Number         string  `gorm:"size:25" json:"number"`
	Slug           string  `gorm:"size:200;unique" json:"slug"`
	DiseaseID      uint    `json:"disease_id"`
	Disease        Disease `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Weight         uint    `gorm:"size:5" json:"weight"`
	PainComplaints string  `gorm:"type:text" json:"pain_complaints"`
	Diagnosis      string  `gorm:"type:text" json:"diagnosis"`
	IsSubmited     bool    `gorm:"type:bool;default:false" json:"is_submited"`
	gorm.Model     `json:"-"`
}
