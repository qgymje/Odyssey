package models

import "time"

const (
	DB_SMSCODE_PHONE = "phone"
	DB_SMSCODE_CODE  = "code"
)

type SMSCode struct {
	id    uint64
	Phone string
	Code  string

	UsedAt    time.Time
	CreatedAt time.Time
}
