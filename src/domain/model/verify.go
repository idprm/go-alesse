package model

import "gorm.io/gorm"

type Verify struct {
	ID         uint64 `gorm:"primaryKey" json:"id"`
	Msisdn     string `gorm:"size:15;not null" json:"msisdn"`
	Otp        string `gorm:"size:5" json:"otp"`
	IsVerify   bool   `gorm:"type:bool" json:"is_verify"`
	gorm.Model `json:"-"`
}
