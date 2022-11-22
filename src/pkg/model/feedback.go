package model

import (
	"time"

	"gorm.io/gorm"
)

type Feedback struct {
	ID         uint64 `gorm:"primaryKey" json:"id"`
	ChatID     uint64 `json:"chat_id"`
	Chat       Chat
	Slug       string    `gorm:"size:50;unique" json:"slug"`
	Rating     float32   `gorm:"size:3" json:"rating"`
	Suggestion string    `gorm:"type:text" json:"suggestion"`
	IsSubmited bool      `gorm:"type:bool;default:false" json:"is_submited"`
	SubmitedAt time.Time `gorm:"default:null" json:"finished_at"`
	gorm.Model `json:"-"`
}
