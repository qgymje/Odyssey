package models

import "time"

// RouteAnnotation model 表示比赛线路上的标注的点, 比如早餐点, 休息站, 等等,
type RouteAnnotation struct {
	RouteID   int64      `db:"route_id" json:"route_id"` // 此annotation名字
	Title     string     `db:"title" json:"title"`
	Subtitle  NullString `db:"subtitle" json:"subtitle"`
	Image     NullString `db:"image" json:"image"`
	Latitude  float64    `db:"latitude" json:"latitude"`
	Longitude float64    `db:"longitude" json:"longitude"`
	Mark      NullString `db:"mark" json:"mark"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
}
