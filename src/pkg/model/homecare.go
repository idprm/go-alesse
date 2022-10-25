package model

import (
	"time"

	"github.com/idprm/go-alesse/src/pkg/common"
)

type Homecare struct {
	common.Model
	ChatID             uint64 `json:"chat_id"`
	Chat               Chat
	EarlyDiagnosis     string    `gorm:"type:text" json:"early_diagnosis"`
	Reason             string    `gorm:"type:text" json:"reason"`
	VisitAt            time.Time `gorm:"default:null" json:"visit_at"`
	Slug               string    `gorm:"size:50" json:"slug"`
	Treatment          string    `gorm:"type:text" json:"treatment"`
	FinalDiagnosis     string    `gorm:"type:text" json:"final_diagnosis"`
	DrugAdministration string    `gorm:"type:text" json:"drug_administration"`
	IsVisited          bool      `gorm:"type:boolean;default:false" json:"is_visited"`
	IsFinished         bool      `gorm:"type:boolean;default:false" json:"is_finished"`
	IsActive           bool      `gorm:"type:boolean;default:false" json:"is_active"`
}
