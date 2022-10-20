package model

import "github.com/idprm/go-alesse/src/common"

type Zenziva struct {
	common.Model
	Msisdn   string `gorm:"size:15;not null" json:"msisdn"`
	Action   string `gorm:"size:120" json:"action"`
	Response string `gorm:"type:text"`
}
