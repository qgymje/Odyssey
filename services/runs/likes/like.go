package likes

import "Odyssey/models"

type RunLike struct {
	likeModel *models.RunLike
}

func NewRunLike(runID, userID int64) *RunLike {
	return &RunLike{
		likeModel: &models.RunLike{
			RunID:  runID,
			UserID: userID,
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
