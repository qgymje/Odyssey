package models

import "time"

// use gorm instead of squirrel

// Game model 表示一场夜跑赛事
type Game struct {
	Id uint64 `json:"game_id"`
	// 赛事名, 通常不能多于64个utf8字符
	Name string `json:"name"`
	// 赛事口号, 标标题, 用于宣传
	Slogan     string      `json:"slogan"`
	Oraginazer *Oraginazer `json:"oraginzer_id"`

	// 参数人数限制
	MaximumParticipant int `json:"maximum_participant"` // -1 表示无限
	MinumumParticipant int `json:"minumum_participant"` // -1 表示无限

	// 参加比赛费用
	Cost float32 `json:"cost"`

	// 比赛报名开始时间
	RegisterTime time.Time `json:"register_time"`
	// 比赛开始时间
	StartTime time.Time `json:"start_time"`
	// 比赛持续时间
	Duration time.Duration `json:"duration"`
	Route    *Route        `json:"route_id"` //如何在输出的时候将route的数据带上?

	// 总公里数
	Distance float32 `json:"distance"`

	CreatedAt time.Time `json:"created_at"`
}
