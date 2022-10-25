package model

type Healthcenter struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Name    string `gorm:"size:150;not null" json:"name"`
	Photo   string `gorm:"size:150;not null" json:"photo"`
	Address string `gorm:"type:text" json:"address"`
	Phone   string `gorm:"size:20" json:"phone"`
}
