package model

import "gorm.io/gorm"

type PrescriptionMedicine struct {
	PrescriptionID uint64 `json:"prescription_id"`
	Prescription   Prescription
	MedicineID     uint64 `json:"medicine_id"`
	Medicine       Medicine
	gorm.Model     `json:"-"`
}
