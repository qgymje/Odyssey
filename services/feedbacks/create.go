package feedbacks

import (
	"Odyssey/forms"
	"Odyssey/models"
)

type Feedback struct {
	fbModel *models.Feedback
}

func NewFeedback(form *forms.FeedbackForm) *Feedback {
	fb := new(Feedback)
	fb.fbModel = &models.Feedback{
		UserId:  form.UserId,
		Content: form.Content,
		IsRead:  false,
	}
	return fb
}

func (fb *Feedback) Do() error {
	if err := fb.fbModel.Create(); err != nil {
		return err
	}
	return nil
}
