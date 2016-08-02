package models

import "time"

// Feedback model 表示用户发来的反馈
type Feedback struct {
	ID        int64     `json:"feedback_id"`
	UserID    NullInt64 `db:"user_id" json:"user_id"` // 如果为null则为匿名用户
	Content   string    `json:"content"`
	IsRead    bool      `db:"is_read" json:"is_read"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`

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
	rows, err := GetDB().Queryx(`select f.id, f.user_id, f.content, f.is_read, f.created_at, u.id as user_id, u.phone, u.nickname, u.created_at from feedbacks as f left join users as u on f.user_id = u.id where f.user_id IS NOT NULL order by ? limit ?,?;`, order, offset, limit)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var f Feedback
		f.User = &User{}
		if err = rows.Scan(&f.ID,
			&f.UserID,
			&f.Content,
			&f.IsRead,
			&f.CreatedAt,
			&f.User.ID,
			&f.User.Phone,
			&f.User.Nickname,
			&f.User.CreatedAt,
		); err != nil {
			return nil, err
		}
		feedbacks = append(feedbacks, &f)
	}

	return feedbacks, nil
}
