package models

import (
	"Odyssey/utils"
	"time"
)

// SMSCode model 表示一次生成短信验证码纪录
type SMSCode struct {
	TableName struct{} `sql:"smscodes"`
	ID        int      `json:"smscode_id"`
	Phone     string   `json:"phone"`
	Code      string   `json:"code"`

	UsedAt    time.Time `sql:",null" json:"used_at"`
	CreatedAt time.Time `json:"created_at"`
}

// Create 生成一条db纪录
func (s *SMSCode) Create() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("models.smscode.Create error: ", err)
		}
	}()

	s.CreatedAt = time.Now()
	err = GetDB().Create(s)

	return
}

// IsUsed 判断一个code是否已经被使用过了
func (s *SMSCode) IsUsed() bool {
	return !s.UsedAt.IsZero()
}

// Update 更新一条验证码纪录
func (s *SMSCode) Update(where map[string]interface{}, update map[string]interface{}) (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("models.smscode.Update error: ", err)
		}
	}()

	query := GetDB().Model(s)
	for key, val := range update {
		query = query.Set(key, val)
	}
	for key, val := range where {
		query = query.Where(key, val)
	}
	// 判断第一个返回值
	_, err = query.Update()

	return
}

// FindSMSCode 根据手机号查找一条验证码信息
func FindSMSCode(phone string) (sms SMSCode, err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("models.smscode.FindSMSCode error: ", err)
		}
	}()

	err = GetDB().Model(&sms).Where("phone=?", phone).Order("id DESC").Limit(1).Select()

	return
}
