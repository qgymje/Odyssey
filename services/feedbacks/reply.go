package feedbacks

import (
	"Odyssey/models"
	"Odyssey/services/notices"
)

type FeedbackReply struct {
	feedbackModel *models.Feedback
	content       string
}

// Reply 回复用户反馈
func (r *FeedbackReply) Do(feedbackID int64) (err error) {
	return
}

func (r *FeedbackReply) addNotice() (err error) {
	return notices.AddNotice(r)
}

//=========================================
// start implement Notice interface
//=========================================
func (r *FeedbackReply) Type() models.NoticeType {
	return models.NoticeTypeFeedbackReply
}

func (r *FeedbackReply) Messages() map[int64]string {
	m := map[int64]string{
		r.feedbackModel.UserID.Int64: r.content,
	}
	return m
}

//=========================================
// end implement Notice interface
//=========================================
