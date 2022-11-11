package model

import "time"

type HomecareOfficer struct {
	ID                 uint64 `gorm:"primaryKey" json:"id"`
	HomecareID         uint64 `json:"homecare_id"`
	Homecare           Homecare
	Slug               string `gorm:"size:50;unique" json:"slug"`
	Treatment          string `gorm:"type:text" json:"treatment"`
	FinalDiagnosis     string `gorm:"type:text" json:"final_diagnosis"`
	DrugAdministration string `gorm:"type:text" json:"drug_administration"`
	DoctorID           uint   `gorm:"default:null" json:"doctor_id"`
	Doctor             Doctor
	OfficerID          uint `gorm:"default:null" json:"officer_id"`
	Officer            Officer
	DriverID           uint `gorm:"default:null" json:"driver_id"`
	Driver             Driver
	VisitedAt          time.Time `gorm:"default:null" json:"visited_at"`
	IsVisited          bool      `gorm:"type:bool;default:false" json:"is_visited"`
}
