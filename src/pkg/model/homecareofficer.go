package model

import (
	"time"

	"gorm.io/gorm"
)

type HomecareOfficer struct {
	ID                   uint64 `gorm:"primaryKey" json:"id"`
	HomecareID           uint64 `json:"homecare_id"`
	Homecare             Homecare
	Slug                 string `gorm:"size:50;unique" json:"slug"`
	PhysicaleExamination string `gorm:"type:text" json:"physical_examination"`
	MedicalTreatment     string `gorm:"type:text" json:"medical_treatment"`
	FinalDiagnosis       string `gorm:"type:text" json:"final_diagnosis"`
	DoctorID             uint   `gorm:"default:null" json:"doctor_id"`
	Doctor               Doctor
	OfficerID            uint `gorm:"default:null" json:"officer_id"`
	Officer              Officer
	DriverID             uint `gorm:"default:null" json:"driver_id"`
	Driver               Driver
	VisitedAt            time.Time `gorm:"default:null" json:"visited_at"`
	FinishedAt           time.Time `gorm:"default:null" json:"finished_at"`
	IsVisited            bool      `gorm:"type:bool;default:false" json:"is_visited"`
	IsFinished           bool      `gorm:"type:bool;default:false" json:"is_finished"`
	gorm.Model           `json:"-"`
}
