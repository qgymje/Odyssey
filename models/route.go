package models

import "time"

// Route model 描述一个比赛的路线
type Route struct {
	ID                 int64       `json:"route_id"`             // 线路对应的比赛
	GameID             int64       `db:"game_id" json:"game_id"` // 集合点
	GatheringLatitude  NullFloat64 `db:"gathering_latitude" json:"gathering_latitude"`
	GatheringLongitude NullFloat64 `db:"gathering_longitude" json:"gathering_longitude"`
	StartLatitude      float64     `db:"start_latitude" json:"start_latitude"` // 起跑点
	StartLongitude     float64     `db:"start_longitude" json:"start_longitude"`
	FinishLatitude     float64     `db:"finish_latitude" json:"finish_latitude"` // 结束点
	FinishLongitude    float64     `db:"finish_longitude" json:"finish_longitude"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt NullTime  `db:"updated_at" josn:"-"`

	Locations   []*RouteLocation   // 定义为数组, 奇数位为latitude, 偶数位为logitude, 在pg里也保存也array // 用于保存路线地图的关键点, 用于显示线路图
	Annotations []*RouteAnnotation // 用于显示地图上一些点, 比如厕所点, 领奖点等
}

// RouteLocation 纪录一个线路上用于划线的坐标点
type RouteLocation struct {
	RouteID   int64   `db:"route_id" json:"route_id"`
	Latitude  float64 `db:"latitude" json:"latitude"`
	Longitude float64 `db:"longitude" json:"longitude"`
}
