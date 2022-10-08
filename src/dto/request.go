package dto

import "time"

type ReqRegister struct {
	Msisdn   string    `json:"msisdn"`
	Identity string    `json:"identity"`
	Name     string    `json:"name"`
	Dob      time.Time `json:"dob"`
	Gender   bool      `json:"gender"`
	Address  string    `json:"address"`
	IsBpjs   bool      `json:"is_bpjs"`
}

type ReqLogin struct {
	Msisdn string `json:"msisdn"`
}

type ReqVerify struct {
	Otp string `json:"otp"`
}
