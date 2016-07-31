package models

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

// RunLocation model纪录用户的跑步过程中GPS数据
// 仿照iOS CLRunLocation的结构
type RunLocation struct {
	ID        int64     `json:"run_location_id"`
	RunID     int64     `db:"run_id" json:"run_id"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Altitude  float64   `json:"altitude"`
	Timestamp time.Time `json:"timestamp"`
	Course    float64   `json:"course"`
	Speed     float64   `json:"speed"`
	CreatedAt time.Time `db:"created_at" json:"-"`
}

// CreateRunLocations 跑步GPS数据入库
func CreateRunLocations(runID int64, ls []*RunLocation) (err error) {
	now := time.Now()
	for i, _ := range ls {
		ls[i].RunID = runID
		ls[i].CreatedAt = now
	}

	vals := []interface{}{}
	rawSQL := `insert into run_locations(run_id, latitude, longitude, altitude, timestamp, course, speed,created_at) values `
	for _, l := range ls {
		rawSQL += `(?,?,?,?,?,?,?,?),`
		vals = append(vals, l.RunID, l.Latitude, l.Longitude, l.Altitude, l.Timestamp, l.Course, l.Speed, l.CreatedAt)
	}
	rawSQL = rawSQL[:len(rawSQL)-1]
	log.Println(rawSQL)

	result, err := GetDB().Exec(rawSQL, vals...)
	if err != nil {
		return
	}
	if cnt, err := result.RowsAffected(); err != nil {
		if int(cnt) != len(ls) {
			return fmt.Errorf("locations insert error: inserted %d, slice len is %d", cnt, len(ls))
		}
		return err
	}

	return
}

func FindRunLocations(runIDs []int64) (runLocations []*RunLocation, err error) {
	query, args, err := sqlx.In(`select * from run_locations where run_id in (?);`, runIDs)
	if err != nil {
		return
	}
	query = GetDB().Rebind(query)
	rows, err := GetDB().Queryx(query, args...)
	if err != nil {
		log.Println(err)
		return
	}
	for rows.Next() {
		var l RunLocation
		err = rows.StructScan(&l)
		runLocations = append(runLocations, &l)
	}

	return
}
