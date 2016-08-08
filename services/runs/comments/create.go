package comments

import (
	"Odyssey/models"
	"Odyssey/utils"

	"github.com/pkg/errors"
)

type RunComment struct {
	commentModel *models.RunComment
}

type RunCommentConfig struct {
	RunID           int64
	Content         string
	ParentCommentID int64
	UserID          int64
}

func NewRunComment(config *RunCommentConfig) *RunComment {
	rc := &RunComment{
		commentModel: &models.RunComment{
			RunID:           config.RunID,
			UserID:          config.UserID,
			ParentCommentID: models.ToNullInt64(config.ParentCommentID),
			Content:         config.Content,
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
