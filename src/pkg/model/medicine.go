package model

type Medicine struct {
	ID       uint64 `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"size:200;not null" json:"name"`
	Category string `gorm:"size:20" json:"category"`
	Type     string `gorm:"size:25" json:"type"`
	Unit     string `gorm:"size:25" json:"unit"`
	IsActive bool   `gorm:"type:bool;default:true" json:"is_active"`
}
