package models

import "time"

// Game model 表示一场夜跑赛事
type Game struct {
	ID                 int64         `json:"game_id"`
	Name               string        `json:"name"`   // 赛事名, 通常不能多于64个utf8字符
	Slogan             NullString    `json:"slogan"` // 赛事口号, 标标题, 用于宣传
	OraginazerID       int64         `bson:"oraginzer_id" json:"oraginzer_id"`
	MaximumParticipant int           `bson:"maximum_participant" json:"maximum_participant"` // -1 表示无限  参数人数限制
	MinumumParticipant NullInt       `bson:"minumum_participant" json:"minumum_participant"`
	Cost               float32       `bson:"const" json:"cost"`                  // 参加比赛费用
	RegisterTime       time.Time     `bson:"register_time" json:"register_time"` // 比赛报名开始时间
	StartTime          time.Time     `bson:"start_time" json:"start_time"`       // 比赛开始时间
	Duration           time.Duration `bson:"duration"`                           // 比赛持续时间
	RouteID            int64         `bson:"route_id" json:"route_id"`           //如何在输出的时候将route的数据带上?

	Distance         float32             `bson:"distance" json:"distance"` // 总公里数
	ValidationStatus GameValidationState `bson:"validation_status" json:"validation_status"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt NullTime  `bson:"updated_at" json:"updated_at"`

	Route *Route `json:"route"` // 线路
}

// Route model 描述一个比赛的路线
type Route struct {
	GameID             int64       `bson:"game_id" json:"game_id"` // 集合点
	GatheringLatitude  NullFloat64 `bson:"gathering_latitude" json:"gathering_latitude"`
	GatheringLongitude NullFloat64 `bson:"gathering_longitude" json:"gathering_longitude"`
	StartLatitude      float64     `bson:"start_latitude" json:"start_latitude"` // 起跑点
	StartLongitude     float64     `bson:"start_longitude" json:"start_longitude"`
	FinishLatitude     float64     `bson:"finish_latitude" json:"finish_latitude"` // 结束点
	FinishLongitude    float64     `bson:"finish_longitude" json:"finish_longitude"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt NullTime  `bson:"updated_at" josn:"-"`

	Locations   []*RouteLocation   // 定义为数组, 奇数位为latitude, 偶数位为logitude, 在pg里也保存也array // 用于保存路线地图的关键点, 用于显示线路图
	Annotations []*RouteAnnotation // 用于显示地图上一些点, 比如厕所点, 领奖点等
}

// RouteLocation 纪录一个线路上用于划线的坐标点
type RouteLocation struct {
	Latitude  float64 `bson:"latitude" json:"latitude"`
	Longitude float64 `bson:"longitude" json:"longitude"`
}

// RouteAnnotation model 表示比赛线路上的标注的点, 比如早餐点, 休息站, 等等,
type RouteAnnotation struct {
	Title     string     `bson:"title" json:"title"`
	Subtitle  NullString `bson:"subtitle" json:"subtitle"`
	Image     NullString `bson:"image" json:"image"`
	Latitude  float64    `bson:"latitude" json:"latitude"`
	Longitude float64    `bson:"longitude" json:"longitude"`
	Mark      NullString `bson:"mark" json:"mark"`
	CreatedAt time.Time  `bson:"created_at" json:"created_at"`
}
