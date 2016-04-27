package models

import (
	"Odyssey/utils"
	"fmt"
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

	Token string

	UpdatedAt time.Time
	CreatedAt time.Time
	DeletedAt time.Time

	Base
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) Create() error {
	u.CreatedAt = time.Now()
	query := sq.Insert(u.TableName()).
		Columns("phone", "nickname", "password", "salt", "height", "weight", "token", "updated_at", "created_at", "deleted_at").
		Values(u.Phone, u.Nickname, u.Password, u.Salt, u.Height, u.Weight, u.Token, u.UpdatedAt, u.CreatedAt, u.DeletedAt).
		Suffix("RETURNING \"id\"").
		RunWith(GetDB()).
		PlaceholderFormat(sq.Dollar)

	// 注意这里必须要传指针
	query.QueryRow().Scan(&u.Id)
	return nil
}

func (u *User) Update(where map[string]interface{}, update map[string]interface{}) error {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("models.user.Update error: %v", err)
		}
	}()

	u.UpdatedAt = time.Now()
	query := sq.Update(u.TableName()).
		SetMap(sq.Eq(update))

	for k, v := range where {
		fmt.Printf("k = %v, v = %v \n", k, v)
		query.Where(k, v)
	}

	sql, _, err := query.ToSql()
	if err != nil {
		return err
	} else {
		utils.GetLog().Debug("models.user.Update sql = : %s", sql)
	}

	result, err := query.RunWith(GetDB()).Exec()

	if err != nil {
		return err
	}
	if n, err := result.RowsAffected(); n == 0 && err != nil {
		return err
	}

	return nil
}

func (u *User) Delete(where map[string]interface{}) error {
	u.DeletedAt = time.Now()

	return nil
}

func (u *User) IsDeleted() bool {
	return u.DeletedAt.IsZero()
}

func FindUsers(where map[string]interface{}) ([]*User, error) {
	return nil, nil
}

func FindUser(where map[string]interface{}) (*User, error) {
	return nil, nil
}
