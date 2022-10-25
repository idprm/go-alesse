package model

import (
	"encoding/json"
	"time"

	"github.com/idprm/go-alesse/src/pkg/common"
)

type User struct {
	common.Model
	Msisdn   string    `gorm:"size:15;unique;not null" json:"msisdn"`
	Name     string    `gorm:"size:200;not null" json:"name"`
	Identity string    `gorm:"size:20" json:"identity"`
	Dob      time.Time `gorm:"default:null" json:"dob"`
	Gender   string    `gorm:"size:15" json:"gender"`
	Address  string    `gorm:"type:text" json:"address"`
	IsVerify bool      `gorm:"type:bool" json:"is_verify"`
	IsBpjs   bool      `gorm:"type:bool" json:"is_bpjs"`
}

func New() *User {
	return &User{}
}

// MarshalJSON ...
func (s *User) MarshalJSON() ([]byte, error) {
	type Alias User

	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(s),
	})
}
