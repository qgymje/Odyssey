package likes

import "github.com/qgymje/Odyssey/models"

type RunLike struct {
	likeModel *models.RunLike
}

type RunLikeConfig struct {
	RunID, UserID int64
}

func NewRunLike(config *RunLikeConfig) *RunLike {
	return &RunLike{
		likeModel: &models.RunLike{
			RunID:  config.RunID,
			UserID: config.UserID,
		},
	}
}

func (l *RunLike) Like() (err error) {
	if err = l.likeModel.Like(); err != nil {
		return
	}
	return
}

func (l *RunLike) Unlike() (err error) {
	if err = l.likeModel.Unlike(); err != nil {
		return
	}
	return
}
