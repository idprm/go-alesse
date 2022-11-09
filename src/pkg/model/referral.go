package model

import "gorm.io/gorm"

type Referral struct {
	ID                 uint64 `gorm:"primaryKey" json:"id"`
	DoctorID           uint   `json:"-"`
	Doctor             Doctor `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DoctorSpecialistID uint   `json:"-"`
	DoctorSpecialist   Doctor `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ChannelName        string `gorm:"size:200" json:"channel_name"`
	ChannelUrl         string `gorm:"size:200" json:"channel_url"`
	gorm.Model         `json:"-"`
}
