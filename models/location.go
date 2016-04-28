package models

import (
	"Odyssey/utils"
	"time"

	sq "github.com/lann/squirrel"
)

// 仿照iOS CLLocation的结构
type Location struct {
	Id        uint64
	RunId     uint64
	Latitude  float64   `json:"lat"`
	Longitude float64   `json:"lng"`
	Altitude  float64   `json:"alt"`
	Timestamp time.Time `json: "ts"`
	Course    float64   `json:"course"`
	Speed     float64   `json:"speed"`

	CreatedAt time.Time
}

func (Location) TableName() string {
	return "locations"
}

type Locations []Location

// 插入location data
func (ls Locations) Create(runId uint64) error {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("models.location.Create error: %v", err)
		}
	}()

	createdAt := time.Now()
	query := sq.Insert(Location{}.TableName()).
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

func FindLocations(where map[string]interface{}) (Locations, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("models.FindLocations error: %v", err)
		}
	}()

	query := sq.Select("*").From(Location{}.TableName()).OrderBy("created_at desc")
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

	ls := Locations{}
	for rows.Next() {
		var l Location
		err = rows.Scan(&l.Id, &l.RunId, &l.Latitude, &l.Longitude, &l.Altitude, &l.Timestamp, &l.Course, &l.Speed, &l.CreatedAt)
		if err != nil {
			return nil, err
		}
		ls = append(ls, l)
	}
	return ls, nil
}
