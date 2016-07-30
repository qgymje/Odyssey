package models

import "time"

// SMSCode model 表示一次生成短信验证码纪录
type SMSCode struct {
	ID     int64
	Phone  string
	Code   string
	UsedAt NullTime `db:"used_at"`

	CreatedAt time.Time `db:"created_at"`
}

// Create 生成一条db纪录
func (s *SMSCode) Create() (err error) {
	s.CreatedAt = time.Now()

	result := GetDB().MustExec(`insert into sms_codes(phone, code, created_at) values(?,?,?)`, s.Phone, s.Code, s.CreatedAt)
	if _, err = result.RowsAffected(); err != nil {
		return
	}
	s.ID, err = result.LastInsertId()

	return
}

// IsUsed 判断一个code是否已经被使用过了
func (s *SMSCode) IsUsed() bool {
	return !s.UsedAt.Time.IsZero()
}

// UseCode 当注册成功之后将used_at更新
func (s *SMSCode) UseCode() (err error) {
	result := GetDB().MustExec(`update sms_codes set used_at = ? where id = ?`, s.UsedAt, s.ID)
	if _, err = result.RowsAffected(); err != nil {
		return
	}
	s.UsedAt = NullTime{Time: time.Now()}
	return
}

func (s *SMSCode) GeneratedInDuration(duration time.Duration) bool {
	return time.Since(s.CreatedAt) < duration
}

// FindSMSCode 根据手机号查找一条验证码信息
func FindSMSCodeByPhone(phone string) (*SMSCode, error) {
	var sms SMSCode
	var err error
	err = db.Get(&sms, `select * from sms_codes where phone=? order by id desc limit 1`, phone)
	return &sms, err
}
