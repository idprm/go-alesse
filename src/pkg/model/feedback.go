package model

import "gorm.io/gorm"

type Feedback struct {
	ID         uint64 `gorm:"primaryKey" json:"id"`
	ChatID     uint64 `json:"chat_id"`
	Chat       Chat
	Rating     float32 `gorm:"size:3" json:"rating"`
	Suggestion string  `gorm:"type:text" json:"suggestion"`
	IsSubmited bool    `gorm:"type:bool;default:false" json:"is_submited"`
	gorm.Model `json:"-"`
}
