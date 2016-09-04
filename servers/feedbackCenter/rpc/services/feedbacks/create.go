package feedbacks

import (
	"github.com/qgymje/Odyssey/models"
	"github.com/qgymje/Odyssey/utils"

	"github.com/pkg/errors"
)

type Feedback struct {
	fbModel *models.Feedback
}

type FeedbackConfig struct {
	UserID  int64
	Content string
}

func NewFeedback(config *FeedbackConfig) *Feedback {
	fb := new(Feedback)
	fb.fbModel = &models.Feedback{
		UserID:  models.ToNullInt64(config.UserID),
		Content: config.Content,
		IsRead:  false,
	}
	return fb
}

func (fb *Feedback) Do() (err error) {
	defer func() {
		if err != nil {
			err = errors.Wrap(err, "services.feedbacks.Do error")
			utils.GetLog().Error("%+v", err)
		}
	}()
	if err := fb.fbModel.Create(); err != nil {
		return err
	}
	return nil
}
