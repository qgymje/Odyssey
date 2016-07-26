package models

// use gorm instead of squirrel

// Game model 表示一场夜跑赛事
type Game struct {
	Id uint64 `json:"game_id"`
	// 赛事名, 通常不能多于64个utf8字符
	Name string `json:"name"`
	// 赛事口号, 标标题, 用于宣传
	Slogan string `json:"slogan"`
	// 参数人数限制
	Route *Route `json:"route_id"` //如何在输出的时候将route的数据带上?

}
