package models

import "time"

// RouteAnnotation model 表示比赛线路上的标注的点, 比如早餐点, 休息站, 等等,
type RouteAnnotation struct {
	ID        int64     `json:"route_annotation_id"`    // 对应的线路图
	RouteID   int64     `db:"route_id" json:"route_id"` // 此annotation名字
	Title     string    `json:"title"`
	Subtitle  string    `json:"subtitle"`
	Image     string    `json:"image"`
	Locations []float64 // 支持多个地点
	Mark      string    `json:"mark"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
