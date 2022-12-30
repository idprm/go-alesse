package model

type History struct {
	ID        uint64 `gorm:"primaryKey" json:"id"`
	ChatID    uint64 `json:"chat_id"`
	Chat      Chat
	DiseaseID uint `json:"disease_id"`
	Disease   Disease
}
