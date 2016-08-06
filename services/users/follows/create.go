package follows

import (
	"Odyssey/forms"
	"Odyssey/models"
)

type UserFollow struct {
	modelFollow *models.Follow
}

func NewUserFollow(form forms.UserFollowForm) *UserFollow {
	uf := &UserFollow{
		modelFollow: &models.Follow{
			FromUserID: form.FromUserID,
			ToUserID:   form.ToUserID,
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

}

// CommonFollow 查找共同关注的人
func (f *UserFollow) CommonFollow() (err error) {

}
