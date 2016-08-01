package models

import "time"

// RunComment model 表示对一次跑步纪录的评论
type RunComment struct {
	ID              int64     `json:"run_comment_id"`
	RunID           int       `json:"run_id"`
	CommentUserID   int       `db:"comment_user_id" json:"comment_user"`
	ParentCommentID int64     `db:"reply_comment" json:"reply_comment"` // 如果为空, 则表示对跑步纪录的评论, 不为空, 则为对用户的评论的评论
	CreatedAt       time.Time `db:"created_at" json:"created_at"`

	CommentUser   *User
	ParentComment *RunComment
}

// Create 创建一个评论
func (rc *RunComment) Create() (err error) {
	return
}
