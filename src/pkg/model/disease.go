package model

type Disease struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Code     string `gorm:"size:200" json:"code"`
	Name     string `gorm:"type:text" json:"name"`
	IsActive bool   `gorm:"type:bool" json:"is_active"`
}
