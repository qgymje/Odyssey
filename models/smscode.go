package models

import (
	"time"

	"github.com/diegogub/aranGO"
)

// SMSCode model 表示一次生成短信验证码纪录
type SMSCode struct {
	aranGO.Document
	ID        int64
	Phone     string
	Code      string
	UsedAt    time.Time `required:"-"`
	CreatedAt time.Time
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
	result := GetDB().MustExec(`update sms_codes set used_at = ? where id = ?`, s.UsedAt, s.ID)
	if _, err = result.RowsAffected(); err != nil {
		return

	}
	s.UsedAt = time.Now()
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
	err = GetDB().Get(&sms, `select * from sms_codes where phone=? order by id desc limit 1`, phone)
	return &sms, err
}
