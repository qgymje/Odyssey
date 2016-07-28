package models

import (
	"Odyssey/utils"
	"time"
)

// RunLocation model纪录用户的跑步过程中GPS数据
// 仿照iOS CLRunLocation的结构
type RunLocation struct {
	TableName struct{}  `sql:"run_locations"`
	ID        int       `json:"id"`
	Run       *Run      `json:"run"`
	Latitude  float64   `json:"lat"`
	Longitude float64   `json:"lng"`
	Altitude  float64   `json:"alt"`
	Timestamp time.Time `json:"ts"`
	Course    float64   `json:"course"`
	Speed     float64   `json:"speed"`
	//Steps     int       `json:"stpes"` //距离上个location走出的步数
	//HeartRate

	CreatedAt time.Time `json:"-"`
}

// CreateRunLocations 跑步GPS数据入库
func CreateRunLocations(runID int, ls []RunLocation) (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("models.CreateRunLocations error: %v", err)
		}
	}()

	now := time.Now()
	for _, l := range ls {
		l.Run.ID = runID
		l.CreatedAt = now
	}

	err = GetDB().Create(&ls)

	return
}

// FindRunLocations 查找跑步GPS纪录
func FindRunLocations(where map[string]interface{}, order string, limit int, offset int) (ls []*RunLocation, err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("models.FindRunLocations error: %v", err)
		}
	}()

	query := GetDB().Model(&ls)
	for key, val := range where {
		query = query.Where(key, val)
	}
	err = query.Order(order).Limit(limit).Offset(offset).Select()

	return
}
