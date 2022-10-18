package manage

import "time"

type CreateUserRequest struct {
	Msisdn   string    `json:"msisdn"`
	Identity string    `json:"identity"`
	Name     string    `json:"name"`
	Dob      time.Time `json:"dob"`
	Gender   bool      `json:"gender"`
	Address  string    `json:"address"`
	IsBpjs   bool      `json:"is_bpjs"`
}

type AuthRequest struct {
	Msisdn string `json:"msisdn"`
}

type VerifyRequest struct {
	Otp string `json:"otp"`
}

type Interface interface {
}
