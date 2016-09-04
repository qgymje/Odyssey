package follows

import "github.com/qgymje/Odyssey/models"

type UserFollow struct {
	modelFollow *models.Follow
}

type UserFollowConfig struct {
	FromUserID int64
	ToUserID   int64
}

func NewUserFollow(config *UserFollowConfig) *UserFollow {
	uf := &UserFollow{
		modelFollow: &models.Follow{
			FromUserID: config.FromUserID,
			ToUserID:   config.ToUserID,
		},
	}
	return uf
}

func (f *UserFollow) Follow() (err error) {
	if err = f.modelFollow.Follow(); err != nil {
		return
	}
	return
}

func (f *UserFollow) UnFollow() (err error) {
	if err = f.modelFollow.UnFollow(); err != nil {
		return
	}
	return
}

func (f *UserFollow) Followers() (err error) {
	return
}

// CommonFollow 查找共同关注的人
func (f *UserFollow) CommonFollow() (err error) {
	return
}
