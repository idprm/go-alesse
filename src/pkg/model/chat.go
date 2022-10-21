package model

import "github.com/idprm/go-alesse/src/common"

type Chat struct {
	common.Model
	OrderID     uint64 `json:"order_id"`
	Order       Order  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DoctorID    uint   `json:"-"`
	Doctor      Doctor `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID      uint64 `json:"user_id"`
	User        User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ChannelName string `gorm:"size:200" json:"channel_name"`
	ChannelUrl  string `gorm:"size:200" json:"channel_url"`
	ShortLink   string `gorm:"size:50" json:"short_link"`
}
