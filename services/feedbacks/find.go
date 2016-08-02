package feedbacks

import (
	"Odyssey/models"
	"Odyssey/utils"

	"github.com/pkg/errors"
)

func Find() ([]*models.Feedback, error) {
	var fbs []*models.Feedback
	var err error
	defer func() {
		if err != nil {
			err = errors.Wrap(err, "services.feedbacks.Find error")
			utils.GetLog().Error("%+v", err)
		}
	}()

	fbs, err = models.FindFeedbacks("id desc", 20, 0)
	return fbs, err
}

func FindOne(feedbackID int64) {

}
