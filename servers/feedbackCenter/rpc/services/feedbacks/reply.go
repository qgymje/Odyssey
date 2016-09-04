package feedbacks

import (
	"github.com/qgymje/Odyssey/models"
	"github.com/qgymje/Odyssey/services/notifications"
	"github.com/qgymje/Odyssey/utils"

	"github.com/pkg/errors"
)

type FeedbackReply struct {
	feedbackModel *models.Feedback
	content       string
}

type FeedbackReplyConfig struct {
	FeedbackID int64
	Reply      string
}

func NewFeedbackReply(config *FeedbackReplyConfig) *FeedbackReply {
	return &FeedbackReply{
		feedbackModel: &models.Feedback{
			ID: config.FeedbackID,
		},
		content: config.Reply,
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

// Type 实现nofitications接口
//=========================================
// start implement Notice interface
//=========================================
// Type 表明通知的类型
func (r FeedbackReply) Type() models.EventType {
	return models.NoticeTypeFeedbackReply
}

func (r FeedbackReply) Messages() map[int64]string {
	m := map[int64]string{
		r.feedbackModel.UserID.Int64: r.content,
	}
	return m
}

//=========================================
// end implement Notice interface
//=========================================
