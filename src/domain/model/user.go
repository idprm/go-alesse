package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         uint64    `gorm:"primaryKey" json:"id"`
	Msisdn     string    `gorm:"size:15;unique;not null" json:"msisdn"`
	Name       string    `gorm:"size:200;not null" json:"name"`
	Identity   string    `gorm:"size:20;unique" json:"identity"`
	Dob        time.Time `gorm:"default:null" json:"dob"`
	Gender     string    `gorm:"size:15" json:"gender"`
	Address    string    `gorm:"type:text" json:"address"`
	IsVerify   bool      `gorm:"type:bool" json:"is_verify"`
	IsBpjs     bool      `gorm:"type:bool" json:"is_bpjs"`
	gorm.Model `json:"-"`
}
