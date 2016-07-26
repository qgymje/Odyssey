package models

import (
	"Odyssey/utils"
	"time"

	sq "github.com/lann/squirrel"
)

// Feedback model 表示用户发来的反馈
type Feedback struct {
	Id      uint64
	UserId  uint64
	Content string
	IsRead  bool

	CreatedAt time.Time
}

func (Feedback) TableName() string {
	return "feedbacks"
}

func (f *Feedback) Create() error {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("models.feedback.Create error: %v", err)
		}
	}()

	f.CreatedAt = time.Now()
	query := sq.Insert(f.TableName()).
		Columns("user_id", "content", "is_read", "created_at").
		Values(f.UserId, f.Content, f.IsRead, f.CreatedAt)

	result, err := query.RunWith(GetDB()).
		PlaceholderFormat(sq.Dollar).
		Exec()

	if err != nil {
		return err
	}
	if n, err := result.RowsAffected(); n == 0 && err != nil {
		return err
	}
	return nil
}

func FindFeedbacks(where map[string]interface{}) ([]*Feedback, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("models.feedback.FindFeedbacsk error: %v", err)
		}
	}()

	query := sq.Select("id, user_id, content, is_read, created_at").From(Feedback{}.TableName()).OrderBy("created_at desc")
	for k, v := range where {
		query = query.Where(sq.Eq{k: v})
	}

	rows, err := query.RunWith(GetDB()).
		PlaceholderFormat(sq.Dollar).
		Query()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var f Feedback
	fbs := []*Feedback{}
	for rows.Next() {
		err = rows.Scan(&f.Id, &f.UserId, &f.Content, &f.IsRead, &f.CreatedAt)
		if err != nil {
			return nil, err
		}
		fbs = append(fbs, &f)
	}
	return fbs, nil
}
