package models

import "time"

// Notice model 表示一个通知
// 比如通知用户比赛
type Notice struct {
	Id          uint64 `json:"notice_id"`
	EventSource string `json:"event_source"` // 代表来源于哪个表的数据
	EventId     uint64 `json:"event_id"`
	Message     string
	ToUser      *User
	CreatedAt   time.Time
}
