package model

type Disease struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"string" json:"name"`
	IsActive bool   `gorm:"type:bool" json:"is_active"`
}
