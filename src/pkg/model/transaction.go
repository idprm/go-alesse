package model

import "gorm.io/gorm"

type Transaction struct {
	ID           uint64 `gorm:"primaryKey" json:"id"`
	UserID       uint64 `json:"-"`
	User         User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ChatID       uint64 `json:"chat_id"`
	Chat         Chat   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SystemStatus string `gorm:"type:text" json:"system_status"`
	UserStatus   string `gorm:"type:text" json:"user_status"`
	NotifStatus  string `gorm:"type:text" json:"notif_status"`
	gorm.Model   `json:"-"`
}
