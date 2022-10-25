package model

import "github.com/idprm/go-alesse/src/pkg/common"

type Verify struct {
	common.Model
	Msisdn   string `gorm:"size:15;not null" json:"msisdn"`
	Otp      string `gorm:"size:5" json:"otp"`
	IsVerify bool   `gorm:"type:bool" json:"is_verify"`
}
