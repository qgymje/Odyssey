package models

import (
	"database/sql"
	"time"
)

// User model 表示一个用户
type User struct {
	ID        int64
	Phone     string
	Email     sql.NullString // 通过email向register发送用户统计数据
	Nickname  sql.NullString
	Password  string
	Salt      string
	Avatar    sql.NullString
	Sex       NullUint8
	Height    NullUint8
	Weight    NullUint8
	Birthday  NullTime
	Latitude  sql.NullFloat64
	Longitude sql.NullFloat64
	Token     sql.NullString
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt NullTime  `db:"deleted_at"`
}

// Fetch 从db里获取数据, 通常用于已经有了id
func (u *User) Fetch() (err error) {
	err = GetDB().Get(u, "select * from users where id=?", u.ID)
	return
}

// Create 用于创建一个用户
func (u *User) Create() (err error) {
	now := time.Now()
	u.CreatedAt = now
	u.UpdatedAt = now
	//result, err := GetDB().Exec(`insert into users(phone, email, nickname, password, salt, avatar, sex, height, weight, birthday, latitude, longitude, token, created_at, updated_at, deleted_at) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, u.Phone, u.Email, u.Nickname, u.Password, u.Salt, u.Avatar.String, u.Sex, u.Height, u.Weight, u.Birthday, u.Latitude, u.Longitude, u.Token, u.CreatedAt, u.UpdatedAt, u.DeletedAt)
	result, err := GetDB().NamedExec(`insert into users(phone, email, nickname, password, salt, avatar, sex, height, weight, birthday, latitude, longitude, token, created_at, updated_at, deleted_at) values(:phone, :email, :nickname, :password, :salt, :avatar, :sex, :height, :weight, :birthday, :latitude, :longitude, :token, :created_at, :updated_at, :deleted_at)`, u)

	if err != nil {
		return
	}
	if _, err = result.RowsAffected(); err != nil {
		return
	}
	u.ID, err = result.LastInsertId()
	return
}

// Delete 表示注销一个用户
func (u *User) Delete() (err error) {
	u.DeletedAt = NullTime{Time: time.Now()}
	result := GetDB().MustExec(`update users set deleted_at = ? where id = ?`, u.DeletedAt, u.ID)
	_, err = result.RowsAffected()
	return
}

// IsDeleted 判断用户是否已经注销
func (u *User) IsDeleted() bool {
	return u.DeletedAt.Time.IsZero()
}

// RemoveToken 清除token
func (u *User) RemoveToken() (err error) {
	return u.UpdateToken("")
}

// UpdateToken 更新token
func (u *User) UpdateToken(token string) (err error) {
	result := GetDB().MustExec(`update users set token=? where id=?`, token, u.ID)
	if _, err = result.RowsAffected(); err != nil {
		return
	}
	u.Token = sql.NullString{String: token}
	return
}

// FindUserByToken 根据token查找用户
func FindUserByToken(token string) (*User, error) {
	var user User
	var err error
	err = GetDB().Get(&user, `select * from users where token=?`, token)
	return &user, err
}

// FindUserByPhone 根据phone查找用户
func FindUserByPhone(phone string) (*User, error) {
	var user User
	var err error
	err = GetDB().Get(&user, `select * from users where phone=?`, phone)
	return &user, err
}

// IsPhoneRegisted 检查手机号是否已经存在
func IsPhoneRegisted(phone string) bool {
	result := GetDB().MustExec(`select count(*) from users where phone=?`, phone)
	if cnt, err := result.RowsAffected(); err != nil {
		return cnt > 0
	}
	return false
}
