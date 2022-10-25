package model

import "github.com/idprm/go-alesse/src/pkg/common"

type Prescription struct {
	common.Model
	ChatID uint64 `json:"chat_id"`
	Chat   Chat
}
