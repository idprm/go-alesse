package model

import "github.com/idprm/go-alesse/src/pkg/common"

type Order struct {
	common.Model
	UserID   uint64 `json:"-"`
	User     User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DoctorID uint   `json:"-"`
	Doctor   Doctor `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Number   string `json:"number"`
	Total    int    `json:"total"`
	Status   string `gorm:"size:25" json:"status"`
}
