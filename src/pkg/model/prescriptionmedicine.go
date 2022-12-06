package model

type PrescriptionMedicine struct {
	PrescriptionID uint64 `json:"prescription_id"`
	Prescription   Prescription
	MedicineID     uint64 `json:"medicine_id"`
	Medicine       Medicine
	Name           string `gorm:"size:250" json:"name"`
	Quantity       uint   `gorm:"size:9" json:"quantity"`
	Dose           string `gorm:"size:250" json:"dose"`
	Rule           string `gorm:"size:250" json:"rule"`
	Description    string `gorm:"type:text" json:"description"`
}
