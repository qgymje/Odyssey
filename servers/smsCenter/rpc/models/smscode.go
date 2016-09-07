package models

import (
	"time"

	"github.com/qgymje/Odyssey/commons/utils"
)

// SMS model 表示一次生成短信验证码纪录
type SMS struct {
	ID        int64
	Phone     string
	Type      int32 `db:"sms_type"`
	Code      string
	UsedAt    utils.NullTime `db:"used_at"`
	CreatedAt time.Time      `db:"created_at"`
}

// Create 生成一条db纪录
func (s *SMS) Create() (err error) {
	s.CreatedAt = time.Now()
	result := GetDB().MustExec(`insert into sms_codes(phone, sms_type, code, created_at) values(?,?,?)`, s.Phone, s.Type, s.Code, s.CreatedAt)
	if _, err = result.RowsAffected(); err != nil {
		return
	}
	s.ID, err = result.LastInsertId()

	return
}

// IsUsed 判断一个code是否已经被使用过了
func (s *SMS) IsUsed() bool {
	return !s.UsedAt.Time.IsZero()
}

// UseCode 当注册成功之后将used_at更新
func (s *SMS) UseCode() (err error) {
	result := GetDB().MustExec(`update sms_codes set used_at = ? where id = ?`, s.UsedAt, s.ID)
	if _, err = result.RowsAffected(); err != nil {
		return
	}
	s.UsedAt = utils.NullTime{Time: time.Now()}
	return
}

// IsGeneratedInDuration 判断验证码是否过期
func (s *SMS) IsGeneratedInDuration(duration time.Duration) bool {
	return time.Since(s.CreatedAt) < duration
}

// FindSMSCodeByPhone 根据手机号查找一条验证码信息
func FindSMSCodeByPhone(phone string) (*SMS, error) {
	var sms SMS
	var err error
	err = db.Get(&sms, `select * from sms_codes where phone=? order by id desc limit 1`, phone)
	return &sms, err
}
