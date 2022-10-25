package model

type Healthcenter struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Code    string `gorm:"size:20" json:"code"`
	Name    string `gorm:"size:150;not null" json:"name"`
	Type    string `gorm:"size:150" json:"type"`
	Photo   string `gorm:"size:150;not null" json:"photo"`
	Address string `gorm:"type:text" json:"address"`
}
