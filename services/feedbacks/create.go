package feedbacks

import (
	"Odyssey/forms"
	"Odyssey/models"
	"Odyssey/utils"

	"github.com/pkg/errors"
)

type Feedback struct {
	fbModel *models.Feedback
}

func NewFeedback(form *forms.FeedbackForm) *Feedback {
	fb := new(Feedback)
	fb.fbModel = &models.Feedback{
		UserID:  models.MakeNullInt64(form.UserID),
		Content: form.Content,
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
