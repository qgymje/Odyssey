package models

import (
	"time"

	sq "github.com/lann/squirrel"
	_ "github.com/lib/pq"
)

type User struct {
	Id       uint64
	Phone    string
	Nickname string
	Password string
	Salt     string
	Height   float32 // 可能会有小数点
	Weight   float32

	UpdatedAt time.Time
	CreatedAt time.Time
	DeletedAt time.Time

	Base
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) Create() error {
	query := sq.Insert(u.TableName()).
		Columns("phone", "nickname", "password", "salt", "height", "weight", "updated_at", "created_at", "deleted_at").
		Values(u.Phone, u.Nickname, u.Password, u.Salt, u.Height, u.Weight, u.UpdatedAt, u.CreatedAt, u.DeletedAt).
		Suffix("RETURNING \"id\"").
		RunWith(GetDB()).
		PlaceholderFormat(sq.Dollar)
	query.QueryRow().Scan(u.Id)
	return nil
}

func (u *User) Update() error {

	return nil
}

func (u *User) Delete() error {
	return nil
}

func FindUser(filters map[string]interface{}) ([]*User, error) {
	return nil, nil
}

func FindUsers(filters map[string]interface{}) (*User, error) {
	return nil, nil
}
