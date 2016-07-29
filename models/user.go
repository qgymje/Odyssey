package models

import (
	"Odyssey/utils"
	"database/sql"
	"errors"
	"time"
)

// User model 表示一个用户
type User struct {
	ID        int
	Phone     string         `gorm:"index:idx_user_phone;type:varchar(11)"`
	Email     sql.NullString `gorm:"type:varchar(64)"` // 通过email向register发送用户统计数据
	Nickname  string         `gorm:"type:varchar(16)"`
	Password  string         `gorm:"type:char(32)"`
	Salt      string         `gorm:"type:char(6)"`
	Avatar    NullString
	Sex       NullUint8
	Height    NullUint8
	Weight    NullUint8
	Birthday  NullTime
	Latitude  NullFloat64
	Longitude NullFloat64
	Token     NullString `gorm:"index:idx_user_token"`
	CreatedAt time.Time  `gorm:"index:idx_user_created_at"`
	UpdatedAt time.Time
	DeletedAt NullTime
}

// IsPhoneExists 检查手机号是否已经存在
func IsPhoneExists(phone string) bool {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("models.IsPhoneExists error: ", err)
		}
	}()

	var cnt int
	GetDB().Model(&User{}).Where("phone=?", phone).Count(&cnt)
	return cnt > 0
}

// Create 用于创建一个用户
func (u *User) Create() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("models.user.Create error: ", err)
		}
	}()

	now := time.Now()
	u.CreatedAt = now
	u.UpdatedAt = now
	GetDB().Create(u)

	return
}

// Update 更新一个用户的信息
func (u *User) Update(where map[string]interface{}, update map[string]interface{}) (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("models.user.Update error: ", err)
		}
	}()

	update["created_at=?"] = time.Now()

	db := GetDB().Model(u)
	for key, val := range where {
		db = db.Where(key, val)
	}
	// 判断第一个返回值
	db.Updates(update)

	return
}

// Delete 表示注销一个用户
func (u *User) Delete(where map[string]interface{}) (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("models.user.Delete error: ", err)
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
	return u.DeletedAt.Time.IsZero()
}

// FindUsers 根据条件查找用户
// where map[string]interface{}
// key与val必须为sql语法, 比如where["id=?] = 1
func FindUsers(where map[string]interface{}, order string, limit int, offset int) (users []*User, err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("models.user.FindUsers error: ", err)
		}
	}()

	query := GetDB()
	for key, val := range where {
		query = query.Where(key, val)
	}

	query.Order(order).Limit(limit).Offset(offset).Find(&users)

	return
}

// FindUser 根据条件查找一个用户
func FindUser(where map[string]interface{}) (user *User, err error) {
	users, err := FindUsers(where, "id ASC", 1, 0)
	if err != nil {
		return nil, err
	}
	if len(users) == 1 {
		return users[0], nil
	} else {
		return nil, errors.New("models.FindUser conditions unfit")
	}
}
