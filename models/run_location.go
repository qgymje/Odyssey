package models

import (
	"Odyssey/utils"
	"time"

	sq "github.com/lann/squirrel"
)

// RunLocation model纪录用户的跑步过程中GPS数据
// 仿照iOS CLRunLocation的结构
type RunLocation struct {
	Id        uint64    `json:"id"`
	RunId     uint64    `json:"run_id"`
	Latitude  float64   `json:"lat"`
	Longitude float64   `json:"lng"`
	Altitude  float64   `json:"alt"`
	Timestamp time.Time `json: "ts"`
	Course    float64   `json:"course"`
	Speed     float64   `json:"speed"`
	//Steps     int       `json:"stpes"` //距离上个location走出的步数
	//HeartRate

	CreatedAt time.Time
}

func (RunLocation) TableName() string {
	return "locations"
}

/*
func (l RunLocation) UnmarshalJSON([]byte) error {
	return nil
}
*/

type RunLocations []RunLocation

// 插入location data
func (ls RunLocations) Create(runId uint64) error {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("models.location.Create error: %v", err)
		}
	}()

	createdAt := time.Now()
	query := sq.Insert(RunLocation{}.TableName()).
		Columns("run_id", "latitude", "longitude", "altitude", "timestamp", "course", "speed", "created_at")

	for _, l := range ls {
		query = query.Values(runId, l.Latitude, l.Longitude, l.Altitude, l.Timestamp, l.Course, l.Speed, createdAt)
	}

	result, err := query.RunWith(GetDB()).PlaceholderFormat(sq.Dollar).Exec()
	if err != nil {
		return err
	}
	if n, err := result.RowsAffected(); n == 0 && err != nil {
		return err
	}

	return nil
}

func FindRunLocations(where map[string]interface{}) (RunLocations, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("models.FindRunLocations error: %v", err)
		}
	}()

	query := sq.Select("id, run_id, latitude, longitude, altitude, timestamp, course, speed, created_at").From(RunLocation{}.TableName()).OrderBy("created_at desc")
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

	var l RunLocation
	ls := RunLocations{}
	for rows.Next() {
		err = rows.Scan(&l.Id, &l.RunId, &l.Latitude, &l.Longitude, &l.Altitude, &l.Timestamp, &l.Course, &l.Speed, &l.CreatedAt)
		if err != nil {
			return nil, err
		}
		ls = append(ls, l)
	}
	return ls, nil
}
