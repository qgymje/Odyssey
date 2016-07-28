package models

import (
	"Odyssey/utils"
	"time"
)

// User model 表示一个用户
type User struct {
	TableName struct{} `sql:"users"`
	ID        int      `json:"user_id"`
	Phone     string   `json:"phone"`
	Email     string   `json:"email"` // 通过email向register发送用户统计数据
	Nickname  string   `json:"nickname"`
	Password  string   `json:"-"`
	Salt      string   `json:"-"`
	Avatar    string   `json:"avatar"`

	Sex      uint8     `json:"sex"`
	Height   float64   `json:"height"`
	Weight   float64   `json:"weight"`
	Birthday time.Time `json:"birthday"`

	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`

	Token string `json:"token"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"-"`
}

// IsPhoneExists 检查手机号是否已经存在
func IsPhoneExists(phone string) bool {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("models.IsPhoneExists error: %v", err)
		}
	}()

	cnt, err := GetDB().Model(&User{}).Where("phone=?", phone).Count()
	return cnt > 0
}

// Create 用于创建一个用户
func (u *User) Create() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("models.user.Create error: %v", err)
		}
	}()

	now := time.Now()
	u.CreatedAt = now
	u.UpdatedAt = now
	err = GetDB().Create(u)

	return
}

// Update 更新一个用户的信息
func (u *User) Update(where map[string]interface{}, update map[string]interface{}) (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("models.user.Update error: %v", err)
		}
	}()

	update["created_at=?"] = time.Now()

	query := GetDB().Model(u)
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

// Delete 表示注销一个用户
func (u *User) Delete(where map[string]interface{}) (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("models.user.Delete error: %v", err)
		}
	}()

	update := map[string]interface{}{
		"deleted_at=?": time.Now(),
	}

	if err := u.Update(where, update); err != nil {
		return err
	}

	return nil
}

// IsDeleted 判断用户是否已经注销
func (u *User) IsDeleted() bool {
	return u.DeletedAt.IsZero()
}

// FindUsers 根据条件查找用户
// where map[string]interface{}
// key与val必须为sql语法, 比如where["id=?] = 1
func FindUsers(where map[string]interface{}, order string, limit int, offset int) (users []*User, err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("models.user.FindUsers error: %v", err)
		}
	}()

	query := GetDB().Model(&users)
	for key, val := range where {
		query = query.Where(key, val)
	}
	err = query.Order(order).Limit(limit).Offset(offset).Select()

	return
}
