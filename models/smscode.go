package models

import "time"

const (
	DB_SMSCODE_PHONE = "phone"
	DB_SMSCODE_CODE  = "code"
)

type SMSCode struct {
	Id    uint64
	Phone string
	Code  string

	UsedAt    time.Time
	CreatedAt time.Time

	Base
}

func (s *SMSCode) TableName() string {
	return "smscodes"
}

func (s *SMSCode) Create() error {
	return nil
}
