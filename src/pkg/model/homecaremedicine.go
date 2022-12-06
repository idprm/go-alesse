package model

type HomecareMedicine struct {
	HomecareID  uint64 `json:"homecare_id"`
	Homecare    Homecare
	MedicineID  uint64 `json:"medicine_id"`
	Medicine    Medicine
	Name        string `gorm:"size:250" json:"name"`
	Quantity    uint   `gorm:"size:0" json:"quantity"`
	Rule        string `gorm:"size:250" json:"rule"`
	Dose        string `gorm:"size:250" json:"dose"`
	Description string `gorm:"type:text" json:"description"`
}
