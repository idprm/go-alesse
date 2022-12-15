package model

type Category struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Code     string `gorm:"size:50;not null" json:"code"`
	Name     string `gorm:"size:100;not null" json:"name"`
	IsActive bool   `gorm:"type:bool;default:true" json:"is_active"`
}
