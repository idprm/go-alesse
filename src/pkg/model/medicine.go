package model

type Medicine struct {
	ID       uint64 `gorm:"primaryKey" json:"id"`
	Code     string `gorm:"size:20" json:"code"`
	Type     string `gorm:"size:25" json:"type"`
	Name     string `gorm:"size:150;not null" json:"name"`
	Unit     string `gorm:"size:25" json:"unit"`
	IsActive bool   `gorm:"type:bool" json:"is_active"`
}
