package model

type ChatDisease struct {
	ChatID    uint64  `json:"chat_id"`
	Chat      Chat    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DiseaseID uint    `json:"disease_id"`
	Disease   Disease `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
