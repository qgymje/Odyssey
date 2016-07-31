package models

import "time"

// Feedback model 表示用户发来的反馈
type Feedback struct {
	ID        int64
	UserID    NullInt64 `db:"user_id"` // 如果为null则为匿名用户
	Content   string
	IsRead    bool      `db:"is_read"`
	CreatedAt time.Time `db:"created_at"`

	User *User
}

type FeedbackTableInfo struct {
	TableName string
	UserID    string
	Content   string
	IsRead    string
	CreatedAt string
}

var FeedbackTable FeedbackTableInfo

func init() {
	FeedbackTable = FeedbackTableInfo{
		TableName: "feedbacks",
		UserID:    "user_id",
		Content:   "content",
		IsRead:    "is_read",
		CreatedAt: "created_at",
	}
}

// Create 纪录一个操作
func (f *Feedback) Create() (err error) {
	f.CreatedAt = time.Now()

	result, err := GetDB().NamedExec(`insert into feedbacks (user_id, content, is_read, created_at) value(:user_id, :content, :is_read, :created_at)`, f)
	if err != nil {
		return
	}
	if _, err = result.RowsAffected(); err != nil {
		return
	}
	f.ID, err = result.LastInsertId()
	return
}

// FindFeedbacks 查找反馈纪录
func FindFeedbacks(order string, limit int, offset int) ([]*Feedback, error) {
	var feedbacks []*Feedback
	var err error
	if err = GetDB().Select(&feedbacks, `select f.*, u.* from feedbacks as f left join users as u on f.user_id = u.id order by ? limit ?,?;`, order, offset, limit); err != nil {
		return nil, err
	}
	return feedbacks, nil
}
