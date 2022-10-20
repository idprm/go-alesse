package model

import "github.com/idprm/go-alesse/src/common"

type Medical struct {
	common.Model
	UserID   uint64 `json:"-"`
	User     User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DoctorID uint   `json:"-"`
	Doctor   Doctor `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
