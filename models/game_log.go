package models

// GameLog 表示一次创建比赛的存照, 每次修改都会纪录, 用于统计最容易被修改的纪录, 优化前端操作
type GameLog struct {
	LogID int64 `db:"game_log_id" json:"game_log_id"`
	Game  `json:"game_id"`
}
