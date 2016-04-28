package models

import (
	"Odyssey/utils"
	"time"

	sq "github.com/lann/squirrel"
)

const (
	DB_SMSCODE_PHONE = "phone"
	DB_SMSCODE_CODE  = "code"
)

type SMSCode struct {
	Id    uint64
	Phone string
	Code  string

	UsedAt    time.Time
	CreatedAt time.Time

	Base
}

func (SMSCode) TableName() string {
	return "smscodes"
}

func (s *SMSCode) Create() error {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("models.smscode.Create error: %v", err)
		}
	}()

	s.CreatedAt = time.Now()
	query := sq.Insert(s.TableName()).
		Columns("phone", "code", "used_at", "created_at").
		Values(s.Phone, s.Code, s.UsedAt, s.CreatedAt)

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

func (s *SMSCode) IsUsed() bool {
	if s.UsedAt.IsZero() {
		return false
	}
	return true
}

func (s *SMSCode) Update(where map[string]interface{}, update map[string]interface{}) error {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("models.smscode.Update error: %v", err)
		}
	}()

	query := sq.Update(s.TableName()).
		SetMap(sq.Eq(update))

	for k, v := range where {
		query = query.Where(sq.Eq{k: v})
	}

	sql, _, err := query.ToSql()
	if err != nil {
		return err
	} else {
		utils.GetLog().Debug("models.smscode.Update sql = : %s", sql)
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

func FindSMSCode(where map[string]interface{}) (*SMSCode, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("models.feedback.FindSMSCode error: %v", err)
		}
	}()

	query := sq.Select("*").From(SMSCode{}.TableName()).OrderBy("created_at desc").Limit(1)
	for k, v := range where {
		query = query.Where(sq.Eq{k: v})
	}

	s := &SMSCode{}
	err = query.RunWith(GetDB()).
		PlaceholderFormat(sq.Dollar).
		QueryRow().
		Scan(&s.Id, &s.Phone, &s.Code, &s.UsedAt, &s.CreatedAt)

	if err != nil {
		return nil, err
	}

	return s, nil
}
