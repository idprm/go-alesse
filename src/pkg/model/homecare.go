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
	Reason         string    `gorm:"type:text" json:"reason"`
	VisitAt        time.Time `gorm:"default:null" json:"visit_at"`
	Slug           string    `gorm:"size:50;unique" json:"slug"`
	SubmitedAt     time.Time `gorm:"default:null" json:"submited_at"`
	FinishedAt     time.Time `gorm:"default:null" json:"finished_at"`
	IsSubmited     bool      `gorm:"type:bool;default:false" json:"is_submited"`
	IsFinished     bool      `gorm:"type:bool;default:false" json:"is_finished"`
	IsActive       bool      `gorm:"type:bool;default:true" json:"is_active"`
	gorm.Model     `json:"-"`
}
