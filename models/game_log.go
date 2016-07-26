package models

// GameLog 表示一次创建比赛的存照, 每次修改都会纪录
type GameLog struct {
	Id   uint64 `json:"game_log_id"`
	Game `json:"game_id"`
}
