package model

import "github.com/idprm/go-alesse/src/common"

type Sendbird struct {
	common.Model
	Msisdn  string `gorm:"msisdn" json:"msisdn"`
	Action  string `gorm:"action" json:"action"`
	Payload string `gorm:"type:text" json:"payload"`
}
