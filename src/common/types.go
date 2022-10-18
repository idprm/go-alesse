package common

import "time"

// Model is base for database struct
type Model struct {
	ID        uint64     `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at" gorm:"not null;"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"not null;"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
}
