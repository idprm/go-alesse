package model

import "github.com/idprm/go-alesse/src/pkg/common"

type Referral struct {
	common.Model
	DoctorID           uint   `json:"-"`
	Doctor             Doctor `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DoctorSpecialistID uint   `json:"-"`
	DoctorSpecialist   Doctor `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ChannelName        string `gorm:"size:200" json:"channel_name"`
	ChannelUrl         string `gorm:"size:200" json:"channel_url"`
	ShortLink          string `gorm:"size:50" json:"short_link"`
}
