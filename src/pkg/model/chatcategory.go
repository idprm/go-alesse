package model

type ChatCategory struct {
	ChatID     uint64   `json:"chat_id"`
	Chat       Chat     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CategoryID uint     `json:"category_id"`
	Category   Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
