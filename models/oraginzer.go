package models

// Oraginazer model 表示一个组织者, 代表个人或者一个团队
// Oraginazer也是一个用户, 但拥有创建比赛的权力
type Oraginazer struct {
	Id   uint64 `json:"oraginzer_id"`
	User *User  `json:"user"`
}
