package model

import "gorm.io/gorm"

type Sendbird struct {
	ID         uint64 `gorm:"primaryKey" json:"id"`
	Msisdn     string `gorm:"size:15;not null" json:"msisdn"`
	Action     string `gorm:"size:120" json:"action"`
	Response   string `gorm:"type:text" json:"response"`
	gorm.Model `json:"-"`
}
