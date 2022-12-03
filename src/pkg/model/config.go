package model

type Config struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"string" json:"name"`
	Value string `gorm:"string" json:"value"`
}
