package model

import (
	"time"

	"gorm.io/gorm"
)

type Pharmacy struct {
	ID              uint64 `gorm:"primaryKey" json:"id"`
	ChatID          uint64 `json:"chat_id"`
	Chat            Chat
	Number          string    `gorm:"size:25" json:"number"`
	Weight          uint32    `gorm:"size:5" json:"weight"`
	PainComplaints  string    `gorm:"type:text" json:"pain_complaints"`
	DiseaseID       uint      `json:"disease_id"`
	Disease         Disease   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Diagnosis       string    `gorm:"type:text" json:"diagnosis"`
	AllergyMedicine string    `gorm:"type:text" json:"allergy_medicine"`
	Slug            string    `gorm:"size:200;unique" json:"slug"`
	SubmitedAt      time.Time `gorm:"default:null" json:"submited_at"`
	IsSubmited      bool      `gorm:"type:bool;default:false" json:"is_submited"`
	IsActive        bool      `gorm:"type:bool;default:true" json:"is_active"`
	gorm.Model      `json:"-"`
}
