package models

/*
import (
	"Odyssey/utils"
	"time"
)

// Feedback model 表示用户发来的反馈
type Feedback struct {
	TableName struct{} `sql:"feedbacks"`
	ID        int
	User      *User
	Content   string
	IsRead    bool

	CreatedAt time.Time
}

// Create 纪录一个操作
func (f *Feedback) Create() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("models.feedback.Create error: ", err)
		}
	}()

	f.CreatedAt = time.Now()

	err = GetDB().Create(f)

	return
}

// FindFeedbacks 查找反馈纪录
func FindFeedbacks(where map[string]interface{}, order string, limit int, offset int) (feedbacks []*Feedback, err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("models.feedback.FindFeedbacsk error: ", err)
		}
	}()
	query := GetDB().Model(&feedbacks)
	for key, val := range where {
		query = query.Where(key, val)
	}
	err = query.Order(order).Limit(limit).Offset(offset).Select()

	return
}

*/
