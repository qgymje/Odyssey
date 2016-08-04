package comments

import (
	"Odyssey/forms"
	"Odyssey/models"
	"Odyssey/utils"

	"github.com/pkg/errors"
)

type RunComment struct {
	commentModel *models.RunComment
}

func NewRunComment(form *forms.RunCommentForm) *RunComment {
	rc := &RunComment{
		commentModel: &models.RunComment{
			RunID:           form.RunID,
			UserID:          form.UserID,
			ParentCommentID: models.ToNullInt64(form.ParentCommentID),
			Content:         form.Content,
		},
	}
	return rc
}

func (c *RunComment) Do() (err error) {
	defer func() {
		if err != nil {
			err = errors.Wrap(err, "services.runs.comments.Do error")
			utils.GetLog().Error("%+v", err)
		}
	}()

	if err = c.save(); err != nil {
		return
	}
	return
	// maybe notify?
}

func (e *RunComment) CommentID() int64 {
	return e.commentModel.ID
}

func (c *RunComment) save() (err error) {
	return c.commentModel.Create()
}
