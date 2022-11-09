package model

import "gorm.io/gorm"

type Chat struct {
	ID          uint64 `gorm:"primaryKey" json:"id"`
	OrderID     uint64 `json:"order_id"`
	Order       Order  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DoctorID    uint   `json:"-"`
	Doctor      Doctor `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID      uint64 `json:"user_id"`
	User        User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ChannelName string `gorm:"size:200;unique" json:"channel_name"`
	ChannelUrl  string `gorm:"size:200" json:"channel_url"`
	gorm.Model  `json:"-"`
}
