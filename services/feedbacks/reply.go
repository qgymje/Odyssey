package feedbacks

import (
	"Odyssey/forms"
	"Odyssey/models"
	"Odyssey/services/notifications"
	"Odyssey/utils"

	"github.com/pkg/errors"
)

type FeedbackReply struct {
	feedbackModel *models.Feedback
	content       string
}

func NewFeedbackReply(form *forms.FeedbackReplyForm) *FeedbackReply {
	return &FeedbackReply{
		feedbackModel: &models.Feedback{
			ID: form.FeedbackID,
		},
		content: form.Reply,
	}
}

// Do 回复用户反馈
func (r *FeedbackReply) Do() (err error) {
	defer func() {
		if err != nil {
			err = errors.Wrap(err, "services.feedbacks.FeedbackReply.Do error")
			utils.GetLog().Error("%+v", err)
		}
	}()

	if err = r.feedbackModel.Reply(r.content); err != nil {
		return
	}
	return
}

func (r *FeedbackReply) addNotice() (err error) {
	return notifications.AddNotice(r)
}

//=========================================
// start implement Notice interface
//=========================================
// Type 表明通知的类型
func (r *FeedbackReply) Type() models.EventType {
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
