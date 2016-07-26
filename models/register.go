package models

import "time"

// Register model 表示用户报名比赛的纪录
type Register struct {
	Id        uint64 `json:"register_id"`
	User      *User
	Game      *Game
	Payment   *Payment
	CreatedAt time.Time
}
