package model

import (
	"time"

	"gorm.io/gorm"
)

type Homecare struct {
	ID                 uint64 `gorm:"primaryKey" json:"id"`
	ChatID             uint64 `json:"chat_id"`
	Chat               Chat
	Number             string    `gorm:"size:25" json:"number"`
	EarlyDiagnosis     string    `gorm:"type:text" json:"early_diagnosis"`
	Reason             string    `gorm:"type:text" json:"reason"`
	VisitAt            time.Time `gorm:"default:null" json:"visit_at"`
	Slug               string    `gorm:"size:50;unique" json:"slug"`
	Treatment          string    `gorm:"type:text" json:"treatment"`
	FinalDiagnosis     string    `gorm:"type:text" json:"final_diagnosis"`
	DrugAdministration string    `gorm:"type:text" json:"drug_administration"`
	IsSubmited         bool      `gorm:"type:boolean;default:false" json:"is_submited"`
	IsVisited          bool      `gorm:"type:boolean;default:false" json:"is_visited"`
	IsFinished         bool      `gorm:"type:boolean;default:false" json:"is_finished"`
	IsActive           bool      `gorm:"type:boolean;default:false" json:"is_active"`
	gorm.Model         `json:"-"`
}
