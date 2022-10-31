package model

type Medicine struct {
	ID       uint64 `gorm:"primaryKey" json:"id"`
	Code     string `gorm:"size:20" json:"code"`
	Name     string `gorm:"size:150;not null" json:"name"`
	IsActive bool   `gorm:"type:bool" json:"is_active"`
}
