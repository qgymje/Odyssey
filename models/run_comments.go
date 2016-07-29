package models

import "time"

// RunComment model 表示对一次跑步纪录的评论
type RunComment struct {
	ID           uint64      `json:"run_comment_id"`
	RunID        int         `json:"run_id"`
	CommentUser  *User       `json:"comment_user"`
	ReplyComment *RunComment `json:"reply_comment"` // 如果为空, 则表示对跑步纪录的评论, 不为空, 则为对用户的评论的评论
	CreatedAt    time.Time   `json:"created_at"`
}

// Create 创建一个评论
func (rc *RunComment) Create() (err error) {
	return
}
