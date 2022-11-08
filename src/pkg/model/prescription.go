package model

import "gorm.io/gorm"

type Prescription struct {
	ID              uint64 `gorm:"primaryKey" json:"id"`
	ChatID          uint64 `json:"chat_id"`
	Chat            Chat
	Number          string `gorm:"size:25" json:"number"`
	Slug            string `gorm:"size:50" json:"slug"`
	PainComplaint   string `gorm:"type:text" json:"pain_complaint"`
	Diagnosis       string `gorm:"type:text" json:"diagnosis"`
	AllergyMedicine string `gorm:"type:text" json:"allergy_medicine"`
	IsSubmited      bool   `gorm:"type:boolean;default:false" json:"is_submited"`
	gorm.Model      `json:"-"`
}