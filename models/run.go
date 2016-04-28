package models

import (
	"Odyssey/utils"
	"time"

	sq "github.com/lann/squirrel"
)

// 一个跑步的纪录
type Run struct {
	Id        uint64
	UserId    uint64
	Distance  float64
	Duration  float64
	IsPublic  bool   //是否发布?
	Comment   string // 自己的评价
	Locations []Location

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// 仿照iOS CLLocation的结构
type Location struct {
	Latitude  float64   `json:"lat"`
	Longitude float64   `json:"lng"`
	Altitude  float64   `json:"alt"`
	TimeStamp time.Time `json: "ts"`
	Course    float64   `json:"course"`
	Speed     float64   `json:"speed"`
}

func (Run) TableName() string {
	return "runs"
}

func FindRuns(where map[string]interface{}) ([]*Run, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("models.run.FindRun error: %v", err)
		}
	}()

	query := sq.Select("*").From(Run{}.TableName()).OrderBy("created_at desc")
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

	rs := []*Run{}
	for rows.Next() {
		var r Run
		err = rows.Scan(&r.Id, &r.UserId, &r.Distance, &r.Duration, &r.IsPublic, &r.Locations, &r.CreatedAt)
		if err != nil {
			return nil, err
		}
		rs = append(rs, &r)
	}
	return rs, nil
}
