package controllers

import (
	"github.com/qgymje/Odyssey/services/feedbacks"

	"github.com/gin-gonic/gin"
)

type FeedbackBinding struct {
	UserID  int64
	Content string `bs:"content" binding:"required"`

	*BaseBinding

	config *feedbacks.FeedbackConfig
}

func NewFeedbackBinding(c *gin.Context, userID int64) (*FeedbackBinding, error) {
	bs := &FeedbackBinding{
		BaseBinding: newBaseBinding(),
		UserID:      userID,
		config:      &feedbacks.FeedbackConfig{},
	}

	if err := c.Bind(bs); err != nil {
		bs.Msg.formatBindError(err)
		return bs, err
	}

	if err := bs.Valid(); err != nil {
		return bs, err
	}

	return bs, nil
}

func (fb *FeedbackBinding) Config() *feedbacks.FeedbackConfig {
	fb.config.UserID = fb.UserID
	fb.config.Content = fb.Content

	return fb.config
}

func (fb *FeedbackBinding) Valid() error {
	return nil
}

type FeedbackReplyBinding struct {
	FeedbackID int64  `bs:"feedback_id" binding:"required"`
	Reply      string `bs:"reply" binding:"required"`

	config *feedbacks.FeedbackReplyConfig

	*BaseBinding
}

func NewFeedbackReplyBinding(c *gin.Context) (*FeedbackReplyBinding, error) {
	bs := &FeedbackReplyBinding{
		BaseBinding: newBaseBinding(),
	}

	if err := c.Bind(bs); err != nil {
		bs.Msg.formatBindError(err)
		return bs, err
	}

	if err := bs.Valid(); err != nil {
		return bs, err
	}
	return bs, nil

}

func (fr *FeedbackReplyBinding) Config() *feedbacks.FeedbackReplyConfig {
	return fr.config
}

func (fr *FeedbackReplyBinding) Valid() (err error) {
	if err = fr.validReply(); err != nil {
		return
	}
	return
}

func (fr *FeedbackReplyBinding) validReply() (err error) {
	if fr.Reply == "" {
		fr.Msg.setError("reply", "回复为空")
		return
	}
	return
}
