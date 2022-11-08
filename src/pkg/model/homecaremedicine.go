package model

type HomecareMedicine struct {
	HomecareID  uint64 `json:"homecare_id"`
	Homecare    Homecare
	MedicineID  uint64 `json:"medicine_id"`
	Medicine    Medicine
	Name        string `gorm:"size:55" json:"name"`
	Quantity    uint   `gorm:"size:8" json:"quantity"`
	Rule        string `gorm:"size:55" json:"rule"`
	Time        string `gorm:"size:15" json:"time"`
	Description string `gorm:"type:text" json:"description"`
}
