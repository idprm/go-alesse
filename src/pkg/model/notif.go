package model

import "gorm.io/gorm"

type Notif struct {
	ID         uint64 `gorm:"primaryKey" json:"id"`
	UserID     uint64 `json:"user_id"`
	User       User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Content    string `gorm:"type:text" json:"content"`
	gorm.Model `json:"-"`
}
