package model

type PrescriptionMedicine struct {
	PrescriptionID uint64 `json:"prescription_id"`
	Prescription   Prescription
	MedicineID     uint64 `json:"medicine_id"`
	Medicine       Medicine
	Name           string `gorm:"size:55" json:"name"`
	Quantity       uint   `gorm:"size:8" json:"quantity"`
	Dose           string `gorm:"size:15" json:"dose"`
	Rule           string `gorm:"size:55" json:"rule"`
	Description    string `gorm:"type:text" json:"description"`
}
