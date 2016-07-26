package models

import (
	"Odyssey/utils"
	"time"

	sq "github.com/lann/squirrel"
)

// 一个跑步的纪录
type Run struct {
	Id       uint64  `json:"run_id"`
	UserId   uint64  `json:"user_id"`
	Distance float64 `json:"distance"`
	Duration int     `json:"duration"`
	//Setps     int       `json:"steps"` // 步数
	IsPublic     bool         `json:"is_public"`
	Comment      string       `json:"comment"`
	RunLocations RunLocations `json:"run_locaitons"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (Run) TableName() string {
	return "runs"
}

func (r *Run) Create() error {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("models.run.Create error: %v", err)
		}
	}()

	r.CreatedAt = time.Now()
	query := sq.Insert(r.TableName()).
		Columns("user_id", "distance", "duration", "is_public", "comment", "created_at", "updated_at", "deleted_at").
		Values(r.UserId, r.Distance, r.Duration, r.IsPublic, r.Comment, r.CreatedAt, r.UpdatedAt, r.DeletedAt).
		Suffix("RETURNING \"id\"").
		RunWith(GetDB()).
		PlaceholderFormat(sq.Dollar)

	// 注意这里必须要传指针
	if err = query.QueryRow().Scan(&r.Id); err != nil {
		return err
	}
	return nil
}

func FindRuns(where map[string]interface{}) ([]*Run, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("models.run.FindRun error: %v", err)
		}
	}()

	query := sq.Select("id, user_id, distance, duration, is_public, comment, created_at, updated_at, deleted_at").From(Run{}.TableName()).OrderBy("created_at desc")
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

	var r Run
	rs := []*Run{}
	for rows.Next() {
		err = rows.Scan(&r.Id, &r.UserId, &r.Distance, &r.Duration, &r.IsPublic, &r.Comment, &r.CreatedAt, &r.UpdatedAt, &r.DeletedAt)
		if err != nil {
			return nil, err
		}
		rs = append(rs, &r)
	}
	return rs, nil
}
