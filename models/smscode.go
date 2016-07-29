package models

import (
	"fmt"
	"time"
)

// SMSCode model 表示一次生成短信验证码纪录
type SMSCode struct {
	ID     int      `json:"smscode_id"`
	Phone  string   `gorm:"not null;type:varchar(11);index:idx_smscode_phone" json:"phone"`
	Code   string   `gorm:"not null;type:varchar(6)" json:"code"`
	UsedAt NullTime `json:"used_at"`

	CreatedAt time.Time `json:"created_at"`
}

// Create 生成一条db纪录
func (s *SMSCode) Create() (err error) {
	s.CreatedAt = time.Now()
	GetDB().Create(s)

	return
}

// IsUsed 判断一个code是否已经被使用过了
func (s *SMSCode) IsUsed() bool {
	return !s.UsedAt.Time.IsZero()
}

// UseCode 当注册成功之后将used_at更新
func (s *SMSCode) UseCode() (err error) {
	GetDB().Model(s).Update("used_at", time.Now())
	return
}

// Update 更新一条验证码纪录
func (s *SMSCode) Update(where map[string]interface{}, update map[string]interface{}) (err error) {
	query := GetDB().Model(s)
	for key, val := range where {
		query = query.Where(key, val)
	}
	query.Updates(update)

	return
}

// FindSMSCode 根据手机号查找一条验证码信息
func FindSMSCode(phone string) (*SMSCode, error) {
	var err error
	var sms SMSCode
	GetDB().Where("phone=?", phone).Order("id DESC").Limit(1).First(&sms)
	if sms.Code == "" {
		return nil, fmt.Errorf("smscode not exists with phone: %s", phone)
	}
	return &sms, err
}
