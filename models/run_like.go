package models

import "time"

// RunLike 表示一个赞
type RunLike struct {
	ID         int64
	RunID      int64     `db:"run_id" json:"run_id"`
	UsreID     int64     `db:"user_id" json:"user_id"`
	IsCanceled bool      `db:"is_canceled" json:"-"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
}
