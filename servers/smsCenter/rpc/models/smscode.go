package models

import (
	"errors"
	"time"

	"github.com/diegogub/aranGO"
)

// SMSCode model 表示一次生成短信验证码纪录
type SMSCode struct {
	aranGO.Document
	Phone     string    `json:"phone"`
	Code      string    `json:"code"`
	UsedAt    time.Time `json:"used_at" required:"-"`
	CreatedAt time.Time `json:"created_at"`
}

/*
func (s *SMSCode) GetKey() string {
	return s.Key
}

func (s *SMSCode) GetCollection() string {
	return "smscodes"
}

func (s *SMSCode) GetError() (string, bool) {
	return s.Message, s.Error
}
*/

// Create 生成一条db纪录
func (s *SMSCode) Create() (err error) {
	s.CreatedAt = time.Now()
	err = GetSession().DB(DB_NAME).Col(DOC_SMSCodes).Save(&s)
	return
}

// IsUsed 判断一个code是否已经被使用过了
func (s *SMSCode) IsUsed() bool {
	return !s.UsedAt.IsZero()
}

// UseCode 当注册成功之后将used_at更新
func (s *SMSCode) UseCode() (err error) {
	s.UsedAt = time.Now()
	err = GetSession().DB(DB_NAME).Col(DOC_SMSCodes).Replace(s.Key, s)
	return
}

// IsGeneratedInDuration 判断是否在duration时间内再次请求了?
func (s *SMSCode) IsGeneratedInDuration(duration time.Duration) bool {
	return time.Since(s.CreatedAt) < duration
}

// FindSMSCodeByPhone 根据手机号查找一条验证码信息
func FindSMSCodeByPhone(phone string) (*SMSCode, error) {
	var sms SMSCode
	var err error
	query := aranGO.NewQuery(`for s in smscodes filter s.phone== "` + phone + `" sort s.created_at desc limit 1   return s`)
	course, err := GetSession().DB(DB_NAME).Execute(query)
	if err != nil {
		return nil, err
	}
	if course.FetchOne(&sms) {
		return &sms, nil
	} else {
		return nil, errors.New("no row found")
	}
}
