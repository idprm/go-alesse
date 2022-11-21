package model

type PharmacyMedicine struct {
	PharmacyID  uint64 `json:"pharmacy_id"`
	Pharmacy    Pharmacy
	MedicineID  uint64 `json:"medicine_id"`
	Medicine    Medicine
	Name        string `gorm:"size:55" json:"name"`
	Quantity    uint   `gorm:"size:8" json:"quantity"`
	Rule        string `gorm:"size:55" json:"rule"`
	Dose        string `gorm:"size:15" json:"dose"`
	Description string `gorm:"type:text" json:"description"`
}
