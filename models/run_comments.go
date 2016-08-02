package models

import "time"

// RunComment model 表示对一次跑步纪录的评论
type RunComment struct {
	ID              int64 `json:"run_comment_id"`
	RunID           int64 `db:"run_id" json:"run_id"`
	UserID          int64 `db:"user_id" json:"user_id"`
	ParentCommentID int64 `db:"parent_comment_id" json:"parent_comment_id"` // 如果为空, 则表示对跑步纪录的评论, 不为空, 则为对用户的评论的评论

	Content   string    `json:"conente"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`

	CommentUser *User
}

// Create 创建一个评论
func (rc *RunComment) Create() (err error) {
	return
}
