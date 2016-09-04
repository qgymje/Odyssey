package models

import "time"

// UserGroup model 表示一个跑团
type UserGroup struct {
	ID      uint64 `json:"user_group_id"`
	Creator *User  `json:"creator"`
	Name    string `json:"name"`
	Logo    string `json:"logo"`
	// 跑团最大人数
	MaximumUser int `json:"maximum_user"` // default为100人
	// 跑团成员列表
	Members   []*User   `json:"members"`
	CreatedAt time.Time `json:"created_at"`
}

// 用户组-用户 中间表
type UserGroupUser struct {
}
