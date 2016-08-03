package models

import "time"

// use gorm instead of squirrel

// Game model 表示一场夜跑赛事
type Game struct {
	ID                 int64         `json:"game_id"`
	Name               string        `json:"name"`   // 赛事名, 通常不能多于64个utf8字符
	Slogan             string        `json:"slogan"` // 赛事口号, 标标题, 用于宣传
	OraginazerID       int64         `db:"oraginzer_id" json:"oraginzer_id"`
	MaximumParticipant int           `db:"maximum_participant" json:"maximum_participant"` // -1 表示无限 // 参数人数限制
	MinumumParticipant int           `db:"minumum_participant" json:"minumum_participant"` // -1 表示无限
	Cost               float32       `json:"cost"`                                         // 参加比赛费用
	RegisterTime       time.Time     `db:"register_time" json:"register_time"`             // 比赛报名开始时间
	StartTime          time.Time     `db:"start_time" json:"start_time"`                   // 比赛开始时间
	Duration           time.Duration `json:"duration"`                                     // 比赛持续时间
	RouteID            int64         `db:"route_id" json:"route_id"`                       //如何在输出的时候将route的数据带上?

	Distance  float32   `json:"distance"` // 总公里数
	CreatedAt time.Time `db:"created_at" json:"created_at"`

	Route *Route `json:"route"` // 线路
}
