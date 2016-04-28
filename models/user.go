package models

import (
	"Odyssey/utils"
	"time"

	sq "github.com/lann/squirrel"
)

type User struct {
	Id       uint64  `json:"id"`
	Phone    string  `json:"phone"`
	Nickname string  `json:"nickname"`
	Password string  `json:"-"`
	Salt     string  `json:"-"`
	Height   float64 `json:"height"`
	Weight   float64 `json:"weight"`

	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`

	Token string `json:"token"`

	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"-"`

	Base
}

func NewUser() *User {
	u := new(User)
	return u
}

func (User) TableName() string {
	return "users"
}

func (u *User) Create() error {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("models.user.Create error: %v", err)
		}
	}()

	u.CreatedAt = time.Now()
	query := sq.Insert(u.TableName()).
		Columns("phone", "nickname", "password", "salt", "height", "weight", "latitude", "longitude", "token", "updated_at", "created_at", "deleted_at").
		Values(u.Phone, u.Nickname, u.Password, u.Salt, u.Height, u.Weight, u.Latitude, u.Longitude, u.Token, u.UpdatedAt, u.CreatedAt, u.DeletedAt).
		Suffix("RETURNING \"id\"").
		RunWith(GetDB()).
		PlaceholderFormat(sq.Dollar)

	// 注意这里必须要传指针
	if err = query.QueryRow().Scan(&u.Id); err != nil {
		return err
	}
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
		SetMap(sq.Eq(update)).
		Set("updated_at", u.UpdatedAt)

	for k, v := range where {
		query = query.Where(sq.Eq{k: v})
	}

	sql, _, err := query.ToSql()
	if err != nil {
		return err
	} else {
		utils.GetLog().Debug("models.user.Update sql = : %s", sql)
	}

	result, err := query.RunWith(GetDB()).
		PlaceholderFormat(sq.Dollar).
		Exec()

	if err != nil {
		return err
	}
	if n, err := result.RowsAffected(); n == 0 && err != nil {
		return err
	}

	return nil
}

func (u *User) Delete(where map[string]interface{}) error {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("models.user.Delete error: %v", err)
		}
	}()

	update := map[string]interface{}{
		"deleted_at": time.Now(),
	}

	if err := u.Update(where, update); err != nil {
		return err
	}

	return nil
}

func (u *User) IsDeleted() bool {
	return u.DeletedAt.IsZero()
}

func FindUsers(where map[string]interface{}) ([]*User, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("models.user.FindUsers error: %v", err)
		}
	}()

	query := sq.Select("*").From(User{}.TableName())
	for k, v := range where {
		query = query.Where(sq.Eq{k: v})
	}

	rows, err := query.RunWith(GetDB()).
		PlaceholderFormat(sq.Dollar).
		Query()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*User{}
	for rows.Next() {
		var u User
		err = rows.Scan(&u.Id, &u.Phone, &u.Nickname, &u.Password, &u.Salt, &u.Weight, &u.Height, &u.Latitude, &u.Longitude, &u.Token, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, &u)
	}
	return users, nil
}
