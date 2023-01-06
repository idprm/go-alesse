package model

import (
	"time"

	"gorm.io/gorm"
)

type Homecare struct {
	ID             uint64 `gorm:"primaryKey" json:"id"`
	ChatID         uint64 `json:"chat_id"`
	Chat           Chat
	Number         string    `gorm:"size:25" json:"number"`
	PainComplaints string    `gorm:"type:text" json:"pain_complaints"`
	EarlyDiagnosis string    `gorm:"type:text" json:"early_diagnosis"`
	VisitAt        time.Time `gorm:"default:null" json:"visit_at"`
	Slug           string    `gorm:"size:200;unique" json:"slug"`
	SubmitedAt     time.Time `gorm:"default:null" json:"submited_at"`
	IsSoon         bool      `gorm:"type:bool;default:true" json:"is_soon"`
	IsSubmited     bool      `gorm:"type:bool;default:false" json:"is_submited"`
	IsActive       bool      `gorm:"type:bool;default:true" json:"is_active"`
	gorm.Model     `json:"-"`
}
